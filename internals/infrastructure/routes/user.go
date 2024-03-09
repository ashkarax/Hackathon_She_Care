package routes

import (
	userHandler "shecare/internals/infrastructure/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup, userHandler userHandler.UserHandler) {
	engin.POST("/", userHandler.Signup)
	engin.POST("", userHandler.Signup)
}
