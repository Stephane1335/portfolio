package main

import (
	"go_portfolio/databases"
	"go_portfolio/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	databases.ConnectDatabase()
	routes.SetupRoutes(r)
	r.Run()
}
