package scripts

import (
	"github.com/golang-demos/golang-mongo-scripts/models"
)

var DemoScripts []models.DemoScript

func RegisterScripts() []*models.DemoScript {
	demoScripts := []*models.DemoScript{
		DemoScript_Insert_Demo(),
		DemoScript_FindOne_Demo(),
	}

	return demoScripts
}
