package entities

import (
	"context"
	"time"

	firestore "cloud.google.com/go/firestore"
)

type ToDoService interface {
	CreateToDo(client firestore.Client, ctx context.Context) (*firestore.WriteResult, error)
}

type Todo struct {
	Name       string
	IsTaskDone bool
	CreatedAt  time.Time
}

func (t Todo) CreateToDo(client firestore.Client, ctx context.Context) (*firestore.WriteResult, error) {
	res, err := client.Collection("ToDo").NewDoc().Create(ctx, map[string]interface{}{
		"name":       t.Name,
		"createdAt":  t.CreatedAt,
		"isTaskDone": t.IsTaskDone,
	})

	return res, err
}
