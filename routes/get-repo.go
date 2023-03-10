package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetRepoById(w http.ResponseWriter, r *http.Request) {

	type Body struct {
		InstallationId string `json:"installation_id"`
	}

	//Get body
	body_data, _ := ioutil.ReadAll(r.Body)
	//Body to JSON

	var bodyData Body
	json.Unmarshal(body_data, &bodyData)

	id := bodyData.InstallationId

	//Get installation id

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
	defer resp.Body.Close()

	//Log Body in JSON
	body, _ := ioutil.ReadAll(resp.Body)

	type Response struct {
		Token string `json:"token"`
	}

	var response Response
	json.Unmarshal(body, &response)

	//Get access token
	access_token := response.Token

	//Get repos
	URL = "https://api.github.com/installation/repositories?per_page=100&visibility=all"

	//Create request
	req, err = http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println(err)
	}
	//Authorization
	req.Header.Set("Authorization", "token "+access_token)
	// Set the request header
	req.Header.Set("Content-Type", "application/json")
	//Accept: application/json
	req.Header.Set("Accept", "application/json")

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	//Log Body in JSON
	body, _ = ioutil.ReadAll(resp.Body)
	//Body to JSOn
	var data interface{}
	json.Unmarshal(body, &data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": data, "token": token})
	//Return
}
