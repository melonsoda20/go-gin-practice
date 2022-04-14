package services

import (
	"context"
	"gin-todo-app/database/repositories"
	"gin-todo-app/models"

	firestore "cloud.google.com/go/firestore"
)

func CreateToDo(client firestore.Client, ctx context.Context, req models.CreateToDoReqDTO) (bool, models.GenericResponse) {
	return repositories.CreateToDo(client, ctx, req)
}

func GetAllToDo(client firestore.Client, ctx context.Context) (isSuccess bool, resp models.GenericResponse) {
	return repositories.GetAllToDo(client, ctx)
}

func GetToDo(client firestore.Client, ctx context.Context, ID string) (isSuccess bool, resp models.GenericResponse) {
	return repositories.GetToDo(client, ctx, ID)
}
