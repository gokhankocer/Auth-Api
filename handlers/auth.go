package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gokhankocer/User-Api/database"
	"github.com/gokhankocer/User-Api/entities"
	"github.com/gokhankocer/User-Api/helper"
	"github.com/gokhankocer/User-Api/models"
)

func Login(c *gin.Context) {

	var body models.UserRequest
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	user, err := entities.FindUserByName(body.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Name"})
		return
	}
	if user.Name != body.Name {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Name"})
		return
	}

	password := user.VerifyPassword(body.Password)
	if password != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password"})
		return
	}

	jwt, err := helper.GenerateJwt(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Create Token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": jwt})

}

func Logout(c *gin.Context) {

	token, _ := helper.GetToken(c)

	err := database.RDB.Set(c, token.Raw, 1, 0).Err()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Loged Out"})

}

func Authorize(c *gin.Context) {
	user, err := helper.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Not Authorized"})
		return
	}
	c.JSON(http.StatusOK, user)
}
