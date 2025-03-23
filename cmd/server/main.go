package main

import (
	"go-cc/src/config"
	"go-cc/src/database"
	"go-cc/src/routes"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	db.ConnectToDB()

	r := routes.SetupRoutes()

	log.Println("Starting server on port", config.AppPort())
	log.Fatal(http.ListenAndServe(":"+config.AppPort(), r))
}
