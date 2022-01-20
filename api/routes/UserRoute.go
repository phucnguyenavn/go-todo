package routes

import (
	"go-todo/api/controller"
	"go-todo/infratructure"
)

type UserRoute struct {
	UserController controller.UserController
	Router         infratructure.Router
}

func NewUserRoute(userController controller.UserController, router infratructure.Router) UserRoute {
	return UserRoute{UserController: userController, Router: router}
}

func (u UserRoute) Setup() {
	user := u.Router.Gin.Group("/user")
	{
		user.POST("/", u.UserController.NewUser)
	}
}
