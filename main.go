package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/RobsonFeitosa/configs/db"
)

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handles.Create)
	r.Put("/{id}", handles.Update)
	r.Delete("/{id}", handles.Delete)
	r.Get("/", handles.List)
	r.Get("/{id}", handles.Get)

	http.ListenAndServer(fmt.Sprintf(":%s", config.GetServerPort()), r)
}