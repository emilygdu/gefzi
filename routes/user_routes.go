package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.GET("/users", GetUsers)
	r.POST("/users", CreateUser)
}

//GET alle User

//Post erstelle neuen User
