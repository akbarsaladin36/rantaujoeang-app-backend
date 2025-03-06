package main

import (
	"fmt"
	"rantaujoeang-app-backend/database"
	"rantaujoeang-app-backend/migrations"
	"rantaujoeang-app-backend/router"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	database.ConnectDB()
	migrations.MigrateTables()
	router.ConnectRoutes()
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Failed to load .env file")
	}
}
