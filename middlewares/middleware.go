package middlewares

import (
	"gin-todo-app/constants"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func SetMiddleware(app firebase.App, client firestore.Client, redis redis.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(constants.MiddlewareKeysConst.FirebaseAppKey, app)
		c.Set(constants.MiddlewareKeysConst.FirebaseClientKey, client)
		c.Set(constants.MiddlewareKeysConst.RedisAppKey, redis)
		c.Next()
	}
}
