package routes

import (
	"gefzi/database"
	"gefzi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(r *gin.Engine) {
	events := r.Group("/events")
	{
		events.GET("/", getEvents)
		events.POST("/", createEvents)
	}
}

func getEvents(c *gin.Context) {
	var events []models.Event
	database.DB.Find(&events)
	c.JSON(http.StatusOK, events)
}

func createEvents(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&event)
	c.JSON(http.StatusCreated, event)
}
