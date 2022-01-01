package scripts

import (
	"strconv"

	"github.com/golang-demos/golang-mongo-scripts/database"
	"github.com/golang-demos/golang-mongo-scripts/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Logical_Conditions_Demo() *models.DemoScript {
	return &models.DemoScript{
		Name:        "Logical Conditions Demo",
		Description: "Examples to demonstrate various conditions using Logical operators.",
		Handler: func() {
			var filter bson.D
			var recordsCount int

			// Count employes with a designation as "Project Manager" and (CTC > 1,200,000)
			filter = bson.D{{
				Key: "$and",
				Value: []bson.D{
					{{
						Key:   "designation",
						Value: "Software Architect",
					}},
					{{
						Key:   "$gt",
						Value: 1200000,
					}},
				},
			}}
			recordsCount = database.GetCountForFilter("employes", filter)
			PrintTaskNote("Count employes with a designation as \"Project Manager\" and (CTC > 1,200,000) : " + strconv.Itoa(recordsCount) + " records found")
		},
	}
}
