package service

import (
	"context"
	"fmt"

	"github.com/takurooo/go-todo-app/auth"
	"github.com/takurooo/go-todo-app/entity"
	"github.com/takurooo/go-todo-app/store"
)

type ListTaskService struct {
	DB   store.Queryer
	Repo TaskLister
}

func (l *ListTaskService) ListTasks(ctx context.Context) (entity.Tasks, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	ts, err := l.Repo.ListTasks(ctx, l.DB, id)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return ts, nil
}
