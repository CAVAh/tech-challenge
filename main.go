package main

import (
	"github.com/CAVAh/api-tech-challenge/src/db/gorm"
	"github.com/CAVAh/api-tech-challenge/src/web/routes"
)

func main() {
	gorm.ConnectDB()
	routes.HandleRequests()
}
