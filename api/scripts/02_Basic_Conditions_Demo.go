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

			// Count all employes [No-Condition]
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

			// Count employes with [shiftid != 3]
			filter = bson.D{{
				Key: "shiftid",
				Value: bson.D{{
					Key:   "$ne",
					Value: 3,
				}},
			}}
			recordsCount = database.GetCountForFilter("employes", filter)
			PrintTaskNote("Count Employes with condition [shiftid != 3] : " + strconv.Itoa(recordsCount) + " records found")

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

			// Count employes with [age >= 34]
			filter = bson.D{{
				Key: "age",
				Value: bson.D{{
					Key:   "$gte",
					Value: 34,
				}},
			}}
			recordsCount = database.GetCountForFilter("employes", filter)
			PrintTaskNote("Count Employes with condition [age >= 34] : " + strconv.Itoa(recordsCount) + " records found")

			// Count employes with [shiftid < 3]
			filter = bson.D{{
				Key: "shiftid",
				Value: bson.D{{
					Key:   "$lt",
					Value: 3,
				}},
			}}
			recordsCount = database.GetCountForFilter("employes", filter)
			PrintTaskNote("Count Employes with condition [shiftid < 3] : " + strconv.Itoa(recordsCount) + " records found")

			// Count employes with [ctc <= 1,150,000]
			filter = bson.D{{
				Key: "ctc",
				Value: bson.D{{
					Key:   "$lte",
					Value: 1150000,
				}},
			}}
			recordsCount = database.GetCountForFilter("employes", filter)
			PrintTaskNote("Count Employes with condition [ctc <= 1,150,000] : " + strconv.Itoa(recordsCount) + " records found")

			// Count employes with a designation from ["Project Manager", "Software Architect", "Senior Software Developer"]
			filter = bson.D{{
				Key: "designation",
				Value: bson.D{{
					Key: "$in",
					Value: []string{
						"Project Manager",
						"Software Architect",
						"Senior Software Developer",
					},
				}},
			}}
			recordsCount = database.GetCountForFilter("employes", filter)
			PrintTaskNote("Count employes with a designation from [\"Project Manager\", \"Software Architect\", \"Senior Software Developer\"] : " + strconv.Itoa(recordsCount) + " records found")

			// Count employes with a designation NOT from ["Project Manager", "Software Architect", "Senior Software Developer"]
			filter = bson.D{{
				Key: "designation",
				Value: bson.D{{
					Key: "$nin", // NOT IN
					Value: []string{
						"Software Architect",
						"Lead Software Developer",
					},
				}},
			}}
			recordsCount = database.GetCountForFilter("employes", filter)
			PrintTaskNote("Count employes with a designation NOT from [\"Software Architect\", \"Lead Software Developer\"] : " + strconv.Itoa(recordsCount) + " records found")

			// Logical Conditions
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
