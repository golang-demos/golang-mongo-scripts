package scripts

import (
	"strconv"

	"github.com/golang-demos/golang-mongo-scripts/database"
	"github.com/golang-demos/golang-mongo-scripts/models"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/golang-demos/chalk"
)

func Evaluation_Query_Demo() *models.DemoScript {
	return &models.DemoScript{
		Name:        "Element Query Operators",
		Description: "Examples to demonstrate use of $exists.",
		Handler: func() {
			var filter bson.D
			var recordsCount int

			PrintTaskTitle("api/scripts/05_Evaluation_operators.go")

			// Count employes with [cafeteria_credit == cafeteria_credit_spent]
			filter = bson.D{{
				Key: "$expr",
				Value: bson.D{{
					Key: "$eq",
					Value: []string{
						"$cafeteria_credit",
						"$cafeteria_credit_spent",
					},
				}},
			}}
			recordsCount = database.GetCountForFilter("employes", filter)
			PrintTaskNote("Count employes with " + chalk.BlueLight().String() + "[cafeteria_credit == cafeteria_credit_spent]" + chalk.Reset() + " : " + strconv.Itoa(recordsCount) + " records found")

		},
	}
}
