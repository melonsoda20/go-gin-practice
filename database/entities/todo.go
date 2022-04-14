package entities

import (
	"context"
	"time"

	firestore "cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type ToDoService interface {
	CreateToDo(client firestore.Client, ctx context.Context) (*firestore.WriteResult, error)
	GetAllToDo(client firestore.Client, ctx context.Context) ([]Todo, error)
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

func (t Todo) GetAllToDo(client firestore.Client, ctx context.Context) ([]Todo, error) {
	todoDatas := []Todo{}
	iter := client.Collection("ToDo").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return todoDatas, err
		}
		var todoData Todo
		mapErr := doc.DataTo(&todoData)
		if mapErr != nil {
			return todoDatas, mapErr
		}
		todoDatas = append(todoDatas, todoData)
	}

	return todoDatas, nil
}
