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
	GetToDo(client firestore.Client, ctx context.Context, ID string) (Todo, error)
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

func (t Todo) GetToDo(client firestore.Client, ctx context.Context, ID string) (Todo, error) {
	var todo Todo
	dsnap, err := client.Collection("ToDo").Doc(ID).Get(ctx)
	if err != nil {
		return todo, err
	}
	mapErr := dsnap.DataTo(&todo)
	if mapErr != nil {
		return todo, mapErr
	}
	return todo, nil
}

func (t Todo) UpdateToDo(client firestore.Client, ctx context.Context, ID string) (*firestore.WriteResult, error) {
	res, err := client.Collection("ToDo").Doc(ID).Set(ctx, map[string]interface{}{
		"name":       t.Name,
		"isTaskDone": t.IsTaskDone,
	}, firestore.MergeAll)

	return res, err
}
