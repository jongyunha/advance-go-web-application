package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func GetDB(dbConfig *DbConfig) (sqlxDb *sqlx.DB, err error) {
	once.Do(func() {
		if db != nil {
			return
		}

		sqlxDb, err = sqlx.Connect("mysql", dbConfig.Dsn())
		if err != nil {
			Logger.Fatal("failed to connect to database", zap.Error(err))
		}

		db = sqlxDb
		err = db.Ping()
	})
	return db, err
}
