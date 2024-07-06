package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jongyunha/advance-go-web-application/api/app"
)

func Serve(app *app.App) error {
	router := gin.Default()
	router.Use(ErrorHandlingMiddleware())
	bindUserApi(*app, router)
	return router.Run(fmt.Sprintf("%s:%s", app.GetConfig().Host, app.GetConfig().Port))
}
