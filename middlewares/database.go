package middlewares

import (
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

func FirebaseMiddleware(app firebase.App, client firestore.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("firebase_db", app)
		c.Set("firestore_client", client)
		c.Next()
	}
}
