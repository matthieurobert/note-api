package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matthieurobert/amp/api/auth"
	"github.com/matthieurobert/amp/api/config"
	"github.com/matthieurobert/amp/api/entity"
	"github.com/matthieurobert/amp/api/handler"
	"github.com/matthieurobert/amp/api/middleware"
)

func main() {
	config.Init()
	auth.InitAuth()
	entity.CreateSchema(config.POSTGRES.DB)

	port := config.ENV.ApiPort
	fmt.Println(port)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World !")
	})

	r.HandleFunc("/auth/token", middleware.AuthMiddleware(http.HandlerFunc(auth.CreateToken))).Methods("GET")

	r.HandleFunc("/notes", middleware.AuthMiddleware(http.HandlerFunc(handler.GetNotesHandler))).Methods("GET")
	r.HandleFunc("/notes", middleware.AuthMiddleware(http.HandlerFunc(handler.PostTaskHandler))).Methods("POST")

	r.HandleFunc("/register", handler.Register).Methods("POST")

	fmt.Println("Server launched on port: " + strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), r)
}
