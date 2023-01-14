package main

import (
	"fmt"
	"net/http"
	"time"

	config "dauqu.com/github/config"
	routes "dauqu.com/github/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	app := mux.NewRouter()
	config.ConnectDB()

	app.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Set cookie
		cookie := http.Cookie{
			Name:    "token",
			Value:   "123",
			Expires: time.Now().Add(24 * time.Hour),
		}

		http.SetCookie(w, &cookie)
		fmt.Fprintf(w, "Hello World")
	})

	// app.HandleFunc("/", RedirectHandler)
	app.HandleFunc("/api/register", routes.Register).Methods("POST")
	app.HandleFunc("/api/login", routes.Login).Methods("POST")
	app.HandleFunc("/api/is-logged-in", routes.IsLoggedIn).Methods("GET")
	app.HandleFunc("/api/get-repos", routes.GetToken).Methods("GET")
	app.HandleFunc("/api/github", routes.Github).Methods("POST")
	app.HandleFunc("/api/connect-github", routes.ConnectGithub).Methods("POST")
	app.HandleFunc("/api/create-auth", routes.Createauth).Methods("GET")
	app.HandleFunc("/api/get-my-repos", routes.GetMyRepos).Methods("GET")

	//Allow CORS
	credentialsOk := handlers.AllowCredentials()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Accept", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000", "https://github-orpin.vercel.app"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	fmt.Println("Server started on port http://localhost:8000")
	http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk, credentialsOk)(app))
}
