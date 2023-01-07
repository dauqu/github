package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	actions "dauqu.com/github/actions"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func GetMyRepos(w http.ResponseWriter, r *http.Request) {

	//Get cookies from request
	cookie, err := r.Cookie("token") 
	if err != nil {
		fmt.Println(err)
	}

	//Check if token haven't value
	if cookie.Value == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid token"})
		return
	}

	//Verify token
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte("secret"), nil
	})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid token"})
		return
	}

	//Get usernae and email from token
	username := token.Claims.(jwt.MapClaims)["username"]
	//Convert username to string
	user_name := fmt.Sprintf("%v", username)

	//Find by username
	filter := bson.M{"username": user_name}
	var result bson.M
	err = AppsCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "App not installed"})
		return
	}

	//Get installation id
	installation_id := result["installation_id"]
	//Convert installation id to string
	installation_id_string := fmt.Sprintf("%v", installation_id)

	res, err := actions.GetToken()
	if err != nil {
		fmt.Println(err)
	}

	//Post request to get access token
	URL := "https://api.github.com/app/installations/" + installation_id_string + "/access_tokens"

	//Create request
	req, err := http.NewRequest("POST", URL, nil)
	if err != nil {
		fmt.Println(err)
	}
	//Authorization
	req.Header.Set("Authorization", "Bearer "+res)
	// Set the request header
	req.Header.Set("Content-Type", "application/json")
	//Accept: application/json
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//Log Body in JSON
	body, _ := ioutil.ReadAll(resp.Body)

	type Response struct {
		Token string `json:"token"`
	}

	var response Response
	json.Unmarshal(body, &response)

	//Get access token
	access_token := response.Token

	fmt.Println(access_token)

	//Get repos
	URL = "https://api.github.com/installation/repositories?per_page=100&visibility=all"

	//Create request
	req, err = http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println(err)
	}
	//Authorization
	req.Header.Set("Authorization", "token "+access_token)
	// Set the request header
	req.Header.Set("Content-Type", "application/json")
	//Accept: application/json
	req.Header.Set("Accept", "application/json")

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	//Log Body in JSON
	body, _ = ioutil.ReadAll(resp.Body)
	//Body to JSOn
	var data interface{}
	json.Unmarshal(body, &data)


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
