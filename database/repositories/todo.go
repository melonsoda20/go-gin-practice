package repositories

import (
	"context"
	"gin-todo-app/database/entities"
	"gin-todo-app/models"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func GetToDo(client firestore.Client, ctx context.Context, ID string) (isSuccess bool, resp models.GenericResponse) {
	var todoData entities.Todo
	data, err := todoData.GetToDo(client, ctx, ID)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return true, models.GenericResponse{
				Data: nil,
			}
		} else {
			errorResponse := models.ErrorResponse{
				Message: err.Error(),
			}

			return false, models.GenericResponse{
				Data: errorResponse,
			}
		}
	}

	return true, models.GenericResponse{
		Data: data,
	}
}
