package controllers

import (
	"net/http"
	"expense/common"
	"expense/data"
	"expense/models"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gorilla/mux"
)

var (
	collectionExp = data.GetCollection("expenses")
)

// CreateExpense - creation of a new Expense
func CreateExpense(w http.ResponseWriter, r *http.Request) {
	var exp models.Expense

	err := json.NewDecoder(r.Body).Decode(&exp)
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error with decoding form",
		)
		return
	}

	// link to Repository,define ID of expense, pass struct into Create Method
	exp.ID = primitive.NewObjectID()
	repo := &ExpenseRepository{C: collectionExp}
	errExp := repo.Create(exp)
	if errExp != nil {
		common.DisplayError(w, errExp, http.StatusInternalServerError,
			"error with inserting into database",
		)
		return
	}

	common.DisplaySuccess(w, nil, http.StatusOK, exp)
}

// GetExpense - creation of a new Expense
func GetExpense(w http.ResponseWriter, r *http.Request) {
	var edu models.Expense
	vars := mux.Vars(r)
	_id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"err in retrieving id",
		)
		return
	}
	repo := &ExpenseRepository{C: collectionExp}
	edu, errEdu := repo.Get(_id)
	if errEdu != nil {
		common.DisplayError(w, errEdu, http.StatusInternalServerError,
			"err in retrieving education after decoding",
		)
		return
	}

	common.DisplaySuccess(w, nil, http.StatusOK, edu)
}

// GetAllExpense - creation of a new Expense
func GetAllExpense(w http.ResponseWriter, r *http.Request) {
	repo := &ExpenseRepository{C: collectionExp}
	expenses, err := repo.GetAll()
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error retrieving all expenses",
		)
		return
	}

	common.DisplaySuccess(w, true, http.StatusOK, expenses)
}

// UpdateExpense - creation of a new Expense
func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	var exp models.Expense

	vars := mux.Vars(r)
	_id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error in retrieving id",
		)
		return
	}

	errExp := json.NewDecoder(r.Body).Decode(&exp)
	if errExp != nil {
		common.DisplayError(w, errExp, http.StatusInternalServerError,
			"error in decoding expense",
		)
		return
	}


	exp.ID = _id
	repo := &ExpenseRepository{C: collectionExp}
	result := repo.Update(exp)
	if result != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error in updating",
		)
		return
	}
	// show success message in response
	common.DisplaySuccess(w, nil, http.StatusCreated, &exp)
}

// DeleteExpense - creation of a new Expense
func DeleteExpense(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			common.DisplayError(w, err, http.StatusInternalServerError,
				"error in retrieving id expense path",
			)
			return
		}

		repo := &ExpenseRepository{C: collectionExp}
		errExp := repo.Delete(_id)
		if errExp != nil {
			common.DisplayError(w, errExp, http.StatusInternalServerError,
				"error in deleting id expense db",
			)
			return
		}
		// show success message in response
		common.DisplaySuccess(w, true, http.StatusOK, "deleted")
}