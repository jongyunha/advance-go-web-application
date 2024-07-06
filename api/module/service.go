//go:build wireinject
// +build wireinject

package module

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/jongyunha/advance-go-web-application/api/entity"
	"github.com/jongyunha/advance-go-web-application/api/service"
)

var userServiceSet = wire.NewSet(
	entity.NewSqlxTransactionManager,
	entity.NewDefaultUserRepository,
	wire.Bind(new(entity.UserRepository), new(*entity.DefaultUserRepository)),
	wire.Bind(new(entity.TransactionManager), new(*entity.SqlxTransactionManager)),
	service.NewDefaultUserService,
	wire.Bind(new(service.UserService), new(*service.DefaultUserService)),
)

func InitializeUserService(db *sqlx.DB) service.UserService {
	wire.Build(userServiceSet)
	return nil
}
