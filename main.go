package main

import (
	"log"
	"subway/server"
	"subway/server/db"

	"os"

	"github.com/joho/godotenv"
)

// @title Subway surfers api
// @version 1.0
// @description This is the api doucmentation for subway surfers
// @host localhost:3000
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connection := db.InitDB()
	db.Transfer(connection)

	app := server.NewServer(connection)
	server.ConfigureRoutes(app)

	if err := app.Run(os.Getenv("PORT")); err != nil {
		log.Print(err)
	}
}
