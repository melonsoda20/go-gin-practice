package repositories

import (
	"context"
	"gin-todo-app/database/entities"
	"gin-todo-app/models"
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
		errorResponse := models.ErrorResponse{
			Message: err.Error(),
		}

		return false, models.GenericResponse{
			Data: errorResponse,
		}
	}

	return true, models.GenericResponse{
		Data: todoData,
	}
}

func GetAllToDo(client firestore.Client, ctx context.Context) (isSuccess bool, resp models.GenericResponse) {
	var todoData entities.Todo

	todoDatas, err := todoData.GetAllToDo(client, ctx)

	if err != nil {
		errorResponse := models.ErrorResponse{
			Message: err.Error(),
		}

		return false, models.GenericResponse{
			Data: errorResponse,
		}
	}

	return true, models.GenericResponse{
		Data: todoDatas,
	}
}
