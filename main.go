package main

import (
	"todo_app/database"
	"todo_app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.InitDB()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
