package middlewares

import (
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func SetMiddleware(app firebase.App, client firestore.Client, redis *redis.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("firebase_db", app)
		c.Set("firestore_client", client)
		c.Set("redis", redis)
		c.Next()
	}
}
