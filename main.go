package main

import (
	rt "github.com/lukaszozimek/license_service/router"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally

	if port == "" {
		port = "8000" //localhost
	}

	log.Printf("Server started")

	router := rt.NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
