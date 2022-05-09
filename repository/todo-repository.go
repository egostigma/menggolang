package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"main/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TodoRepository interface {
	Save(todo *model.Todo)
	FindAll() []*model.Todo
	FindOne(id string) *model.Todo
}

type database struct {
	client *mongo.Client
}

const (
	DATABASE   = "menggolang"
	COLLECTION = "todo"
)

func NewTodoRepository() TodoRepository {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// mongodb+srv://USERNAME:PASSWORD@HOST:PORT
	MONGODB := os.Getenv("MONGODB_SERVER")
	// MONGODB := "mongodb://db"

	// Set client options
	clientOptions := options.Client().ApplyURI(MONGODB)

	clientOptions = clientOptions.SetMaxPoolSize(50)

	// Connect to MongoDB
	userClient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = userClient.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return &database{
		client: userClient,
	}
}

func (db *database) Save(todo *model.Todo) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), todo)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *database) FindAll() []*model.Todo {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var results []*model.Todo
	for cursor.Next(context.TODO()) {
		var v *model.Todo
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, v)
	}
	return results
}

func (db *database) FindOne(id string) *model.Todo {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor := collection.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}})
	var result *model.Todo
	var v *model.Todo
	err := cursor.Decode(&v)
	if err != nil {
		log.Fatal(err)
	}
	result = v

	return result
}
