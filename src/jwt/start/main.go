package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// Define a secret key for signing the token
	secretKey := []byte("your-256-bit-secret")
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["foo"] = "bar"
	claims["exp"] = jwt.NewNumericDate(time.Now().Add(time.Hour))

	// Sign the token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return
	}

	fmt.Println("Generated Token:", tokenString)
}
