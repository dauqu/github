package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func GetToken(w http.ResponseWriter, r *http.Request) {

	type Body struct {
		Username string `json:"username"`
	}

	//Bind JSOn
	var body Body
	json.NewDecoder(r.Body).Decode(&body)

	//Create context
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	//Get Token from username
	var github bson.M
	err := GitCollection.FindOne(ctx, bson.M{"username": body.Username}).Decode(&github)
	if err != nil {
		fmt.Println(err)
	}

	//Get access token
	access_token := github["access_token"]
	access_type := github["token_type"]

	//ReturnJSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"access_token": access_token, "token_type": access_type})
}
