package main

import (
	"github.com/CAVAh/api-tech-challenge/database"
	"github.com/CAVAh/api-tech-challenge/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequests()
}
