package main

import (
	"context"
	"database/sql"
	"flag"
	"greenlight/internal/data"
	"greenlight/internal/jsonlog"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}

	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
}

type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
}

func main() {
	var cfg config

	// Get the variables from the command line
	flag.IntVar(&cfg.port, "port", 8000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "(dev|prod|staging|test)")

	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "PostgreSQL DSN")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	// Rate limiting configuration
	flag.Float64Var(&cfg.limiter.rps, "rate-limiter", 2, "The number of requests per second")
	flag.IntVar(&cfg.limiter.burst, "burst", 4, "The amount of tokens initially in the bucket")
	flag.BoolVar(&cfg.limiter.enabled, "enabled", true, "Enable rate limiter")

	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	defer db.Close()
	logger.PrintInfo("database pool connection established", nil)

	migrationDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	migrator, err := migrate.NewWithDatabaseInstance("file:///Users/nate/fun/greenlight/migrations", "postgres", migrationDriver)
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	err = migrator.Up()

	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err, nil)
	}

	logger.PrintInfo("Database migrations applied", nil)

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// Set the maximum number of open connections(in-use + idle) connections in the pool
	// Passing a value that is less than or equals to zero will mean no limit
	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	// Set the maximum number of idle connections in a pool
	// Passing a value that is less than or equals to zero will mean no limit
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	// Use the time.ParseDuration() function to convert the idle timeout string to
	// time.Duration type
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	// Create a context with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
