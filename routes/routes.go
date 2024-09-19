package routes

import (
	handler "note-service/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/notes", handler.AddNoteHandler).Methods("POST")
	r.HandleFunc("/notes", handler.GetNotesHandler).Methods("GET")

	r.HandleFunc("/auth/login", handler.LoginHandler).Methods("POST")
	return r
}
