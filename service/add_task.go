package service

import (
	"context"
	"fmt"

	"github.com/takurooo/go-todo-app/auth"
	"github.com/takurooo/go-todo-app/entity"
	"github.com/takurooo/go-todo-app/store"
)

type AddTaskService struct {
	DB   store.Execer
	Repo TaskAdder
}

func (a *AddTaskService) AddTask(ctx context.Context, title string) (*entity.Task, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	t := &entity.Task{
		UserID: id,
		Title:  title,
		Status: entity.TaskStatusTodo,
	}
	err := a.Repo.AddTask(ctx, a.DB, t)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return t, nil
}
