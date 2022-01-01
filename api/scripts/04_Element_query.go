package scripts

import (
	"strconv"

	"github.com/golang-demos/golang-mongo-scripts/database"
	"github.com/golang-demos/golang-mongo-scripts/models"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/golang-demos/chalk"
)

func Element_Query_Demo() *models.DemoScript {
	return &models.DemoScript{
		Name:        "Element Query Operators",
		Description: "Examples to demonstrate use of $exists.",
		Handler: func() {
			var filter bson.D
			var recordsCount int

			PrintTaskTitle("api/scripts/04_Element_query.go")

			// Count employes having ctc field and has value greater than 1,200,000
			filter = bson.D{{
				Key: "ctc",
				Value: bson.M{
					"$exists": true,
					"$gt":     1200000,
				},
			}}
			recordsCount = database.GetCountForFilter("employes", filter)
			PrintTaskNote("Count employes having ctc field and has " + chalk.BlueLight().String() + "[ctc > 1,200,000]" + chalk.Reset() + " : " + strconv.Itoa(recordsCount) + " records found")

		},
	}
}
