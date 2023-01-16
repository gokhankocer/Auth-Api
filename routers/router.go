package routers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gokhankocer/User-Api/handlers"
	"github.com/gokhankocer/User-Api/middleware"
)

func Setup() {
	router := gin.Default()
	publicRoutes := router.Group("/auth")
	//router.POST("/signup", handlers.Signup)
	publicRoutes.POST("/login", handlers.Login)
	publicRoutes.GET("/logout", middleware.JWTAuthMiddleware(), handlers.Logout)
	publicRoutes.GET("/authorize", handlers.Authorize)
	log.Fatal(router.Run("localhost:8880"))
}
