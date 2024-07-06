package app

import (
	"github.com/jongyunha/advance-go-web-application/api/module"
	"github.com/jongyunha/advance-go-web-application/api/service"
)

type Services struct {
	UserService service.UserService
}

func NewServices(app App) *Services {
	return &Services{
		UserService: module.InitializeUserService(app.db),
	}
}
