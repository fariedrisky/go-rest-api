package main

import (
	"log"
	"net/http"

	"github.com/fariedrisky/go-rest-api/database"
	"github.com/fariedrisky/go-rest-api/routes"
)

func main() {
	database.ConnectDB()
	router := routes.SetupRoutes()
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
