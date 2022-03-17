package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"greenlight/internal/data"
	"log"
	"net/http"
	"os"
	"time"
)

var version = "1.0.0"

type config struct {
	port int
	env string
	db struct {
		dsn string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime string
	}
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
}

func main() {
	var cfg config

	// Get the variables from the command line
	flag.IntVar(&cfg.port,"port",4000,"API server port")
	flag.StringVar(&cfg.env,"env","development","(dev|prod|staging|test)")

	flag.StringVar(&cfg.db.dsn,"db-dsn",os.Getenv("GREENLIGHT_DB_DSN"),"PostgreSQL DSN")

	flag.IntVar(&cfg.db.maxOpenConns,"db-max-open-conns",25,"PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns,"db-max-idle-conns",25,"PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime,"db-max-idle-time","15m","PostgreSQL max connection idle time")


	flag.Parse()

	logger := log.New(os.Stdout,"[INFO]\t",log.Ldate | log.Ltime)

	db,err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Printf("database pool connection established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d",cfg.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s",cfg.env,srv.Addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func openDB(cfg config) (*sql.DB,error) {
	db,err := sql.Open("postgres",cfg.db.dsn)
	if err != nil {
		return nil,err
	}

	// Set the maximum number of open connections(in-use + idle) connections in the pool
	// Passing a value that is less than or equals to zero will mean no limit
	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	// Set the maximum number of idle connections in a pool
	// Passing a value that is less than or equals to zero will mean no limit
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	// Use the time.ParseDuration() function to convert the idle timeout string to
	// time.Duration type
	duration,err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil,err
	}

	db.SetConnMaxIdleTime(duration)

	// Create a context with 5 second timeout
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil,err
	}

	return db,nil
}
