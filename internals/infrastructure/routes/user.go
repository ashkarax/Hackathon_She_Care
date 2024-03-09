package routes

import (
	"shecare/internals/infrastructure/handler"
	userHandler "shecare/internals/infrastructure/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup, userHandler *userHandler.UserHandler, post *handler.PostHandler) {
	engin.POST("/", userHandler.UserSignup)
	engin.POST("", userHandler.UserSignup)

	engin.POST("/post", post.NewPost)
}
