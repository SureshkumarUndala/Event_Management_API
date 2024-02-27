package controllers

import (
	"net/http"

	"github.com/SureshkumarUndala/Event_Management_API/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	var events models.Event

	err := c.ShouldBindJSON(&events)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "all fields are required"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "event is created"})

}

func GetEventbyID(c *gin.Context) {

	return

}

func CreateEvent(c *gin.Context) {
	return
}
