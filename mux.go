package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/takurooo/go-todo-app/auth"
	"github.com/takurooo/go-todo-app/clock"
	"github.com/takurooo/go-todo-app/config"
	"github.com/takurooo/go-todo-app/handler"
	"github.com/takurooo/go-todo-app/service"
	"github.com/takurooo/go-todo-app/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	v := validator.New()
	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	clocker := clock.RealClocker{}

	r := store.Repository{Clocker: clocker}

	rcli, err := store.NewKVS(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	jwter, err := auth.NewJWTer(rcli, clocker)
	if err != nil {
		return nil, cleanup, err
	}

	ru := &handler.RegisterUser{
		Service:   &service.RegisterUser{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/register", ru.ServeHTTP)

	l := &handler.Login{
		Service: &service.Login{
			DB:             db,
			Repo:           &r,
			TokenGenerator: jwter,
		},
		Validator: v,
	}
	mux.Post("/login", l.ServeHTTP)

	at := &handler.AddTask{
		Service:   &service.AddTask{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/tasks", at.ServeHTTP)

	lt := &handler.ListTask{
		Service: &service.ListTask{DB: db, Repo: &r},
	}
	mux.Get("/tasks", lt.ServeHTTP)
	return mux, cleanup, nil
}
