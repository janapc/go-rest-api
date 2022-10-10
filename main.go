package main

import (
	"fmt"

	"go-rest-api/database"
	"go-rest-api/routes"
)

func main() {
	database.ConnectionDatabase()
	fmt.Println("Server start at http://localhost:8000")
	routes.HandleRequest()
}
