// Package is dope yo
package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"server/models"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string
const connectionString = ""

// Database Name
const dbName = "circusdb"

// Collection name
const collName = "skills"

// collection object/instance
var collection *mongo.Collection

// create connection with mongo db
func init() {

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")
}

func GetAllSkills(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllSkills()
	json.NewEncoder(w).Encode(payload)
}

func GetSelectSkills(w http.ResponseWriter, r *http.Request) {
	println("HERE")
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	payload := getSelectSkills(params["search"])
	json.NewEncoder(w).Encode(payload)
}

// CreateSkill makes a skill
func CreateSkill(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var skill models.Skill
	_ = json.NewDecoder(r.Body).Decode(&skill)
	insertOneSkill(skill)
	json.NewEncoder(w).Encode(skill)
}

func DeleteSkill(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	deleteOneSkill(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func EditSkill(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	editSkill(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func getAllSkills() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cursor.Next(context.Background()) {
		var result bson.M
		e := cursor.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}

		results = append(results, result)

	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	cursor.Close(context.Background())
	return results

}

func getSelectSkills(skillName string) []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{"name", skillName}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cursor.Next(context.Background()) {
		var result bson.M
		e := cursor.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}

		results = append(results, result)

	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	cursor.Close(context.Background())
	return results

}

func insertOneSkill(skill models.Skill) {
	insertResult, err := collection.InsertOne(context.Background(), skill)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Record inserted ", insertResult.InsertedID)
}

// delete one skill by ID
func deleteOneSkill(skill string) {
	fmt.Println(skill)
	id, _ := primitive.ObjectIDFromHex(skill)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}

// update skill with changed name
// TODO: Add other changes
func editSkill(skill string) {
	fmt.Println(skill)
	id, _ := primitive.ObjectIDFromHex(skill)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"name": "x"}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}
