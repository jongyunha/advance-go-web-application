package apis

import (
	"fmt"
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
	var request *dto.CreateUserRequest
	if err := c.BindJSON(&request); err != nil {
		errMsg := fmt.Sprintf("failed to bind request, err: %v", err)
		err = errs.NewWarnError(http.StatusBadRequest, errs.InvalidRequest, errMsg)
		_ = c.Error(err)
		return
	}

	id, err := api.app.Service.UserService.Create(c.Request.Context(), request)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response := dto.SuccessResponse[int64]{
		Data:    id,
		Message: "user created successfully",
	}

	c.JSON(http.StatusOK, response)
}
