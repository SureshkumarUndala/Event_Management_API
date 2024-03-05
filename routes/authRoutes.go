package routes

import (
	"github.com/SureshkumarUndala/Event_Management_API/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	auth := server.Group("/api/v1")
	{

		auth.POST("/login", controllers.Login)
		auth.POST("/signup", controllers.RegisterUser)
		auth.POST("/forgotpasswpord", controllers.Resetpassword)

	}

	event := server.Group("/api/v1")
	{
		event.GET("/events", controllers.GetEvents)
		event.GET("/events/:id", controllers.GetEventbyID)
		event.POST("/events", controllers.CreateEvent)
	}

}
