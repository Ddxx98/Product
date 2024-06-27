package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//password := "Tj8N5NNHY5QPiyCq"
//url := "mongodb+srv://deexith2016:Tj8N5NNHY5QPiyCq@cluster0.3veksmi.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
//url := "mongodb+srv://deexith2016:Tj8N5NNHY5QPiyCq@cluster0.3veksmi.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://deexith2016:Tj8N5NNHY5QPiyCq@cluster0.3veksmi.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}