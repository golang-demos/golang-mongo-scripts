package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employes struct {
	ID          primitive.ObjectID  `json:"_id" bson:"_id"`
	Name        string              `json:"name" bson:"name"`
	Age         int16               `json:"age" bson:"age"`
	Designation string              `json:"designation" bson:"designation"`
	ProjectID   primitive.ObjectID  `json:"project_id" bson:"project_id"`
	Ctc         float64             `json:"ctc" bson:"ctc"`
	Experience  int16               `json:"experience" bson:"experience"`
	IsVerified  bool                `json:"isverified" bson:"isverified"`
	JoinDate    primitive.Timestamp `json:"joindate" bson:"joindate"`
	ShiftId     uint8               `json:"shiftid" bson:"shiftid"`
}

type Projects struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

type DemoScript struct {
	Name        string
	Description string
	Handler     func()
}
