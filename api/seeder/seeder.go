package seeder

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/golang-demos/golang-mongo-scripts/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TruncateDatabase() bool {
	database.Db.Collection("employes").DeleteMany(context.Background(), bson.M{})
	database.Db.Collection("projects").DeleteMany(context.Background(), bson.M{})
	return true
}

func SeedDatabase(doTruncate bool) bool {
	faker.Name()

	NumberOfProjects := 10
	NumberOfEmployes := 200

	dbEmployesCount, _ := database.Db.Collection("employes").CountDocuments(context.Background(), bson.D{{}})
	log.Print(dbEmployesCount, " Employes already exists")
	if dbEmployesCount > int64(NumberOfEmployes) && (!doTruncate) {
		return false
	}
	TruncateDatabase()

	var ProjectMappings []primitive.ObjectID
	for i := 0; i < NumberOfProjects; i++ {
		ProjectName := faker.LastName() + []string{
			" Infotech",
			" Solutions",
			" Pvt. Ltd",
		}[rand.Intn(3)]

		result, err := database.Db.Collection("projects").InsertOne(context.Background(), bson.M{
			"name": ProjectName,
		})
		if err != nil {
			log.Print(err)
			break
		}
		ProjectMappings = append(ProjectMappings, result.InsertedID.(primitive.ObjectID))
	}

	for i := 0; i < NumberOfEmployes; i++ {
		ctcSeeder := rand.Intn(20)
		_, err := database.Db.Collection("employes").InsertOne(context.Background(), bson.M{
			"name": faker.Name(),
			"designation": []string{
				"Project Manager",
				"Software Architect",
				"Senior Software Developer",
				"Lead Software Developer",
			}[rand.Intn(4)],
			"project_id": ProjectMappings[rand.Intn(NumberOfProjects)],
			"ctc":        (16 + ctcSeeder) * 50000,
			"experience": (24 + (ctcSeeder * 6)),
			"isverified": [2]bool{true, false}[rand.Intn(2)],
			"joindate":   primitive.Timestamp{T: uint32(time.Now().AddDate(0, 0, (-1 * (60 + (10 * rand.Intn(30))))).Unix())},
			"shiftid":    rand.Intn(4) + 1,
		})
		if err != nil {
			log.Print(err)
			break
		}
	}
	return true
}
