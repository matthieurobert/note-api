package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/matthieurobert/amp/api/entity"
)

func Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var user entity.User

	err = json.Unmarshal(body, &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if user.Email == "" || user.Username == "" || user.Password == "" {
		http.Error(w, "missed data", http.StatusInternalServerError)
	}

	id, err := entity.PostUser(user.Username, user.Password, user.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(id)
}
