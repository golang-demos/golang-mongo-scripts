package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-demos/golang-mongo-scripts/database"
	"github.com/golang-demos/golang-mongo-scripts/models"
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

	var demoScripts []*models.DemoScript = scripts.RegisterScripts()

	var index int16
	var input string

	for {
		index = 0
		for _, dScript := range demoScripts {
			fmt.Printf("\n[%d] %s", index+1, dScript.Name)
			index++
		}
		fmt.Printf("\n[%d] %s", 0, "Exit")
		fmt.Print("\n\nOption : ")
		fmt.Scan(&input)

		inputIndex, _ := strconv.Atoi(input)
		if inputIndex == 0 {
			fmt.Print("\nExiting...")
			break
		}

		if len(demoScripts) < inputIndex {
			fmt.Print("\nWrong Input. Try again")
			continue
		}
	}

}
