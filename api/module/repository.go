//go:build wireinject
// +build wireinject

package module

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/jongyunha/advance-go-web-application/api/entity"
)

var userRepositorySet = wire.NewSet(
	entity.NewDefaultUserRepository,
	wire.Bind(new(entity.UserRepository), new(*entity.DefaultUserRepository)),
)

func InitializeUserRepository(db *sqlx.DB) entity.UserRepository {
	wire.Build(userRepositorySet)
	return nil
}
