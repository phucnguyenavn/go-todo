package main

import (
	"go-todo/api/controller"
	"go-todo/api/repository"
	"go-todo/api/routes"
	"go-todo/api/service"
	"go-todo/infratructure"
	"go-todo/models"
	"os"
)

func init() {
	infratructure.LoadEnv()

}

func main() {
	router := infratructure.NewGinRouter() //router has been initialized and configured
	db := infratructure.NewDatabase()      // database has been initialized and configured
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()
	db.DB.AutoMigrate(&models.User{}) // migrating User model to database table
	router.Gin.Run(os.Getenv("SERVER_PORT"))

}
