package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Dsn struct {
	dsn          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type config struct {
	port int
	db   Dsn
}

type application struct {
	config config
	logger *log.Logger
	db     *gorm.DB
}

const version = "1.23"

func main() {
	var cfg config

	err := godotenv.Load(".env")

	if err != nil {
		log.Println("No .env file found")
	}
	dsn := os.Getenv("DB_STRING")

	fmt.Println(dsn)

	flag.IntVar(&cfg.port, "port", 8080, "Application starting port")
	flag.StringVar(&cfg.db.dsn, "db-dsn", dsn, "POSTGRES SQL DATABASE DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "POSTGRES maximum open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "POSTGRES maximum idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "POSTGRES maximum idle time")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := connectDb(cfg)

	if err != nil {
		logger.Fatal(err)
	}

	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	app := &application{
		config: cfg,
		logger: logger,
		db:     db,
	}

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	err = srv.ListenAndServe()
	logger.Fatal(err)
}
