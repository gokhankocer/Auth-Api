package main

import (
	"log"

	"github.com/gokhankocer/User-Api/database"
	"github.com/gokhankocer/User-Api/entities"
	"github.com/gokhankocer/User-Api/routers"
	"github.com/joho/godotenv"
)

func main() {
	loadDatabase()
	loadEnv()
	routers.Setup()
}
func loadDatabase() {
	database.Connect()
	database.ConnectRedis()
	database.DB.Migrator().CreateTable(&entities.User{})
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Print("Env successfully loaded")
	}
}
