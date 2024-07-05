package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"sync"
)

var once sync.Once

type Stage string

const (
	Development Stage = "development"
	Production  Stage = "production"
	Local       Stage = "local"
)

type AppConfig struct {
	Stage Stage
	*DbConfig
}

type DbConfig struct {
	DbHost            string
	DbPort            string
	DbUser            string
	DbPass            string
	DbName            string
	DbSsl             string
	DbMaxIdleConns    int
	DbMaxOpenConns    int
	DbConnMaxLifetime int
	DbConnMaxIdleTime int
}

func (d *DbConfig) Dsn() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", d.DbUser, d.DbPass, d.DbHost, d.DbPort, d.DbName)
}

func NewAppConfig(stage Stage) (*AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dbMaxOpenConns := os.Getenv("DB_MAX_OPEN_CONNS")
	dbMaxOpenConnsToInt, err := strconv.Atoi(dbMaxOpenConns)
	if err != nil {
		return nil, err
	}

	dbMaxIdleConns := os.Getenv("DB_MAX_IDLE_CONNS")
	dbMaxIdleConnsToInt, err := strconv.Atoi(dbMaxIdleConns)
	if err != nil {
		return nil, err
	}

	dbConnMaxLifetime := os.Getenv("DB_CONN_MAX_LIFETIME")
	dbConnMaxLifetimeToInt, err := strconv.Atoi(dbConnMaxLifetime)
	if err != nil {
		return nil, err
	}

	dbConnMaxIdleTime := os.Getenv("DB_CONN_MAX_IDLE_TIME")
	dbConnMaxIdleTimeToInt, err := strconv.Atoi(dbConnMaxIdleTime)
	if err != nil {
		return nil, err
	}

	dbConfig := &DbConfig{
		DbHost:            os.Getenv("DB_HOST"),
		DbPort:            os.Getenv("DB_PORT"),
		DbUser:            os.Getenv("DB_USER"),
		DbPass:            os.Getenv("DB_PASSWORD"),
		DbName:            os.Getenv("DB_NANE"),
		DbSsl:             os.Getenv("DB_SSL"),
		DbMaxOpenConns:    dbMaxOpenConnsToInt,
		DbMaxIdleConns:    dbMaxIdleConnsToInt,
		DbConnMaxLifetime: dbConnMaxLifetimeToInt,
		DbConnMaxIdleTime: dbConnMaxIdleTimeToInt,
	}
	return &AppConfig{
		Stage:    stage,
		DbConfig: dbConfig,
	}, nil
}
