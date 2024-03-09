package server

import (
	"shecare/internals/config"
	"shecare/internals/infrastructure/handler"
	"shecare/internals/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

func Server(config config.Config, user *handler.UserHandler, admin *handler.AdminHandler, post *handler.PostHandler) error {
	engin := gin.Default()

	routes.UserRoutes(engin.Group("/user"), user, post)
	routes.AdminRoutes(engin.Group("admin"), admin)

	err := engin.Run(config.Port)
	if err != nil {
		return err
	}
	return nil
}
