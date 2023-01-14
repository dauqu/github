package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"dauqu.com/github/config"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var AppsCollection *mongo.Collection = config.GetCollection(config.DB, "installed_apps")

func InstallApp(w http.ResponseWriter, r *http.Request) {
	//Read body
	body, _ := ioutil.ReadAll(r.Body)

	//Create context
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	type Body struct {
		InstallationID string `json:"installation_id"`
	}

	var response Body
	json.Unmarshal(body, &response)

	//Get cookies from request
	cookie, err := r.Cookie("token")
	if err != nil {
		fmt.Println(err)
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

	//Get installation id from request
	installation_id := response.InstallationID

	//Check if app is already installed
	result, err := AppsCollection.FindOne(ctx, bson.M{"installation_id": installation_id}).DecodeBytes()
	if err != nil {
		fmt.Println(err)
	}

	//Get installation id from database
	installation_id_db := result.Lookup("installation_id").StringValue()

	fmt.Println(installation_id_db)
	fmt.Println(installation_id)

	//Check if installation id is already in database
	if installation_id == installation_id_db {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "App is already installed"})
		return
	}

	//Create a new document
	_, err = AppsCollection.InsertOne(ctx, bson.M{
		"installation_id": installation_id,
		"username":        user_name,
		"created_at":      time.Now().Format("2006-01-02 15:04:05"),
		"updated_at":      time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		fmt.Println(err)
	}

	//Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Successfully connected to github"})

}
