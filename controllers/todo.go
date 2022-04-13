package controllers

import (
	"context"
	"gin-todo-app/database/repositories"
	"gin-todo-app/models"
	"gin-todo-app/services"
	"net/http"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

func CreateToDo(c *gin.Context) {
	ctx := context.Background()
	req := models.CreateToDoReqDTO{}

	bind_err := c.ShouldBindJSON(&req)
	if bind_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bind_err.Error()})
		return
	}

	firebase_db, _ := c.Get("firebase_db")
	app := firebase_db.(firebase.App)

	firestore_client, _ := c.Get("firestore_client")
	client := firestore_client.(firestore.Client)

	validation_results := services.ValidateCreateToDo(req)
	if !validation_results.IsSuccess {
		c.JSON(http.StatusBadRequest, gin.H{"error": validation_results.Message})
		return
	}

	isCreateSuccessful, results := repositories.CreateToDo(app, client, ctx, req)
	if !isCreateSuccessful {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results": results,
	})

}
