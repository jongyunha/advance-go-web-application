package apis

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jongyunha/advance-go-web-application/api/app"
	"github.com/jongyunha/advance-go-web-application/api/provider"
)

func Serve(app *app.App) error {
	router := gin.Default()
	router.Use(ErrorHandlingMiddleware())

	authMiddleware, err := jwt.New(provider.InitParams())
	if err != nil {
		return err
	}
	router.Use(authMiddleware.MiddlewareFunc())

	bindUserApi(*app, router)
	return router.Run(fmt.Sprintf("%s:%s", app.GetConfig().Host, app.GetConfig().Port))
}
