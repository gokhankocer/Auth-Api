package main

import (
	"log"

	"github.com/gokhankocer/User-Api/database"
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
	//database.DB.Migrator().DropTable(&entities.User{})
	//database.DB.Migrator().CreateTable(&entities.User{})
	//database.DB.Migrator().AddColumn(&entities.User{}, "Is_Active")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Print("Env successfully loaded")
	}
}
