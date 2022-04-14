package services

import (
	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

func GetFirebaseApp(c *gin.Context) (value interface{}, exists bool) {
	firebase_db, exists := c.Get("firebase_db")
	if !exists {
		return nil, false
	}
	app := firebase_db.(firebase.App)

	return app, true
}

func GetFirestoreClient(c *gin.Context) (value firestore.Client, exists bool) {
	firestore_client, isClientExists := c.Get("firestore_client")
	client := firestore_client.(firestore.Client)

	return client, isClientExists
}
