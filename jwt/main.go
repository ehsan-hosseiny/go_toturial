package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey []byte

func main() {

	secretKey = []byte("ehsan")

	claims := jwt.MapClaims{
		"user_id":   12,
		"expire_at": time.Now().Add(time.Hour * 4).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(tokenString)
	validateToken(tokenString)

}

func validateToken(tokenParam string) {
	token, err := jwt.Parse(tokenParam, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}	
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims["user_id"])
		fmt.Println("token is valid")
	} else {
		fmt.Println("token is not valid")
	}

}
