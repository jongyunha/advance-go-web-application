package apis

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jongyunha/advance-go-web-application/api/errs"
	"net/http"
)

// ErrorHandlingMiddleware is a middleware function that handles errors
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			ginErr := c.Errors.Last()
			var appErr *errs.ApplicationError
			if errors.As(ginErr.Err, &appErr) {
				appErr.Log()
				c.JSON(appErr.StatusCode(), gin.H{
					"code":    appErr.Code(),
					"message": appErr.Message(),
				})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": ginErr.Error(),
				})
			}
		}
	}
}
