package scripts

import "github.com/golang-demos/golang-mongo-scripts/models"

func DemoScript_Insert_Demo() *models.DemoScript {
	return &models.DemoScript{
		Name:        "Insert Demo",
		Description: "Demo for inserting documents into database",
		Handler: func() {

		},
	}
}
