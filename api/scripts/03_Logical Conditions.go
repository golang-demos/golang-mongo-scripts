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

			PrintTaskTitle("api/scripts/03_Logical_Conditions.go")

			// Count employes with a designation as "Project Manager" and (CTC > 1,200,000)
			filter = bson.D{{
				Key: "$and",
				Value: []bson.D{
					{{
						Key:   "designation",
						Value: "Project Manager",
					}},
					{{
						Key: "ctc",
						Value: bson.D{{
							Key:   "$gt",
							Value: 1200000,
						}},
					}},
				},
			}}
			recordsCount = database.GetCountForFilter("employes", filter)
			PrintTaskNote("Count employes with a designation as \"Project Manager\" and (CTC > 1,200,000) : " + strconv.Itoa(recordsCount) + " records found")

			// Count employes with ctc > 1,400,000 OR is verified
			filter = bson.D{{
				Key: "$or",
				Value: []bson.D{
					{{
						Key: "ctc",
						Value: bson.D{{
							Key:   "$gt",
							Value: 1400000,
						}},
					}},
					{{
						Key:   "designation",
						Value: "Software Architect",
					}},
				},
			}}
			recordsCount = database.GetCountForFilter("employes", filter)
			PrintTaskNote("Count employes with ctc > 1,400,000 OR is verified : " + strconv.Itoa(recordsCount) + " records found")

			// Count employes with experience not less than 4 years
			filter = bson.D{{
				Key: "experience",
				Value: bson.D{{
					Key: "$not",
					Value: bson.D{{
						Key:   "$lt",
						Value: 48,
					}},
				}},
			}}
			recordsCount = database.GetCountForFilter("employes", filter)
			PrintTaskNote("Count employes with experience not less than 4 years : " + strconv.Itoa(recordsCount) + " records found")

		},
	}
}
