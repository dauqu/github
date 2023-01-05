package main

import (
	"net/http"

	routes "dauqu.com/github/routes"
	config "dauqu.com/github/config"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	app := mux.NewRouter()
	config.ConnectDB()

	// app.HandleFunc("/", RedirectHandler)
	app.HandleFunc("/api/github", routes.Github).Methods("POST")

	//Allow CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(app))
}

// func RedirectHandler(w http.ResponseWriter, r *http.Request) {

// 	//Get code from query
// 	code := r.URL.Query().Get("code")

// 	//Print code
// 	fmt.Println(code)

// 	//Get access token
// 	GetAccessToken(code)

// }

// func GetAccessToken(code string) error {

// 	// Set the request parameters
// 	CREATE_ID := "Iv1.72f299b0ba45be0a"
// 	CREATE_SECRET := "4273f4bd8116e6865fb47688b6e1cd1dee14fe88"
// 	REDIRECT_URI := "http://localhost:8000"

// 	URL := "https://github.com/login/oauth/access_token" + "?client_id=" + CREATE_ID + "&client_secret=" + CREATE_SECRET + "&code=" + code + "&redirect_uri=" + REDIRECT_URI

// 	// Create the POST request
// 	req, err := http.NewRequest("POST", URL, nil)
// 	if err != nil {
// 		return err
// 	}

// 	// Set the request header
// 	req.Header.Set("Content-Type", "application/json")
// 	//Accept: application/json
// 	req.Header.Set("Accept", "application/json")

// 	// Send the request and get the response
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	//Log Body in JSON
// 	body, _ := ioutil.ReadAll(resp.Body)

// 	type Response struct {
// 		AccessToken  string `json:"access_token"`
// 		RefreshToken string `json:"refresh_token"`
// 		Scope        string `json:"scope"`
// 		ExpiresIn    int    `json:"expires_in"`
// 	}

// 	var response Response

// 	json.Unmarshal(body, &response)

// 	fmt.Println(response.AccessToken)

// 	return nil
// }
