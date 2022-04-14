package repositories

import (
	"context"
	"gin-todo-app/database/entities"
	"gin-todo-app/models"
	"gin-todo-app/services"
	"time"

	"cloud.google.com/go/firestore"
)

func CreateToDo(client firestore.Client, ctx context.Context, req models.CreateToDoReqDTO) (bool, models.GenericResponse) {
	todoData := entities.Todo{
		Name:       req.Name,
		CreatedAt:  time.Now().UTC(),
		IsTaskDone: req.IsTaskDone,
	}

	_, err := todoData.CreateToDo(client, ctx)

	if err != nil {
		services.LogError(err)
		errorResponse := models.ErrorResponse{
			Message: "Failed to create todo",
		}

		return false, models.GenericResponse{
			Data: errorResponse,
		}
	}

	return true, models.GenericResponse{
		Data: todoData,
	}
}
