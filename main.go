package main

import (
	"fmt"
	"net/http"

	config "dauqu.com/github/config"
	routes "dauqu.com/github/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	app := mux.NewRouter()
	config.ConnectDB()

	// app.HandleFunc("/", RedirectHandler)
	app.HandleFunc("/api/register", routes.Register).Methods("POST")
	app.HandleFunc("/api/login", routes.Login).Methods("POST")
	app.HandleFunc("/api/get-token", routes.GetToken).Methods("POST")
	app.HandleFunc("/api/github", routes.Github).Methods("POST")

	//Allow CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	fmt.Println("Server started on port http://localhost:8000")
	http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(app))
}

