package main

import (
	"gefzi/database"
	"gefzi/models"
	"gefzi/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//Datenbank anbinden
	database.Connect()
	//Automatische Tabellen erstellen für Datenbank
	database.DB.AutoMigrate(&models.User{}, &models.Event{})

	r := gin.Default()
	//CORS aktivieren fürs Frontend
	r.Use(cors.Default())

	//Test Get (Beispielroute) -> später entfernen
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Yayy der erste Test klappt",
		})
	})

	//API-Routen registrieren
	routes.RegisterGroupCalendarRoutes(r)
	routes.RegisterUserRoutes(r)
	routes.RegisterEventRoutes(r)

	//Logik für Login

	//Logik für freie Zeiträume berechnen

	//Server starten (localhost)
	r.Run(":8080")
}
