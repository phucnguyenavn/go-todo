package controller

import (
	"github.com/gin-gonic/gin"
	"go-todo/main/api/service"
	"go-todo/main/models"
	"go-todo/main/util"
	"net/http"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{userService: userService}
}



func (u *UserController) NewUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = u.userService.NewUser(user)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to create post")
		return
	}
	util.SuccessJSON(c, http.StatusCreated, "Successfully Created User")
}
