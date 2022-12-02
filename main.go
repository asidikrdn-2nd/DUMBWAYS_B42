package main

import (
	"dumbways-task_7/router"
	"fmt"
	"net/http"
)

func main() {
	// mengambil routingan dari package router
	router := router.Router()

	fmt.Println("Server running on http://localhost:4135")
	http.ListenAndServe(":4135", router)
}
