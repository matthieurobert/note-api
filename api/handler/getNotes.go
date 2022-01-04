package handler

import (
	"encoding/json"
	"net/http"

	"github.com/matthieurobert/amp/api/auth"
	"github.com/matthieurobert/amp/api/entity"
)

func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	_, userInfo, _ := auth.Strategy.AuthenticateRequest(r)
	user, _ := entity.GetUserByUsername(userInfo.GetUserName())

	notes, err := entity.GetNotes(user.Id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(notes)
}
