package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/integrationninjas/go-app/handlers"
	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()

	// Initialize logging
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Register API handlers
	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/items", handlers.ItemsHandler)
	http.HandleFunc("/randomuser", handlers.GetRandomUser)
	port := os.Getenv("PORT")
	fmt.Println(port)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

// use godot package to load/read the .env file and
// return the value of the key
func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error in loading .env file: ", err)
	}
}
