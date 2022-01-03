package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matthieurobert/amp/api/entity"
)

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var note entity.Note

	err = json.Unmarshal(body, &note)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if note.Body == "" {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = entity.PostNote(&note)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("Task created")
}
