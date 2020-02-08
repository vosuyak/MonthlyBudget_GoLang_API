package router

import (
	"github.com/gorilla/mux"

	"monthly-budget/src/controller/expense"
)

// ExpenseRoutes - routes
func ExpenseRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/expense", expense.CreateExpense).Methods("POST")
	r.HandleFunc("/expense/{id}", expense.GetExpense).Methods("GET")
	r.HandleFunc("/expenses", expense.GetAllExpense).Methods("GET")
	r.HandleFunc("/expense/{id}", expense.UpdateExpense).Methods("PUT")
	r.HandleFunc("/expense/{id}", expense.DeleteExpense).Methods("DELETE")
	return r
}
