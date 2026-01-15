package routes

import (
	"gefzi/internal/database"
	"gefzi/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Routen Registrieren
func RegisterUserRoutes(c *gin.Engine) {
	c.GET("/users", GetUsers)
	c.GET("/users/:id", GetUserById)
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

// GET User by id
func GetUserById(c *gin.Context) {
	idStr := c.Param("id")         // Id aus der URL lesen
	id, err := strconv.Atoi(idStr) //String in int umwandeln -> da id in der URL immer als string übergeben wird
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id", //Fehlermeldung wenn id nicht vorhanden oder falsche id gesendet
		})
		return
	}
	var user models.User                                            //Variable users (wird später mit DB-Daten gefüllt)
	result := database.DB.Preload("GroupCalendar").First(&user, id) //Datenbankabfrage
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return //Ausgabe Fehlermeldung wenn user nicht gefunden
	}

	c.JSON(http.StatusOK, user) //rückgabe User by Id daten & Erfolgsmeldung
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
