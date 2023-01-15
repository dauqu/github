package routes

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"os"
	"time"
)

func AppAuth(key_path string) error {

	//Get current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	//Get current time
	time := time.Now().Unix()

	// Load the private key
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(key_path))
	if err != nil {
		panic(err)
	}

	// Create a new token
	token := jwt.New(jwt.SigningMethodRS256)

	// Set some claims
	token.Claims = jwt.MapClaims{
		"iss": 277956,
		"iat": time,
		"exp": time + 600,
	}

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}

	//Store token in a file
	err = ioutil.WriteFile(dir+"/routes/token.txt", []byte(tokenString), 0644)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
