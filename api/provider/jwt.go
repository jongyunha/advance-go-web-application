package provider

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jongyunha/advance-go-web-application/api/entity"
	"github.com/jongyunha/advance-go-web-application/api/errs"
	"time"
)

const (
	identityKey = "id"
)

func InitParams() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: payloadFunc(),

		IdentityHandler: identityHandler(),
		//Authenticator:   authenticator(),
		//Authorizator:    authorizator(),
		Unauthorized:  unauthorized(),
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}

func payloadFunc() func(data any) jwt.MapClaims {
	return func(data any) jwt.MapClaims {
		if v, ok := data.(*entity.User); ok {
			return jwt.MapClaims{
				identityKey: v.ID,
			}
		}
		return jwt.MapClaims{}
	}
}

func identityHandler() func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		id := claims[identityKey].(int64)
		return &entity.User{
			ID: &id,
		}
	}
}

func unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		err := errs.NewWarnError(code, errs.Unauthorized, message)
		_ = c.Error(err)
	}
}
