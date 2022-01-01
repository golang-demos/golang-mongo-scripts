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

	"github.com/golang-demos/chalk"
	"github.com/joho/godotenv"
)

func main() {

	argsLength := len(os.Args)

	appMode := "dev"
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

	// If option is provided through command line
	singularOption := "-1"
	if argsLength > 1 {
		singularOption = os.Args[1]
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
			fmt.Printf("\n %2d : %s", index+1, dScript.Name)
			index++
		}
		fmt.Printf("\n %2d : %s\n\n", 0, "Exit")

		if singularOption == "-1" {
			fmt.Print("Enter Option : ")
			fmt.Scan(&input)
		} else {
			input = singularOption
		}

		// convert input in number
		inputIndex, err := strconv.Atoi(input)
		if err != nil {
			log.Print(chalk.Red(), "Please enter a number between 0 to ", (index - 1), chalk.Reset())
			continue
		}

		// Exit if 0 is entered
		if inputIndex == 0 {
			fmt.Print("\nExiting...")
			break
		}

		// Detect wrong input number. Reiterate if found.
		if len(demoScripts) < inputIndex || inputIndex < 0 {
			fmt.Print("\n", chalk.Red(), "Wrong Input. Try again", chalk.Reset())
			continue
		}

		// Execute selected Handler
		demoScripts[inputIndex-1].Handler()

		if singularOption != "-1" {
			break
		}
	}

}
