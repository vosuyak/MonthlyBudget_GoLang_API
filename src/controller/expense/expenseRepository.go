package expense

import (
	"context"
	"monthly-budget/src/model"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ExpenseRepository - Repo
type ExpenseRepository struct {
	C *mongo.Collection
}

// Create - Create Method
func (r *ExpenseRepository) Create(expense model.Expense) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.C.InsertOne(ctx, expense)
	return err
}

// Get - Get Method
func (r *ExpenseRepository) Get(id primitive.ObjectID) (model.Expense, error) {
	var expense model.Expense
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	opts := options.FindOne()
	err := r.C.FindOne(ctx, filter, opts).Decode(&expense)
	return expense, err
}

// GetAll - GetAll Method
func (r *ExpenseRepository) GetAll() ([]model.Expense, error) {
	var expenses []model.Expense
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	item,err := r.C.Find(ctx,bson.M{})
	var expense model.Expense
	for item.Next(ctx){
		item.Decode(&expense)
		expenses = append(expenses,expense)
	}
	defer item.Close(ctx)
	return expenses, err
}

// Update - Update Method
func (r *ExpenseRepository) Update(expense model.Expense) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id":expense.ID}
	update := bson.M{
		"$set":bson.M{
			"title": expense.Title,
			"description": expense.Description,
			"amount": expense.Amount,
			"owner": expense.Owner,
		},
	}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := r.C.FindOneAndUpdate(ctx,filter,update,opts).Decode(&expense)
	return err
}

// Delete - Delete Method
func (r *ExpenseRepository) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	_, err := r.C.DeleteOne(ctx, filter, opts)
	return err
}

