package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gokhankocer/User-Api/database"
	"github.com/gokhankocer/User-Api/entities"
	"github.com/gokhankocer/User-Api/helper"
	"github.com/gokhankocer/User-Api/models"
)

func Login(c *gin.Context) {
	var body models.UserLoginRequest
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	user, err := entities.FindUserByEmail(body.Email)
	if err != nil {
		log.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email"})
		return
	}

	if !user.IsActive {
		log.Println("error", "User Error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is not active"})
		return
	}

	if user.VerifyPassword(body.Password) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password"})
		return
	}

	jwt, err := helper.GenerateJwt(user)
	if err != nil {
		log.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Create Token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwt})
	fmt.Println("Successfully loged in")
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
