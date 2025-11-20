package routes

import (
	"gefzi/database"
	"gefzi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Routen Registrieren
func RegisterUserRoutes(c *gin.Engine) {
	c.GET("/users", GetUsers)
	c.POST("/users", CreateUser)
}

// GET alle User
func GetUsers(c *gin.Context) {
	var users []models.User //Variable users (Slice(Liste)), Für speicherung Daten
	result := database.DB.Preload("GroupCalendar").Find(&users)
	if result.Error != nil { //Fehlerbehandlung
		c.JSON(http.StatusInternalServerError, gin.H{ //Beispiel: Datenbank-Datei existiert nicht
			"error": result.Error.Error(),
		})
		return //Fehlermeldung HTTP 500
	}
	c.JSON(http.StatusOK, users) //Erfolgreiche Antwort -> HTTP 200
}

// Post erstelle neuen User
func CreateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil { //liest JSON-Body Request aus -> setzt in input
		c.JSON(http.StatusBadRequest, gin.H{ //wenn Body fehlerhaft, greift Fehlermeldung
			"error": err.Error(), //Beispiel: Typen falsch
		})
		return //Fehlermeldung (HTTP 400)
	}

	result := database.DB.Create(&input) //schreibt neuen Datensatz in Tabelle
	if result.Error != nil {             // Fehlerbehandlung
		c.JSON(http.StatusInternalServerError, gin.H{ //Beispiel: Datei schreibgeschützt -> gibt HTTP 500 zurück
			"error": result.Error.Error(),
		})
		return //Fehlermeldung (HTTP 500)
	}
	c.JSON(http.StatusCreated, input) //Erfolgsmeldung HTTP 201 Created -> Datensatz angelegt
}
