package helper

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gokhankocer/User-Api/database"
	"github.com/gokhankocer/User-Api/entities"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwt(user entities.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET")))
}
func ValidateJWT(c *gin.Context) error {
	token, err := GetToken(c)
	if err != nil {
		return err
	}

	if isBlacklisted(c, token) {
		return errors.New("invalid token provided")
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}

	return errors.New("invalid token provided")
}

func CurrentUser(c *gin.Context) (entities.User, error) {
	err := ValidateJWT(c)
	if err != nil {
		fmt.Println(err, "error1")
		return entities.User{}, err
	}
	token, err := GetToken(c)
	if err != nil {
		fmt.Println(err, "error2")
	}
	claims, a := token.Claims.(jwt.MapClaims)

	fmt.Println(a, "error3")

	userId := uint(claims["user_id"].(float64))
	user, err := entities.FindUserById(userId)
	if err != nil {
		fmt.Println(err, "error4")
		return entities.User{}, err

	}
	fmt.Println(user.ID, "user id")
	return user, nil
}
func GetToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		fmt.Println(os.Getenv("SECRET"))

		return []byte(os.Getenv("SECRET")), nil
	})
	return token, err
}

func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}

func isBlacklisted(c *gin.Context, token *jwt.Token) bool {

	res, err := database.RDB.Exists(c, token.Raw).Result()
	if err != nil {
		panic(err)
	}
	if res == 1 {
		return true
	}
	return false

}
