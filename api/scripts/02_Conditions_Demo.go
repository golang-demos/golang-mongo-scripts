package scripts

import (
	"strconv"

	"github.com/golang-demos/golang-mongo-scripts/database"
	"github.com/golang-demos/golang-mongo-scripts/models"
	"go.mongodb.org/mongo-driver/bson"
)

func DemoScript_Conditions_Demo() *models.DemoScript {
	return &models.DemoScript{
		Name:        "Conditions Demo",
		Description: "Examples to demonstrate various conditions that can be set while finding or counting documents.",
		Handler: func() {
			var filter bson.D
			var recordsCount int

			// Count all employes
			filter = bson.D{{}}
			recordsCount = database.GetCountForFilter("employes", filter)
			PrintTaskNote("Count All Employes [No-Condition] : " + strconv.Itoa(recordsCount) + " records found")

			// Count employes with [shiftid = 2]
			filter = bson.D{{
				Key:   "shiftid",
				Value: 2,
			}}
			recordsCount = database.GetCountForFilter("employes", filter)
			PrintTaskNote("Count Employes with condition [shiftid = 2] : " + strconv.Itoa(recordsCount) + " records found")

			// Count employes with [experience > 60]
			filter = bson.D{{
				Key: "experience",
				Value: bson.D{{
					Key:   "$gt",
					Value: 60,
				}},
			}}
			recordsCount = database.GetCountForFilter("employes", filter)
			PrintTaskNote("Count Employes with condition [experience > 60] : " + strconv.Itoa(recordsCount) + " records found")

		},
	}
}
