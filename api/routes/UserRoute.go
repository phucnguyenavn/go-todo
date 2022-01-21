package routes

import (
	"go-todo/api/controller"
	"go-todo/infrastructure"
)

type UserRoute struct {
	UserController controller.UserController
	Router         infrastructure.Router
}

func NewUserRoute(userController controller.UserController, router infrastructure.Router) UserRoute {
	return UserRoute{UserController: userController, Router: router}
}

func (u UserRoute) Setup() {
	user := u.Router.Gin.Group("/user")
	{
		user.POST("/", u.UserController.NewUser)
	}
}
