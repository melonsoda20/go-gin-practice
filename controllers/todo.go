package controllers

import (
	"fmt"
	"gin-todo-app/constants"
	"gin-todo-app/models"
	"gin-todo-app/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// CreateToDo example
// @Summary      Create New ToDo
// @Description  Create New ToDo Data
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param   some_id      path   int     true  "Some ID"
// @Param   some_id      body web.Pet true  "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't retrieve todo at the moment"})
		return
	}

	redis, redisExists := services.GetRedis(c)

	if redisExists {
		isCacheExists, exists_err := services.CheckCacheExists(redis, constants.CacheKeys.GetAllToDoKey)
		if exists_err != nil {
			services.LogError(exists_err)
			services.LogErrorMessage("ToDo is not going to be retrieved from redis")
		} else {
			if isCacheExists.(bool) {
				toDoFromCache, err := services.GetCacheData(redis, constants.CacheKeys.GetAllToDoKey)
				if err != nil {
					services.LogError(err)
					services.LogErrorMessage("ToDo is not going to be retrieved from redis")
				}
				toDoInBytes := []byte(fmt.Sprint(toDoFromCache))
				results := models.GenericResponse{}

				json_err := services.DeserializeJSON(toDoInBytes, &results)
				if json_err != nil {
					services.LogErrorMessage("Failed to deserialize ToDo data from cache")
				} else {
					c.JSON(http.StatusOK, gin.H{
						"results": results,
					})
				}
			}
		}
	} else {
		services.LogErrorMessage(fmt.Sprintf("Error: %v", "Failed to retrieve redis from middleware"))
		services.LogErrorMessage("ToDo is not going to be retrieved from redis")
	}

	isGetSuccessful, results := services.GetAllToDo(client, ctx)
	if !isGetSuccessful {
		services.LogErrorMessage(fmt.Sprintf("Error: %v", results.Data))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve ToDos"})
		return
	}

	if redisExists {
		results_json, json_err := services.SerializeJson(results)
		if json_err != nil {
			services.LogErrorMessage(fmt.Sprintf("Error: %v", json_err.Error()))
			services.LogErrorMessage("ToDo is not going to be stored to redis")
		} else {
			services.SetCacheData(redis, constants.CacheKeys.GetAllToDoKey, string(results_json))
		}
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

func DeleteToDo(c *gin.Context) {
	ctx := services.GetBackgroundContext()
	ID := c.Param("id")

	client, isClientExists := services.GetFirestoreClient(c)
	if !isClientExists {
		services.LogErrorMessage("firestore client does not exists")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't create todo at the moment"})
		return
	}

	isGetSuccessful, results := services.DeleteToDo(client, ctx, ID)
	if !isGetSuccessful {
		services.LogErrorMessage(fmt.Sprintf("Error: %v", results.Data))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ToDo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results": results,
	})
}
