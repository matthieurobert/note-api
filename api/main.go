package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matthieurobert/amp/api/config"
	"github.com/matthieurobert/amp/api/entity"
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

	fmt.Println("Server launched on port: " + strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), r)
}
