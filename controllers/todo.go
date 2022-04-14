package controllers

import (
	"gin-todo-app/database/repositories"
	"gin-todo-app/models"
	"gin-todo-app/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateToDo(c *gin.Context) {
	ctx := services.GetBackgroundContext()
	req := models.CreateToDoReqDTO{}

	bind_err := c.ShouldBindJSON(&req)
	if bind_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bind_err.Error()})
		return
	}

	client, isClientExists := services.GetFirestoreClient(c)
	if !isClientExists {
		services.LogErrorMessage("firestore client does not exists")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't create todo at the moment"})
		return
	}

	validation_results := services.ValidateCreateToDo(req)
	if !validation_results.IsSuccess {
		message := strings.Join(validation_results.Message[:], ",")
		services.LogErrorMessage("Validation error: " + message)
		c.JSON(http.StatusBadRequest, gin.H{"error": validation_results.Message})
		return
	}

	isCreateSuccessful, results := repositories.CreateToDo(client, ctx, req)
	if !isCreateSuccessful {
		services.LogErrorMessage(results.Data.(string))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results": results,
	})

}
