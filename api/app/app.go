package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/jongyunha/advance-go-web-application/api/core"
)

type App struct {
	db         *sqlx.DB
	config     *core.AppConfig
	Service    *Services
	Repository *Repositories
}

func (a *App) GetConfig() *core.AppConfig {
	return a.config
}

func New(stage core.Stage) (*App, error) {
	err := core.InitLogger(stage)
	if err != nil {
		return nil, err
	}

	appConfig, err := core.NewAppConfig(core.Development)
	if err != nil {
		return nil, err
	}

	db, err := core.GetDB(appConfig.DbConfig)
	if err != nil {
		return nil, err
	}

	newApp := &App{
		db:     db,
		config: appConfig,
	}

	newApp.Repository = NewRepositories(*newApp)
	newApp.Service = NewServices(*newApp)

	return newApp, nil
}
