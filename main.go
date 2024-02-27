package main

import (
	"github.com/SureshkumarUndala/Event_Management_API/db"
	routes "github.com/SureshkumarUndala/Event_Management_API/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}
