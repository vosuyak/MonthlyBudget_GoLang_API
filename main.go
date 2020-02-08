package main

import (
	_ "monthly-budget/src/data"
	"os"
	"monthly-budget/src/router"
	"net/http"
)

func main() {
	r := router.ExpenseRoutes()
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), r)
}
