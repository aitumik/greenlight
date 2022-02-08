package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var version = "1.0.0"

type config struct {
	port int
	env string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// Get the variables from the command line
	flag.IntVar(&cfg.port,"port",4000,"API server port")
	flag.StringVar(&cfg.env,"env","development","(dev|prod|staging|test)")
	flag.Parse()

	logger := log.New(os.Stdout,"[INSIGHT]\t",log.Ldate | log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d",cfg.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s",cfg.env,cfg.port)
	err := srv.ListenAndServe()
	log.Fatal(err)
}