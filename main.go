package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	initDB()

	http.HandleFunc("/netflix", home)
	http.HandleFunc("/netflix/random", random)

	fmt.Println("Server listen on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
