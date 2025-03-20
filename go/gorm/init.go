package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	gormConfig = &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Silent),
		TranslateError: true,
	}
)

func InitDB() *gorm.DB {
	db := InitMysql()

	// Use *sql.DB interface to set connection pool for *gorm.DB
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(0)
	sqlDB.SetMaxOpenConns(2)

	return db
}

func InitMysql() *gorm.DB {
	const dsn = "root:my-secret-pw@tcp(localhost:3306)/mysql"

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		panic(err)
	}

	result := db.Exec("DROP DATABASE IF EXISTS work")
	if result.Error != nil {
		panic(result.Error)
	}

	result = db.Exec("CREATE DATABASE IF NOT EXISTS work")
	if result.Error != nil {
		panic(result.Error)
	}

	result = db.Exec("use work")
	if result.Error != nil {
		panic(result.Error)
	}

	return db
}

func InitSQLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), gormConfig)
	if err != nil {
		panic(err)
	}
	return db
}
