package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDb(cfg config) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(cfg.db.dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("cannot connect to database: %w", err)
	}
	sqlDB, err := db.DB()

	if err != nil {
		return nil, fmt.Errorf("error getting database connection pool: %w", err)
	}
	sqlDB.SetMaxOpenConns(cfg.db.maxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, _ := time.ParseDuration(cfg.db.maxIdleTime)
	sqlDB.SetConnMaxIdleTime(duration)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("cannot ping database: %w", err)
	}

	return db, nil
}
