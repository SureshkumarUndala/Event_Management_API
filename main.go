package main

import (
	"os"

	"github.com/SureshkumarUndala/Event_Management_API/db"
	routes "github.com/SureshkumarUndala/Event_Management_API/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	f, _ := os.Create("gin.log")

	server.Use(gin.LoggerWithWriter(f))

	routes.RegisterRoutes(server)

	server.Run(":8080")

}
