package controllers

import (
	"fmt"
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

	isCreateSuccessful, results := services.CreateToDo(client, ctx, req)
	if !isCreateSuccessful {
		services.LogErrorMessage(results.Data.(string))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results": results,
	})

}

func GetAllToDo(c *gin.Context) {
	ctx := services.GetBackgroundContext()

	client, isClientExists := services.GetFirestoreClient(c)
	if !isClientExists {
		services.LogErrorMessage("firestore client does not exists")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't create todo at the moment"})
		return
	}

	isGetSuccessful, results := services.GetAllToDo(client, ctx)
	if !isGetSuccessful {
		services.LogErrorMessage(fmt.Sprintf("Error: %v", results.Data))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve ToDos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results": results,
	})
}

func GetToDo(c *gin.Context) {
	ctx := services.GetBackgroundContext()
	ID := c.Param("id")

	client, isClientExists := services.GetFirestoreClient(c)
	if !isClientExists {
		services.LogErrorMessage("firestore client does not exists")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't create todo at the moment"})
		return
	}

	isGetSuccessful, results := services.GetToDo(client, ctx, ID)
	if !isGetSuccessful {
		services.LogErrorMessage(fmt.Sprintf("Error: %v", results.Data))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve ToDos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results": results,
	})
}

func UpdateToDo(c *gin.Context) {
	ctx := services.GetBackgroundContext()
	ID := c.Param("id")
	req := models.UpdateToDoReqDTO{}

	bind_err := c.ShouldBindJSON(&req)
	if bind_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bind_err.Error()})
		return
	}

	client, isClientExists := services.GetFirestoreClient(c)
	if !isClientExists {
		services.LogErrorMessage("firestore client does not exists")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't update todo at the moment"})
		return
	}

	validation_results := services.ValidateUpdateToDo(req)
	if !validation_results.IsSuccess {
		message := strings.Join(validation_results.Message[:], ",")
		services.LogErrorMessage("Validation error: " + message)
		c.JSON(http.StatusBadRequest, gin.H{"error": validation_results.Message})
		return
	}

	isCreateSuccessful, results := services.UpdateToDo(client, ctx, req, ID)
	if !isCreateSuccessful {
		services.LogErrorMessage(results.Data.(string))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results": results,
	})
}
