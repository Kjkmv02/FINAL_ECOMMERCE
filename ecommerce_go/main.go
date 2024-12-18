package main

import (
	"log"
	"net/http"

	"ecommerce_go/config"
	"ecommerce_go/routes"
)

func main() {
	config.Connect()
	routes.SetupRoutes()

	log.Println("Servidor corriendo en http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
