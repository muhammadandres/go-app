package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofor-little/env"
	"github.com/integrationninjas/go-app/handlers"
)

func main() {
	// load .env file
	// godotenv package
	port := goDotEnvVariable("PORT")

	// Initialize logging
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Register API handlers
	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/items", handlers.ItemsHandler)
	http.HandleFunc("/randomuser", handlers.GetRandomUser)

	fmt.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := env.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
