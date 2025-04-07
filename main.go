package main

import (
    "database/sql"
    "flag"
    "log"
    "fmt"   
    "os"
    "os/signal"
    "time"
     _ "github.com/lib/pq"
    "github.com/FitRang/profile-service/domain"
    "github.com/FitRang/profile-service/handlers"
    "github.com/FitRang/profile-service/routes"

    "github.com/gin-gonic/gin"
)

var (
    // Derived flags from ldflags -X
    buildRevision string
    buildVersion  string
    buildTime     string

    // general options
    versionFlag bool
    helpFlag    bool

    //server port
    port string

    // program controller
    done      = make(chan struct{})  
    errc      = make(chan error)
)

func init() {
    flag.BoolVar(&versionFlag, "version", false, "show current verison and exit")
    flag.BoolVar(&helpFlag, "help", false, "show usage and exit")
    flag.StringVar(&port, "port", ":4444", "server port")
}

func setBuildVariables() {
    if buildRevision == "" {
        buildRevision = "dev"
    }
    if buildVersion == "" {
        buildRevision = "dev"
    }
    if buildTime == "" {
        buildTime = time.Now().UTC().Format(time.RFC3339)
    }
}

func parseFlags() {
    flag.Parse()

    if helpsFlag {
        flag.Usage()
        os.Exit(0)
    }
    if versionFlag {
        fmt.Printf("%s %s %s\n", buildRevision, buildVersion, buildTime)
        os.Exit(0)
    }
}

func hanldeInterrupts() {
    log.Println("start handle interrupts")

    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt)

    sig := <- interrupt
    log.Printf("caught sig: %v", sig)
    // close resource here
    done <- struct{}{}
}

func openDB() (*sql.DB, error) {
    var (
        host = "localhost"
        port = 5432
        user = "postgres"
        password = "postgres"
        dbname = "postgres"
    )

	//psqlInfo := os.Getenv("POSTGRESQL_CONN_STRING")
    psqlInfo := "postgresql://postgres:[YOUR-PASSWORD]@db.ygnqeraxlvjozannsnyy.supabase.co:5432/postgres"
    if len(psqlInfo) == 0 {
        psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    }
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }
    return db, nil
}

func main() {
    setBuildVariables()
    parseFlags()
    go handleInterrupts()

    server := gin.Default()
    psqlInfo, err := openDB()
    if err != nil {
        log.Printf("error connecting DB: %v", err)
        return
    }
    log.Println("DB connection is successful")
    defer psqlInfo.Close()

    // create a profile service
    profileService := domain.NewProfileService(psqlInfo)

    profileHandler := hanlders.NewProfileHandler(profileService)
    apiRoutes := routes.NewRoutes(profileHandler)
    routes.AttachRoutes(server, apiRoutes)

    go func() {
        errc <- server.Run(port)
    }()

    select {
    case err := <- errc:
        log.Printf("ListenAndServe error: %v", err)
    case <-done:
        log.Println("shutting down server ...")
    }
    time.AfterFunc(1*time.Second, func() {
        close(done)
        close(errc)
    })
}
