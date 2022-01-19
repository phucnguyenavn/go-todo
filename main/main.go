package main

import (
	"go-todo/main/api/controller"
	"go-todo/main/api/repository"
	"go-todo/main/api/routes"
	"go-todo/main/api/service"
	"go-todo/main/infratructure"
	"go-todo/main/models"
	"os"
)



func init() {
	infratructure.LoadEnv()

}

func main() {
	router := infratructure.NewGinRouter()
	db := infratructure.NewDatabase()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController,router)
	userRoute.Setup()
	db.DB.AutoMigrate(&models.User{})
	router.Gin.Run(os.Getenv("SERVER_PORT"))

}
