package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	HTTPRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "fitrang",
			Subsystem: "profile_service",
			Name:      "http_requests_total",
			Help:      "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "fitrang",
			Subsystem: "profile_service",
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request latency",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	GraphQLOperationDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "fitrang",
			Subsystem: "profile_service",
			Name:      "graphql_operation_duration_seconds",
			Help:      "GraphQL operation execution time",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"operation", "type"},
	)

	KafkaProduceEnqueuedTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "fitrang",
			Subsystem: "profile_service",
			Name:      "kafka_produce_enqueued_total",
			Help:      "Kafka messages successfully enqueued for production",
		},
		[]string{"topic"},
	)

	KafkaProduceDeliveryTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "fitrang",
			Subsystem: "profile_service",
			Name:      "kafka_produce_delivery_total",
			Help:      "Kafka message delivery results",
		},
		[]string{"topic", "status"},
	)

	KafkaProduceLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "fitrang",
			Subsystem: "profile_service",
			Name:      "kafka_produce_latency_seconds",
			Help:      "End-to-end Kafka message delivery latency",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"topic"},
	)

	GRPCRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "fitrang",
			Subsystem: "profile_service",
			Name:      "grpc_requests_total",
			Help:      "Total gRPC requests handled",
		},
		[]string{"service", "method", "status"},
	)

	GRPCRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "fitrang",
			Subsystem: "profile_service",
			Name:      "grpc_request_duration_seconds",
			Help:      "gRPC request latency",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"service", "method"},
	)
)

func Register() {
	prometheus.MustRegister(
		HTTPRequestsTotal,
		HTTPRequestDuration,
		GraphQLOperationDuration,
		KafkaProduceEnqueuedTotal,
		KafkaProduceDeliveryTotal,
		KafkaProduceLatency,
		GRPCRequestsTotal,
		GRPCRequestDuration,
	)
}
