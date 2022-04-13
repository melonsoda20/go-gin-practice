package repositories

import (
	"context"
	"fmt"
	"gin-todo-app/database/entities"
	"gin-todo-app/models"
	"gin-todo-app/services"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/google/uuid"
)

func CreateToDo(app firebase.App, client firestore.Client, ctx context.Context, req models.CreateToDoReqDTO) (bool, models.GenericResponse) {
	todoData := entities.Todo{
		ID:        uuid.New(),
		Name:      req.Name,
		CreatedAt: time.Now(),
	}

	fmt.Println(todoData.ID.String())
	_, err := client.Collection("ToDo").NewDoc().Create(ctx, map[string]interface{}{
		"id":        todoData.ID.String(),
		"name":      todoData.Name,
		"createdAt": todoData.CreatedAt,
	})

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
