package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetToken(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("token")
	if err != nil {
		fmt.Println(err)
	}

	//Verify token
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
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

	//Create context
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	var user bson.M
	err = GitCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "User not found"})
		return
	}
	if err != nil {
		return // err
	}

	//Get request http
	url := "https://api.github.com/user/repos"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+user["access_token"].(string))

	//Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return // err
	}

	//Read response
	body, _ := ioutil.ReadAll(resp.Body)

	//Convert body to JSON
	var repos []interface{}
	json.Unmarshal(body, &repos)

	//Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repos)
}
