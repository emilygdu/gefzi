package routes

import (
	"gefzi/database"
	"gefzi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Routen Registrieren
func RegisterEventRoutes(c *gin.Engine) {
	c.GET("/events", GetEvents)
	c.POST("/events", CreateEvents)
}

// GET alle Events (Termine)
func GetEvents(c *gin.Context) {
	var events []models.Event           //Variable events (Slice(Liste))-> für Events aus DB
	result := database.DB.Find(&events) //holt Daten aus aktiver DB Verbindung, führt SQL Abfrage aus -> .Find(&events) -> fügt alles in events
	if result.Error != nil {            //Fehlerbehandlung
		c.JSON(http.StatusInternalServerError, gin.H{ //Beispiel: Datenbank-Datei existiert nicht
			"error": result.Error.Error(),
		})
		return //Fehlermeldung HTTP 500
	}
	c.JSON(http.StatusOK, events) //Erfolgreiche Antwort -> HTTP 200
}

// POST Event (Termin erstellen)
func CreateEvents(c *gin.Context) {
	var input models.Event                           //leere Variable wird mit Daten aus JSON gefüllt
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
