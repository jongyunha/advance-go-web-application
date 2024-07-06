package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/jongyunha/advance-go-web-application/api/app"
	"github.com/jongyunha/advance-go-web-application/api/dto"
	"github.com/jongyunha/advance-go-web-application/api/errs"
	"net/http"
)

func bindUserApi(
	app app.App,
	router *gin.Engine,
) {
	api := userApi{app: app}
	rg := router.Group("/api/v1")
	rg.POST("/users", api.create)
}

type userApi struct {
	app app.App
}

func (api *userApi) create(c *gin.Context) {
	var request dto.CreateUserRequest
	if err := c.BindJSON(&request); err != nil {
		err = errs.NewWarnError(http.StatusBadRequest, errs.InvalidRequest, "invalid request")
		_ = c.Error(err)
		return
	}
}
