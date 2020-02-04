package main

import (
	_ "expense/data"
	"os"
	"expense/routers"
	"net/http"
)

func main() {
	r := routers.ExpenseRoutes()
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), r)
}
