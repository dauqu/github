package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func InstalledApps(w http.ResponseWriter, r *http.Request) {
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

	//Get all installed apps by username
	cursor, err := AppsCollection.Find(context.Background(), filter)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	//Convert cursor to array
	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}
