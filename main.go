package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty"  bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty"  bson:"lastname,omitempty"`
}

var client *mongo.Client
var DB_NAME string = "persondb"
var MONGO_URI string = "mongodb://mongoadmin:secret@localhost:27017/?authSource=admin"

func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var people []Person
	collection := client.Database(DB_NAME).Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database(DB_NAME).Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person Person
	collection := client.Database(DB_NAME).Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

func DeletePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := client.Database(DB_NAME).Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	resp, err := collection.DeleteOne(ctx, Person{ID: id})
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	message := strconv.Itoa(int(resp.DeletedCount)) + " records deleted"
	response.Write([]byte(`{"message": "` + message + `"}`))
}

func main() {
	fmt.Println("Starting the application")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI))

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", DeletePersonEndpoint).Methods("DELETE")

	http.ListenAndServe(":12345", router)
}
