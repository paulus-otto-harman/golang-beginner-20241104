package main

import (
	"20241104/class/2/handler"
	"20241104/class/2/handler/todo"
	"20241104/class/2/middleware"
	"log"
	"net/http"
)

func main() {

	// Public Routes
	publicMux := http.NewServeMux()

	publicMux.HandleFunc("POST /register", handler.UserRegistration)
	publicMux.HandleFunc("POST /login", handler.Login)

	// Private Routes
	privateMux := http.NewServeMux()

	privateMux.HandleFunc("GET /todos", todo.TodoIndex)
	privateMux.HandleFunc("POST /todos", todo.TodoCreate)
	privateMux.HandleFunc("PUT /todos/{id}", todo.TodoUpdate)
	privateMux.HandleFunc("DELETE /todos/{id}", todo.TodoDelete)

	auth := middleware.Auth(privateMux)

	serverMux := http.NewServeMux()
	serverMux.Handle("/", publicMux)
	serverMux.Handle("/user/", http.StripPrefix("/user", auth))

	log.Println("Server started on port 8080")

	http.ListenAndServe(":8080", serverMux)

}
