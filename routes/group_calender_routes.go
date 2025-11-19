package routes

import (
	"github.com/gin-gonic/gin"
)

// Routen Registrieren
func RegisterGroupCalenderRoutes(r *gin.Engine) {
	r.GET("/groupcalenders", GetGroupCalenders)
	r.POST("/groupcalenders", CreateGroupCalenders)
}

//GET alle Gruppen Kalender

//POST erstelle neuen Gruppen Kalender
