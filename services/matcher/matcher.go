package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/StevenRojas/donatePlasma/services/matcher/pkg/endpoints"
	"github.com/StevenRojas/donatePlasma/services/matcher/pkg/service"
	"github.com/StevenRojas/donatePlasma/services/matcher/pkg/transport"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/go-sql-driver/mysql"
)

var logger log.Logger

func main() {
	httpPort := getHTTPPort()
	logger = setLogger()
	level.Info(logger).Log("msg", "Matcher server starting")
	defer level.Info(logger).Log("msg", "Matcher server ended")

	db := connectToDb(getConnectionString())
	ctx := context.Background()
	repository := service.NewRepository(db, logger)
	service := service.NewService(repository, logger)
	endpoints := endpoints.MakeEndpoints(service)

	errs := make(chan error)

	go func() {
		level.Info(logger).Log("msg", "Listing on port "+httpPort)
		handler := transport.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(httpPort, handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	level.Error(logger).Log("exit", <-errs)
}

func connectToDb(conn string) *sql.DB {
	level.Info(logger).Log("msg", "Connecting to DB", conn)
	db, err := sql.Open("mysql", conn)
	level.Info(logger).Log("msg", "DB response ", err)
	if err != nil {
		level.Error(logger).Log("exit", "Error connecting database server", err)
		os.Exit(-1)
	}
	return db
}

func getHTTPPort() string {
	httpPort, ok := os.LookupEnv("MATCHER_PORT")
	if !ok {
		httpPort = "8011"
	}
	return ":" + httpPort
}

func getConnectionString() string {
	username, ok := os.LookupEnv("MYSQL_USERNAME")
	if !ok {
		username = "root"
	}
	password, ok := os.LookupEnv("MYSQL_PASSWORD")
	if !ok {
		password = "123"
	}
	host, ok := os.LookupEnv("MYSQL_HOST")
	if !ok {
		host = "localhost"
	}
	port, ok := os.LookupEnv("MYSQL_PORT")
	if !ok {
		port = "33016"
	}
	db, ok := os.LookupEnv("MYSQL_DB")
	if !ok {
		db = "donate"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, db)
}

func setLogger() log.Logger {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service:", "Matcher",
			"time:", log.DefaultTimestampUTC,
			"caller:", log.DefaultCaller,
		)
	}
	return logger
}
