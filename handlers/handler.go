package handler

import (
	"encoding/json"
	"net/http"

	"note-service/middleware"
	"note-service/model"
	"note-service/service"
	"note-service/utils"
	"sync"
)

var (
	notes   = make([]model.Note, 0)
	notesMu sync.Mutex
)

func AddNoteHandler(w http.ResponseWriter, r *http.Request) {
	userID := 1
	var note model.Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	note.UserID = userID
	note.ID = len(notes) + 1

	// checking spelling
	isValid, err := service.CheckSpelling(note.Content)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Spell check error")
		return
	}
	if !isValid {
		utils.ErrorResponse(w, http.StatusBadRequest, "Spelling errors detected")
		return
	}

	notesMu.Lock()
	notes = append(notes, note)
	notesMu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	userID := 1
	var userNotes []model.Note

	notesMu.Lock()
	defer notesMu.Unlock()

	for _, note := range notes {
		if note.UserID == userID {
			userNotes = append(userNotes, note)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userNotes)
}

var (
	users = map[string]string{
		"user": "password", // simple hardcoded user
	}
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds map[string]string
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	password, ok := users[creds["username"]]
	if !ok || password != creds["password"] {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Generation JWT
	token, err := middleware.GenerateToken(creds["username"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Could not generate token")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
