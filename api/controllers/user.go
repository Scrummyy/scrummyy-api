package controllers

import (
	"github.com/Scrummyy/scrummyy-api/data/models"
	datatype "github.com/Scrummyy/scrummyy-api/internal/datatypes"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type UserAPI struct {
	config *viper.Viper
}

// UserController ...
type UserController struct{}

var userModel = new(models.UserModel)
var userForm = new(datatype.UserForm)

func RegisterUserHandler(r gin.IRouter, conf *viper.Viper) {
	api := UserAPI{config: conf}
	r.POST("/users/login", api.Login)
	r.POST("/users/register", api.Register)
	r.GET("/users/logout", api.Logout)
}

// getUserID ...
func getUserID(c *gin.Context) (userID int64) {
	//MustGet returns the value for the given key if it exists, otherwise it panics.
	return c.MustGet("userID").(int64)
}

// Login ...
func (ctrl UserAPI) Login(c *gin.Context) {
	var loginForm datatype.LoginForm

	if validationErr := c.ShouldBindJSON(&loginForm); validationErr != nil {
		message := userForm.Login(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	user, token, err := userModel.Login(loginForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid login details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in", "user": user, "token": token})
}

// Register ...
func (ctrl UserAPI) Register(c *gin.Context) {
	var registerForm datatype.RegisterForm

	if validationErr := c.ShouldBindJSON(&registerForm); validationErr != nil {
		message := userForm.Register(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	user, err := userModel.Register(registerForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered", "user": user})
}

// Logout ...
func (ctrl UserAPI) Logout(c *gin.Context) {

	au, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User not logged in"})
		return
	}

	deleted, delErr := authModel.DeleteAuth(au.AccessUUID)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
