package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/stamp-server/inmem"
	"github.com/stamp-server/models"
	"github.com/stamp-server/mongo"
	"github.com/stamp-server/service/auth"
	"github.com/stamp-server/service/user"
	httpService "github.com/stamp-server/transport/http"
	"gopkg.in/mgo.v2"
)

const (
	defaultPort = "8080"
	// defaultRoutingServiceURL = "http://localhost:7878"
	defaultMongoDBURL = "127.0.0.1"
	defaultDBName     = "stamp"
)

func main() {
	var (
		addr = envString("PORT", defaultPort)
		// rsurl  = envString("ROUTINGSERVICE_URL", defaultRoutingServiceURL)
		dburl  = envString("MONGODB_URL", defaultMongoDBURL)
		dbname = envString("DB_NAME", defaultDBName)

		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
		// routingServiceURL = flag.String("service.routing", rsurl, "routing service URL")
		mongoDBURL   = flag.String("db.url", dburl, "MongoDB URL")
		databaseName = flag.String("db.name", dbname, "MongoDB database name")
		inmemory     = flag.Bool("inmem", false, "use in-memory repositories")

		// ctx = context.Background()
	)
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	// Setup repositories
	var (
		userRepo models.UserRepository
	)
	if *inmemory {
		userRepo = inmem.NewUserRepository()
	} else {
		session, err := mgo.Dial(*mongoDBURL)
		if err != nil {
			panic(err)
		}
		defer session.Close()

		session.SetMode(mgo.Monotonic, true)
		userRepo = mongo.NewUserRepository(*databaseName, session, "user")
	}

	userService := user.NewService(userRepo)
	authService := auth.NewService(userRepo)
	var h http.Handler
	{
		h = httpService.NewHTTPHandler(
			userService,
			authService,
			logger,
		)
	}
	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, h)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	logger.Log("terminated", <-errs)
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
