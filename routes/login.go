package routes

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"

	//JWT
	"github.com/golang-jwt/jwt/v4"
)

func Login(w http.ResponseWriter, r *http.Request) {

	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//Bind JSOn
	var body Body
	json.NewDecoder(r.Body).Decode(&body)

	//Create context
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	//Get Token from username
	resp, err := UsersCollection.FindOne(ctx, bson.M{"email": body.Email}).DecodeBytes()
	if err != nil {
		return
	}

	hashedPassword := resp.Lookup("password").StringValue()
	username := resp.Lookup("username").StringValue()
	email := resp.Lookup("email").StringValue()

	//Check if password is correct
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(body.Password))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid password")
		return
	}

	//Create token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return // err
	}

	//Set cookie
	cookie := http.Cookie{
		Name:    "token",
		Value:   t,
		Expires: time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, &cookie)

	//Return token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Return token with message and token
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged in successfully", "token": t})
}
