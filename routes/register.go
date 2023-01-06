package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"dauqu.com/github/config"
	models "dauqu.com/github/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UsersCollection *mongo.Collection = config.GetCollection(config.DB, "users")

func Register(w http.ResponseWriter, r *http.Request) {

	//Create context
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	//Print body
	read_body, _ := ioutil.ReadAll(r.Body)

	var body models.User
	json.Unmarshal(read_body, &body)

	//Check all fields
	if body.FullName == "" || body.Email == "" || body.Username == "" || body.Password == "" {
		//Return error response JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "All fields are required"})
		return
	}

	//Check username exists
	var username models.User
	err := UsersCollection.FindOne(ctx, bson.M{"username": body.Username}).Decode(&username)
	if err != mongo.ErrNoDocuments {
		//Return error response JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Username already exists"})
		return
	}

	//Check if email exists
	var email models.User
	err = UsersCollection.FindOne(ctx, bson.M{"email": body.Email}).Decode(&email)
	if err != mongo.ErrNoDocuments {
		//Return error response JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Email already exists"})
		return
	}

	//Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 8)
	if err != nil {
		fmt.Println(err)
	}

	//Insert user
	result, err := UsersCollection.InsertOne(ctx, bson.M{
		"full_name":   body.FullName,
		"email":       body.Email,
		"username":    body.Username,
		"phone":       body.Phone,
		"license_key": "",
		"password":    string(hashedPassword),
		"created_at":  time.Now().Format("2006-01-02 15:04:05"),
		"updated_at":  time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		fmt.Println(err)
	}

	//ReturnJSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User created", "id": result.InsertedID})
}
