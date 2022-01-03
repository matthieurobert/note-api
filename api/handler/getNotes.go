package handler

import (
	"encoding/json"
	"net/http"

	"github.com/matthieurobert/amp/api/entity"
)

func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	notes, err := entity.GetNotes()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(notes)
}
