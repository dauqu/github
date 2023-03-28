package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

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
	if err == mongo.ErrNoDocuments {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "User not found"})
		return
	}
	if err != nil {
		return // err
	}

	hashedPassword := resp.Lookup("password").StringValue()
	username := resp.Lookup("username").StringValue()
	email := resp.Lookup("email").StringValue()

	//Check if password is correct
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(body.Password))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Incorrect password"})
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
		SameSite: http.SameSiteNoneMode,
		Secure: true,
		Expires: time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, &cookie)

	//Return token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Return token with message and token
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged in successfully", "token": t})
}

// Check if user is logged in
func IsLoggedIn(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Unauthorized"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Bad request " + err.Error()})
		return
	}

	tokenString := cookie.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte("secret"), nil
	})

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "Unauthorized"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Authorized", "username": claims["username"].(string), "email": claims["email"].(string)})
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		json.NewEncoder(w).Encode(map[string]string{"message": "Unauthorized"})
		return
	}
}
