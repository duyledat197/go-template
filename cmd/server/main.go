package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/duyledat197/go-template/config"
	"github.com/duyledat197/go-template/models/domain"
	"github.com/duyledat197/go-template/mongo"
	"github.com/duyledat197/go-template/service/auth"
	"github.com/duyledat197/go-template/service/user"
	httpService "github.com/duyledat197/go-template/transport/http"
	"github.com/go-kit/kit/log"
)

const (
	defaultPort       = config.PORT
	defaultMongoDBURL = config.MONGODBURL
	defaultDBName     = config.DBNAME
)

func main() {
	// env variable
	var (
		addr         = envString("PORT", defaultPort)
		dburl        = envString("MONGODB_URL", defaultMongoDBURL)
		dbname       = envString("DB_NAME", defaultDBName)
		httpAddr     = flag.String("http.addr", ":"+addr, "HTTP listen address")
		mongoDBURL   = flag.String("db.url", dburl, "MongoDB URL")
		databaseName = flag.String("db.name", dbname, "MongoDB database name")
	)

	// collection name
	var (
		USER = "user"
	)
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	// Setup repositories
	var (
		userRepo domain.UserRepository
	)
	mongoClient, err := mongo.NewMongoDbClient(*mongoDBURL)
	if err != nil {
		panic(err)
	}
	// defer closeConnection()
	// Setup collections
	userColl, err := mongo.NewMongoDbCtxCollectionWithClient(mongoClient, *databaseName, USER)
	if err != nil {
		panic(err)
	}
	userRepo = mongo.NewUserRepository(userColl)
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
