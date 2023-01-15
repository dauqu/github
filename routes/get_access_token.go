package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetAccessToken(w http.ResponseWriter, r *http.Request) {
	
	type Body struct {
		InstallationId string `json:"installation_id"`
	}

	//Get body
	body_data, _ := ioutil.ReadAll(r.Body)
	//Body to JSON

	var bodyData Body
	json.Unmarshal(body_data, &bodyData)

	id := bodyData.InstallationId

	//Get current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	//Read token.txt file
	token, err := ioutil.ReadFile(dir + "/routes/token.txt")
	if err != nil {
		fmt.Println(err)
	}

	//Convert token to string
	res := string(token)

	//Post request to get access token
	URL := "https://api.github.com/app/installations/" + id + "/access_tokens"

	//Create request
	req, err := http.NewRequest("POST", URL, nil)
	if err != nil {
		fmt.Println(err)
	}
	//Authorization
	req.Header.Set("Authorization", "Bearer "+res)
	// Set the request header
	req.Header.Set("Content-Type", "application/json")
	//Accept: application/json
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	//Read response
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	type Response struct {
		Token string `json:"token"`
	}

	var response Response

	json.Unmarshal(responseData, &response)

	//Write cookie
	cookie := http.Cookie{Name: "token", Value: response.Token}
	http.SetCookie(w, &cookie)

	//Write response
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}