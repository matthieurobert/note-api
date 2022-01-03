package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matthieurobert/amp/api/config"
	"github.com/matthieurobert/amp/api/entity"
	"github.com/matthieurobert/amp/api/handler"
)

func main() {
	config.Init()

	entity.CreateSchema(config.POSTGRES.DB)

	port := config.ENV.ApiPort
	fmt.Println(port)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World !")
	})

	r.HandleFunc("/notes", handler.GetNotesHandler).Methods("GET")
	r.HandleFunc("/notes", handler.PostTaskHandler).Methods("POST")

	fmt.Println("Server launched on port: " + strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), r)
}
