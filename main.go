package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	config "dauqu.com/github/config"
	routes "dauqu.com/github/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

func main() {
	app := mux.NewRouter()

	//Connect to database
	config.ConnectDB()

	//Get current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	//Read pen file
	pem, err := ioutil.ReadFile(dir + "/routes/key.pem")
	if err != nil {
		fmt.Println(err)
	}

	err = routes.AppAuth(string(pem))
	if err != nil {
		fmt.Println(err)
	}

	//Cron job to refresh token
	c := cron.New()
	c.AddFunc("@every 9m", func() {
		fmt.Println("Token refreshed at", time.Now())
		err = routes.AppAuth(string(pem))
		if err != nil {
			fmt.Println(err)
		}
	})
	c.Start()

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
	app.HandleFunc("/api/install-app", routes.InstallApp).Methods("POST")
	app.HandleFunc("/api/get-my-repos", routes.GetMyRepos).Methods("GET")
	//Get Repos by ID
	app.HandleFunc("/api/get-repos", routes.GetRepoById).Methods("POST")
	app.HandleFunc("/api/installed-apps", routes.InstalledApps).Methods("GET")
	//Catch all events from github
	app.HandleFunc("/api/events", routes.Events).Methods("POST")
	//Access token
	app.HandleFunc("/api/access-token", routes.GetAccessToken).Methods("POST")
	//Clone repos
	app.HandleFunc("/api/clone-repos", routes.CloneRepos).Methods("POST")

	//Allow CORS
	credentialsOk := handlers.AllowCredentials()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Accept", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000", "https://github-orpin.vercel.app"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	fmt.Println("Server started on port http://localhost:8000")
	http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk, credentialsOk)(app))
}
