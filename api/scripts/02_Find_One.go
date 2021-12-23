package scripts

import "github.com/golang-demos/golang-mongo-scripts/models"

func DemoScript_FindOne_Demo() *models.DemoScript {
	return &models.DemoScript{
		Name:        "Find One",
		Description: "Demo scripts for variations using FindOne() function",
		Handler: func() {

		},
	}
}
