package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Expense - data struct
type Expense struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Amount      int                `json:"amount" bson:"amount"`
	Owner       string             `json:"owner" bson:"owner"`
}
