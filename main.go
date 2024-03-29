package main

import (
	"github.com/CAVAh/api-tech-challenge/src/infra/db/gorm"
	"github.com/CAVAh/api-tech-challenge/src/infra/web/routes"
)

func main() {
	gorm.ConnectDB()
	routes.HandleRequests()
}
