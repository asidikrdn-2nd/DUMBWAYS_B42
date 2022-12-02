package main

import (
	"dumbways-task_12/router"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// mengambil data dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// mengambil routingan dari package router
	router := router.Router()

	fmt.Println("Server running on http://localhost:4135")
	http.ListenAndServe("localhost:4135", router)
}
