package main

import (
	"go-todo/api/controller"
	"go-todo/api/repository"
	"go-todo/api/routes"
	"go-todo/api/service"
	"go-todo/infrastructure"
	"go-todo/models"
	"os"
)

func init() {
	infrastructure.LoadEnv()

}

func main() {
	router := infrastructure.NewGinRouter() //router has been initialized and configured
	db := infrastructure.NewDatabase()      // database has been initialized and configured

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()
	db.DB.AutoMigrate(&models.User{})
	router.Gin.Run(os.Getenv("SERVER_PORT"))

}
