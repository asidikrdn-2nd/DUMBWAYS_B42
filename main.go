package main

import (
	"dumbways-task_8/router"
	"fmt"
	"net/http"
)

func main() {
	// mengambil routingan dari package router
	router := router.Router()

	fmt.Println("Server running on http://localhost:4135")
	http.ListenAndServe("localhost:4135", router)
}
