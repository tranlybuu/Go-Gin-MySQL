package model

import (
	"context"
	"go-gin/initializer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

var mongoClient *mongo.Client
var collection *mongo.Collection

type Course struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Image       string             `json:"image" bson:"image"`
	Type        string             `json:"type" bson:"type"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
	Deleted     bool               `json:"deleted" bson:"deleted"`
}

func connectDB() {
	mongoClient = initializer.ConnectMongoDB()
	collection = mongoClient.Database("test").Collection("courses")
}

func disconnectDB() {
	if err := mongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func FindCourseList() []bson.M {
	connectDB()
	defer disconnectDB()
	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results
}

func FindCourseDetail(id string) *Course {
	connectDB()
	defer disconnectDB()
	var course Course
	courseId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": courseId}
	err := collection.FindOne(context.TODO(), filter).Decode(&course)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		panic(err)
	}
	return &course
}

func CreateCourse(course Course) string {
	connectDB()
	defer disconnectDB()
	_, err := collection.InsertOne(context.TODO(), course)
	if err != nil {
		panic(err)
	}
	return "Created successfully"
}

func DeleteCourseById(id string) string {
	connectDB()
	defer disconnectDB()
	courseId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": courseId}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	return "Deleted successfully"
}

func UpdateCourseById(course Course) string {
	connectDB()
	defer disconnectDB()
	filter := bson.M{"_id": course.ID}
	_, err := collection.UpdateOne(context.TODO(), filter, bson.M{"$set": course})
	if err != nil {
		panic(err)
	}
	return "Updated successfully"
}
