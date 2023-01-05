package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

func GetToken(w http.ResponseWriter, r *http.Request) {

	type Body struct {
		Token string `json "token"`
	}

	var body Body
	json.NewDecoder(r.Body).Decode(&body)

	cookie, err := r.Cookie("token")
	if err != nil {
		panic(err)
	}

	//Parse token
	token, err := jwt.Parse(cookie.Name, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte("secret"), nil
	})
	if err != nil {
		panic(err)
	}

	//Get username from token
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"]


	//Create context
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	//Get Token from username
	resp, err := UsersCollection.FindOne(ctx, bson.M{"username": username}).DecodeBytes()
	if err != nil {
		return
	}

	//Get access token
	access_token := resp.Lookup("access_token").StringValue()
	access_type := github["token_type"]

	//ReturnJSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"access_token": access_token, "token_type": access_type})
}
