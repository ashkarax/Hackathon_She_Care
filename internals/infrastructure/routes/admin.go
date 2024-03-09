package routes

import (
	"shecare/internals/infrastructure/handler"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engin *gin.RouterGroup, adminhandler handler.AdminHandler) {
	engin.POST("", adminhandler.AdminLogin)
}
