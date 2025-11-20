package routes

import (
	"gefzi/database"
	"gefzi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Routen Registrieren
func RegisterGroupCalendarRoutes(c *gin.Engine) {
	c.GET("/groupcalendars", GetGroupCalendars)
	c.POST("/groupcalendars", CreateGroupCalendars)
}

// GET alle Gruppen Kalender
func GetGroupCalendars(c *gin.Context) {
	var calendars []models.GroupCalendar                      //Variable calendars (Slice(Liste))-> für Gruppenkalender aus DB
	result := database.DB.Preload("Members").Find(&calendars) //holt Daten aus aktiver DB Verbindung, auch zugehörige User, führt SQL Abfrage aus -> .Find(&calenders) -> fügt alles in calenders
	if result.Error != nil {                                  //Fehlerbehandlung
		c.JSON(http.StatusInternalServerError, gin.H{ //Beispiel: Datenbank-Datei existiert nicht
			"error": result.Error.Error(),
		})
		return //Fehlermeldung HTTP 500
	}
	c.JSON(http.StatusOK, calendars) //Erfolgreiche Antwort -> HTTP 200, Daten sind in calendars und werden in JSON umgewandelt & ausgegeben
}

// POST erstelle neuen Gruppen Kalender
func CreateGroupCalendars(c *gin.Context) {
	var input models.GroupCalendar                   //leere Variable, wird mit JSON Daten aus Request gefüllt
	if err := c.ShouldBindJSON(&input); err != nil { //liest JSON-Body Request aus -> setzt in input
		c.JSON(http.StatusBadRequest, gin.H{ //wenn Body fehlerhaft, greift Fehlermeldung
			"error": err.Error(), //Beispiel: Typen falsch
		})
		return //Fehlermeldung wird ausgegeben (HTTP 400 Bad Request)
	}

	result := database.DB.Create(&input) //schreibt neuen Datensatz in Tabelle
	if result.Error != nil {             // Fehlerbehandlung
		c.JSON(http.StatusInternalServerError, gin.H{ //Beispiel: Datei schreibgeschützt -> gibt HTTP 500 zurück
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, input) //Erfolgsmeldung HTTP 201 Created -> Datensatz angelegt
}
