package handlers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateJWTToken(id string, key string) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)
	token.Claims["id"] = id
	token.Claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	tokenString, err := token.SignedString([]byte(key))
	return tokenString, err
}
