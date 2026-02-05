package routes

import (
	"gefzi/internal/database"
	"gefzi/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Routen Registrieren
func RegisterGroupCalendarRoutes(c *gin.Engine) {
	c.GET("/groupcalendars", GetGroupCalendars)
	c.GET("/groupcalendars/:id", GetGroupCalendarsById)
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

// GET Gruppen Kalender by Id
func GetGroupCalendarsById(c *gin.Context) {
	idStr := c.Param("id")         // Id aus der URL lesen
	id, err := strconv.Atoi(idStr) //String in int umwandeln -> da id in der URL immer als string übergeben wird
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid groupcalendar id", //Fehlermeldung wenn id nicht vorhanden oder falsche id gesendet
		})
		return
	}
	var calendar models.GroupCalendar                             //Variable
	result := database.DB.Preload("Members").First(&calendar, id) //Datenbankabfrage
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "groupcalendar not found",
		})
		return //Ausgabe Fehlermeldung wenn Kalender nicht gefunden
	}

	c.JSON(http.StatusOK, calendar) //rückgabe
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
