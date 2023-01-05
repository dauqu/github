package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	app := mux.NewRouter()

    app.HandleFunc("/redirect", RedirectHandler).Methods("POST")

    http.ListenAndServe(":8000", app)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
    //Print the request body
    fmt.Println(r.Body)
}
