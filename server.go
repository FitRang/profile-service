package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Foxtrot-14/FitRang/profile-service/cache"
	"github.com/Foxtrot-14/FitRang/profile-service/config"
	"github.com/Foxtrot-14/FitRang/profile-service/eventbus"
	"github.com/Foxtrot-14/FitRang/profile-service/graph"
	"github.com/Foxtrot-14/FitRang/profile-service/metrics"
	"github.com/Foxtrot-14/FitRang/profile-service/middleware"
	"github.com/Foxtrot-14/FitRang/profile-service/proto"
	"github.com/Foxtrot-14/FitRang/profile-service/repository"
	"github.com/Foxtrot-14/FitRang/profile-service/services/access"
	"github.com/Foxtrot-14/FitRang/profile-service/services/dossier"
	"github.com/Foxtrot-14/FitRang/profile-service/services/profile"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
)

var (
	restPort    string = "8080"
	rpcPort     string = "8081"
)

func main() {
	godotenv.Load(".env")
	metrics.Register()
	kafkaCfg := config.LoadKafkaConfig()

	eventBus, err := eventbus.NewEventBus(eventbus.Config{
		Brokers:  kafkaCfg.Brokers,
	})
	if err != nil {
		log.Fatalf("failed to init event bus: %v", err)
	}

	kafkaProducer, err := eventBus.NewProducer()
	if err != nil {
		log.Fatalf("failed to init event bus: %v", err)
	}
	defer kafkaProducer.Close()

	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Mongo Connect Error: ", err)
	}

	db := client.Database("profile-service")

	profileRepo := repository.NewProfileRepository(db)
	profileService := profile.NewProfileService(profileRepo, kafkaProducer)

	dossierRepo := repository.NewDossierRepository(db)
	dossierService := dossier.NewDossierService(dossierRepo, profileRepo)

	accessRepo := repository.NewAccessRepository(db)
	accessService := access.NewAccessService(accessRepo, dossierRepo, profileRepo, kafkaProducer)
	repository.Init(db)

	resolver := &graph.Resolver{
		ProfileService: profileService,
		DossierService: dossierService,
		AccessService:  accessService,
	}

	srv := handler.New(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: resolver,
			},
		),
	)

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.Use(&middleware.GraphQLPrometheus{})

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/metrics", promhttp.Handler())
	http.Handle(
		"/query",
		middleware.PrometheusMiddleware(
			middleware.AuthMiddleware(srv),
		),
	)

	go func() {
		lis, err := net.Listen("tcp", ":"+rpcPort)
		if err != nil {
			log.Fatal("gRPC listen error:", err)
		}

		grpcServer := grpc.NewServer(
			grpc.UnaryInterceptor(
				middleware.PrometheusUnaryServerInterceptor(),
			),
		)
		rdb, err := cache.NewRedisClient(os.Getenv("REDIS_URI"))
		if err != nil {
			log.Fatal("while creating redis clien:", err)
		}
		log.Printf("redis connected:%s", os.Getenv("REDIS_URI"))

		profileGRPCService := proto.NewProfileGRPCService(
			profileService,
			dossierService,
			rdb,
		)

		proto.RegisterProfileServiceServer(grpcServer, profileGRPCService)

		log.Printf("gRPC server listening on :%s", rpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("Profile Service (HTTP) running on :%s", restPort)
	log.Fatal(http.ListenAndServe(":"+restPort, nil))
}
