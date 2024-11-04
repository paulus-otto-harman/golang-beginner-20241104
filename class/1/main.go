package main

import (
	"20241104/class/1/handler"
	producthandler "20241104/class/1/handler/product"
)

import (
	"log"
	"net/http"
)

func main() {
	// public route
	http.HandleFunc("POST /login", handler.Login)

	// private route
	http.HandleFunc("GET /products", producthandler.Index)

	log.Printf("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
