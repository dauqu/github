package routes

import (
	"encoding/json"
	"fmt"
	"github.com/go-git/go-git/v5"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func CloneRepos(w http.ResponseWriter, r *http.Request) {

	type Body struct {
		Tag         string `json:"tag"`
		CloneUrl    string `json:"clone_url"`
		RemoteName  string `json:"remote_name"`
		AccessToken string `json:"access_token"`
		Username    string `json:"username"`
	}

	//Get body
	body_data, _ := ioutil.ReadAll(r.Body)
	//Body to JSON
	var bodyData Body
	json.Unmarshal(body_data, &bodyData)

	//Get current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bodyData)

	path_name := dir + "/repos/" + bodyData.Tag

	re := gitCloneURL(bodyData.CloneUrl, bodyData.AccessToken, bodyData.Username)

	fmt.Println(re)

	//Creae path
	err = os.MkdirAll(path_name, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	//Clone private repos
	resp, err := git.PlainClone(path_name, false, &git.CloneOptions{
		URL:      re,
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println(err)
	}

	//Read response from clone
	fmt.Println(resp)
}

func gitCloneURL(url, token, username string) string {
	url = strings.Replace(url, "https://", "", 1)
	return "https://" + username + ":" + token + "@" + url
}
