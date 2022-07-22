package Database

import (
	"context"
	"log"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseConnect() *mongo.Collection{
	client,err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
        log.Fatal(err)
    }

	err1 := client.Ping(context.TODO(),nil)
    if err1 != nil {
        log.Fatal(err1)
    }
    fmt.Println("Connected to MongoDB")
	 usersCollection := client.Database("Golang").Collection("Books_Collection")
	return usersCollection
}
