package scripts

import (
	"fmt"

	"github.com/golang-demos/golang-mongo-scripts/models"
	"github.com/ttacon/chalk"
)

var DemoScripts []models.DemoScript

func RegisterScripts() []*models.DemoScript {
	demoScripts := []*models.DemoScript{
		DemoScript_Insert_Demo(),
		DemoScript_Conditions_Demo(),
	}

	return demoScripts
}

func PrintTaskNote(note string) {
	fmt.Println(chalk.Green, "+", note, chalk.Reset)
}
