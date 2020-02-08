package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
	"monthly-budget/src/core"
	"os"
)

func init() {
	core.LoadEnv()
	GetClient()
}

// GetClient - Connect to Client
func GetClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credential := options.Credential{
		Username: os.Getenv("MONGO_DB_USER"),
		Password: os.Getenv("MONGO_DB_PASSWORD"),
		AuthSource: os.Getenv("MONGO_DB_NAME"),
	}
	clientOpts := options.Client().ApplyURI("mongodb://"+os.Getenv("MONGO_DB_HOST")+":"+os.Getenv("MONGO_DB_PORT")).SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Could not connect to server: \n", err,"\n",credential,"\n"+"mongodb://"+os.Getenv("MONGO_DB_USER")+":"+os.Getenv("MONGO_DB_PASSWORD")+"@"+os.Getenv("MONGO_DB_HOST")+":"+os.Getenv("MONGO_DB_NAME")+":"+os.Getenv("MONGO_DB_NAME"))
	} else {
		fmt.Println("Connected to MongoDB")
	}
	return client
}
