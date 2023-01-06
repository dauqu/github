package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Createauth(w http.ResponseWriter, r *http.Request) {

	//Get current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	//Read pen file
	pem, err := ioutil.ReadFile(dir + "/routes/dauqu.2023-01-06.private-key.pem")
	if err != nil {
		fmt.Println(err)
	}

	//Get current time
	time := time.Now().Unix()

	// Load the private key
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(pem))
	if err != nil {
		panic(err)
	}

	// Create a new token
	token := jwt.New(jwt.SigningMethodRS256)

	// Set some claims
	token.Claims = jwt.MapClaims{
		"iss": 277956,
		"iat": time,
		//Exp in 10 minutes
		"exp": time + 600,
	}

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(tokenString)

}
