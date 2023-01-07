package actions

import (
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"os"
	"time"
)

func GetToken() (token string, err error) {
	//Get current working directory
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	//Read pen file
	pem, err := ioutil.ReadFile(dir + "/routes/dauqu.2023-01-06.private-key.pem")
	if err != nil {
		return "", err
	}

	//Get current time
	time := time.Now().Unix()

	// Load the private key
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(pem))
	if err != nil {
		return "", err
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"iss": 277956,
		"iat": time,
		"exp": time + 600,
	}

	// Create the token
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
