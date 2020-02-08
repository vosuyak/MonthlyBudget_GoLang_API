package data

import (
	"context"
	"time"

	"monthly-budget/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

// GetCollection : Goes to the GetClient Client and returns all Collections
func GetCollection(coll string) *mongo.Collection {
	var collection *mongo.Collection
	switch coll {
	case "expenses":
		collection = GetClient().Database(os.Getenv("MONGO_DB_NAME")).Collection("expenses")
	}
	return collection
}

// CheckIfIDExist - Check if collection entered along with the ID exist
// in the collection, if not return an error
func CheckIfIDExist(coll string, id primitive.ObjectID) error {
	var result error
	switch coll {
	case "expenses":
		var expense model.Expense
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := GetCollection("expenses")
		filter := bson.D{
			{"_id", id},
		}
		opts := options.FindOne()
		result = collection.FindOne(ctx, filter, opts).Decode(&expense)
	}
	return result
}