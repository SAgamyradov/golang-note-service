package main

import (
	"log"
	"net/http"
	"note-service/routes"
)

func main() {
	r := routes.InitRoutes()
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
