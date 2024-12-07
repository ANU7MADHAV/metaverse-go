package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDb(cfg config) (*gorm.DB, error) {
	fmt.Println("hello")

	dsn := "postgresql://neondb_owner:ecpnP1VJag5S@ep-withered-pine-a5byt7vo.us-east-2.aws.neon.tech/prisma_migrate_shadow_db_5a3bc837-e2c9-485b-be77-0776ef4c0379?sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("cannot connect to database: %w", err)
	}
	sqlDB, err := db.DB()

	if err != nil {
		return nil, fmt.Errorf("error getting database connection pool: %w", err)
	}
	sqlDB.SetMaxOpenConns(cfg.db.maxIdleOpenConns)
	sqlDB.SetMaxIdleConns(cfg.db.maxIdleOpenConns)

	duration, _ := time.ParseDuration(cfg.db.maxIdleTime)
	sqlDB.SetConnMaxIdleTime(duration)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("cannot ping database: %w", err)
	}

	return db, nil
}
