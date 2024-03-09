package di

import (
	"fmt"
	"shecare/internals/config"
	"shecare/internals/infrastructure/db"
	"shecare/internals/infrastructure/handler"
	"shecare/internals/infrastructure/repository"
	"shecare/internals/infrastructure/server"
	"shecare/internals/infrastructure/usecase"
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

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUsecase(config, userRepo)
	userHandler := handler.NewUserHandler(userUseCase, *config)

	adminRepo := repository.NewAdminRepository(db)
	adminUseCase := usecase.NewAdminUseCase(adminRepo, *config)
	adminHandler := handler.NewAdminHandler(adminUseCase)

	postRepo := repository.NewPostRepository(db)
	postUseCase := usecase.NewPostUseCase(postRepo)
	postHandler := handler.NewPostHandler(postUseCase)

	err = server.Server(*config, userHandler, adminHandler, postHandler)
	if err != nil {
		return err
	}

	return nil
}
