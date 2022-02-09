package main

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func main() {
	hmacSampleSecret := []byte("secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": "juancarlos.roman@globant.com",
	})

	tokenString, _ := token.SignedString(hmacSampleSecret)

	fmt.Println(tokenString)
}
