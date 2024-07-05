package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/jongyunha/advance-go-web-application/api/config"
	"go.uber.org/zap"
	"log"
)

func main() {
	err := config.InitLogger(config.Development)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	appConfig, err := config.NewAppConfig(config.Development)
	if err != nil {
		config.Logger.Fatal("failed to load app config", zap.Error(err))
	}
	db, err := config.GetDB(appConfig.DbConfig)
	if err != nil {
		config.Logger.Fatal("failed to connect to database", zap.Error(err))
	}

	defer func(db *sqlx.DB) {
		_ = db.Close()
	}(db)
}
