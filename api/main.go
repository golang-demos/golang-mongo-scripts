package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-demos/golang-mongo-scripts/database"
	"github.com/golang-demos/golang-mongo-scripts/scripts"
	"github.com/golang-demos/golang-mongo-scripts/seeder"

	"github.com/joho/godotenv"
)

func main() {

	appMode := "dev"
	argsLength := len(os.Args)
	if argsLength > 0 {
		if os.Args[0] == "prod" {
			appMode = "prod"
		}
	}

	envFile := ".env.dev"
	if appMode == "prod" {
		envFile = ".env"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()
	defer database.DisconnectDB()

	seeder.SeedDatabase(false)

	app := fiber.New()

	scripts.RegisterRoutes(app)

	apiServerPort := os.Getenv("MONGO_APP_PORT")
	log.Fatal(app.Listen(apiServerPort))
}
