package main

import (
	"time"

	"github.com/SureshkumarUndala/Event_Management_API/db"
	"github.com/SureshkumarUndala/Event_Management_API/middlewares"
	routes "github.com/SureshkumarUndala/Event_Management_API/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	db.InitDB()

	routes.RegisterRoutes(server)

	middlewares.Logger.Info().Printf("server started on 8080 port %d", time.Now().UnixNano())

	server.Use(func(c *gin.Context) {
		middlewares.Logger.Error().Printf("Invalid end point middlware was triggered %d", time.Now().UnixNano())

		c.JSON(404, gin.H{"status": "404", "message": "Invalid Endpoint"})

	})
	server.Run(":8080")

}
