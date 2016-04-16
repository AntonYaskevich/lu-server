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
	"github.com/AntonYaskevich/lu-server/utils"
	"log"
)

var userNotFoundError = utils.ApiError{
	Status: http.StatusNotFound,
	Title:  "User not found",
}

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateJWTToken(id string, key string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["_id"] = id
	token.Claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	tokenString, err := token.SignedString([]byte(key))
	return tokenString, err
}

func Create(c *gin.Context) {
	user := models.User{}
	err := c.Bind(&user)

	if err != nil {
		c.Error(err)
		return
	}
	repository := userdb.UserDB{repository.New()}

	user, err = repository.Create(user);
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, user)
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

func Get(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "me" {
		id = c.MustGet("_id").(string)
	}
	repository := userdb.UserDB{repository.New()}

	user, err := repository.Read(id)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusNotFound, userNotFoundError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetAll(c *gin.Context) {
	repository := userdb.UserDB{repository.New()}

	users, err := repository.GetAll()

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func Update(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.Error(err)
		return
	}
	user.Id = c.Params.ByName("id")

	repository := userdb.UserDB{repository.New()}
	updated, err := repository.Update(user)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, updated)
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	repository := userdb.UserDB{repository.New()}
	err := repository.Delete(id)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, id)
}