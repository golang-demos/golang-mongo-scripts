package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber"
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

	seeder.TruncateDatabase()
	seeder.SeedDatabase()

	time.Sleep(2 * time.Second)

	app := fiber.New()

	scripts.RegisterRoutes(app)

	apiServerPort := os.Getenv("MONGO_APP_PORT")
	log.Fatal(app.Listen(apiServerPort))
}
