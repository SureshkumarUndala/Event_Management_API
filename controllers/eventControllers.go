package controllers

import (
	"net/http"

	"github.com/SureshkumarUndala/Event_Management_API/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {

	events, err := models.GetallEvents()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "all events are fetched", "data": events})

}

func GetEventbyID(c *gin.Context) {
	eventID := c.Param("id")

	if eventID == "" {
		c.JSON(400, gin.H{"status": 400, "message": "id was not found in the request params"})
	}

	c.JSON(http.StatusAccepted, gin.H{"status": 400, "message": "event is fetched byt ID"})

}

func CreateEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": err.Error()})
		return
	}

	id, err := event.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 10, "message": err})
		return

	}

	c.JSON(http.StatusAccepted, gin.H{"status": "200", "message": "event created successfully", "id": id})
}
