package routers

import (
	"github.com/gorilla/mux"
	"expense/controllers"
)

// ExpenseRoutes - routes
func ExpenseRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/expense", controllers.CreateExpense).Methods("POST")
	r.HandleFunc("/expense/{id}", controllers.GetExpense).Methods("GET")
	r.HandleFunc("/expenses", controllers.GetAllExpense).Methods("GET")
	r.HandleFunc("/expense/{id}", controllers.UpdateExpense).Methods("PUT")
	r.HandleFunc("/expense/{id}", controllers.DeleteExpense).Methods("DELETE")
	return r
}
