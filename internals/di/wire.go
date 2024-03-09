package di

import (
	"fmt"
	"shecare/internals/config"
	"shecare/internals/infrastructure/db"
	userHandler "shecare/internals/infrastructure/handler"
	"shecare/internals/infrastructure/server"
)

func InitializeDependency() error {
	config, err := config.InitConfig()
	if err != nil {
		return err
	}

	db, err := db.InitDB(config)
	if err != nil {
		return err
	}
	fmt.Println("=", db)

	userHandler := userHandler.NewUserHandler(*config)

	err = server.Server(*config, *userHandler)
	if err != nil {
		return err
	}

	return nil
}
