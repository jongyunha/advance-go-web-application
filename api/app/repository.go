package app

import (
	"github.com/jongyunha/advance-go-web-application/api/entity"
	"github.com/jongyunha/advance-go-web-application/api/module"
)

type Repositories struct {
	UserRepository entity.UserRepository
}

func NewRepositories(app App) *Repositories {
	return &Repositories{
		UserRepository: module.InitializeUserRepository(app.db),
	}
}
