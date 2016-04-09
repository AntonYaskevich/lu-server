package users

import (
	"time"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/AntonYaskevich/lu-server/models"
	"github.com/AntonYaskevich/lu-server/repository"
	"github.com/AntonYaskevich/lu-server/repository/userdb"
)

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateJWTToken(id string, key string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["id"] = id
	token.Claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	tokenString, err := token.SignedString([]byte(key))
	return tokenString, err
}

func Login(c *gin.Context) {
	data := models.Login{}
	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	repository := userdb.UserDB{repository.New()}

	user, err := repository.ReadByUsername(data.Username)
	if err != nil || user == nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "The username or password doesn't match",
		})
		return
	}
	token, err := CreateJWTToken(strconv.FormatUint(user.Id, 16), "key")
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"Authorization": "Bearer " + token,
	})
}
