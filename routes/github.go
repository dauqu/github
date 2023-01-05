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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Github(w http.ResponseWriter, r *http.Request) {
	//Print body
	body, _ := ioutil.ReadAll(r.Body)

	type Response struct {
		Code string `json:"code"`
	}

	var response Response

	json.Unmarshal(body, &response)

	res, err := GetAccessToken(response.Code)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	//ReturnJSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

var GitCollection *mongo.Collection = config.GetCollection(config.DB, "github")

func GetAccessToken(code string) (container_id string, err error) {

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
		"username":      "harsha",
		"created_at":    time.Now().Format("2006-01-02 15:04:05"),
		"updated_at":    time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
