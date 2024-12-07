package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/gorm"
)

type db struct {
	dsn              string
	maxOpenConns     int
	maxIdleOpenConns int
	maxIdleTime      string
}

type config struct {
	port int
	db   db
}

type application struct {
	config config
	logger *log.Logger
	db     *gorm.DB
}

const version = "1.23"

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "Application starting port")

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
