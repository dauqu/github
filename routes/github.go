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
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var GitCollection *mongo.Collection = config.GetCollection(config.DB, "github")

func Github(w http.ResponseWriter, r *http.Request) {

	//Create context
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	//Print body
	body, _ := ioutil.ReadAll(r.Body)

	type Response struct {
		Code string `json:"code"`
	}

	var response Response
	json.Unmarshal(body, &response)

	fmt.Println(response.Code)

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

	//Check username exists
	var user models.User
	err = GitCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		res, err := GetAccessToken(response.Code, user_name)
		if err != nil {
			fmt.Println(err)
		}

		//ReturnJSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
		return
	}
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)

	//if user exists
	if user.Username == user_name {

		res, err := UpdateAccessToken(response.Code, user_name)
		if err != nil {
			fmt.Println(err)
		}

		//ReturnJSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
		return
	}
}

func UpdateAccessToken(code string, username string) (container_id string, err error) {

	// Set the request parameters
	// CREATE_ID := "Iv1.72f299b0ba45be0a"
	CLIENT_ID := "Iv1.72f299b0ba45be0a"
	CREATE_SECRET := "4fef0590e0a26a527a0d159c2bae5fe690bdbb02"
	REDIRECT_URI := "http://localhost:3000/gitcode"

	URL := "https://github.com/login/oauth/access_token" + "?client_id=" + CLIENT_ID + "&client_secret=" + CREATE_SECRET + "&code=" + code + "&redirect_uri=" + REDIRECT_URI

	// Create the POST request
	req, err := http.NewRequest("POST", URL, nil)
	if err != nil {
		return "", err
	}

	// Set the request header
	req.Header.Set("Content-Type", "application/json")
	//Accept: application/json
	req.Header.Set("Accept", "application/json")

	// Send the request and get the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	//Log Body in JSON
	body, _ := ioutil.ReadAll(resp.Body)

	type Response struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
		TokenType    string `json:"token_type"`
	}

	var response Response

	fmt.Println(string(body))

	json.Unmarshal(body, &response)

	//Create context
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	//Update by username
	filter := bson.M{"username": username}

	//Update by username
	update := bson.M{"$set": bson.M{"access_token": response.AccessToken, "refresh_token": response.RefreshToken, "expires_in": response.ExpiresIn, "token_type": response.TokenType}}

	//Update by username
	_, err = GitCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}

	return "Updated", nil
}

func GetAccessToken(code string, username string) (container_id string, err error) {

	// Set the request parameters
	CREATE_ID := "Iv1.72f299b0ba45be0a"
	CREATE_SECRET := "4273f4bd8116e6865fb47688b6e1cd1dee14fe88"
	REDIRECT_URI := "http://localhost:3000/gitcode"

	URL := "https://github.com/login/oauth/access_token" + "?client_id=" + CREATE_ID + "&client_secret=" + CREATE_SECRET + "&code=" + code + "&redirect_uri=" + REDIRECT_URI

	// Create the POST request
	req, err := http.NewRequest("POST", URL, nil)
	if err != nil {
		return "", err
	}

	// Set the request header
	req.Header.Set("Content-Type", "application/json")
	//Accept: application/json
	req.Header.Set("Accept", "application/json")

	// Send the request and get the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	//Log Body in JSON
	body, _ := ioutil.ReadAll(resp.Body)

	type Response struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
		TokenType    string `json:"token_type"`
	}

	var response Response

	// fmt.Println(string(body))

	json.Unmarshal(body, &response)

	//Create context
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	//Create new user
	newUser := models.Github{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
		ExpiresIn:    response.ExpiresIn,
		Type:         response.TokenType,
	}

	//Insert new user
	result, err := GitCollection.InsertOne(ctx, bson.M{
		"access_token":  newUser.AccessToken,
		"refresh_token": newUser.RefreshToken,
		"expires_in":    newUser.ExpiresIn,
		"token_type":    newUser.Type,
		"username":      username,
		"created_at":    time.Now().Format("2006-01-02 15:04:05"),
		"updated_at":    time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
