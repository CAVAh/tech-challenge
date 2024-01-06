package main

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/gorm"
	"github.com/CAVAh/api-tech-challenge/src/adapters/driver/http/routes"
)

func main() {
	gorm.ConnectDB()
	routes.HandleRequests()
}
