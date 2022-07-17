package main

import (
	"os"
	"wallet-tracker-api/database"
	"wallet-tracker-api/router"

	"github.com/joho/godotenv"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger

// @contact.name nathan.lu
// @contact.url https://tedmax100.github.io/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// schemes http
func main() {
	godotenv.Load() // load .env
	database.ConnectDB()
	router := router.SetupRouter()      // load /router/*
	router.Run(":" + os.Getenv("PORT")) // run server in :PORT( env var PORT )
}
