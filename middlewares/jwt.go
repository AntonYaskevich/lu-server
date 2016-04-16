package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	ID = "_id"
)

func Auth(key string) *gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := jwt.ParseFromRequest(c.Request, func(token *jwt.Token) (interface{}, error) {
			privateKey := ([]byte(key))
			return privateKey, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		userId := token.Claims[ID]
		c.Set(ID, userId)
		c.Next()
	}
}