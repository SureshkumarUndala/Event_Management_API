package routes

import (
	"github.com/SureshkumarUndala/Event_Management_API/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEventbyID)
	server.POST("/events", controllers.CreateEvent)

}
