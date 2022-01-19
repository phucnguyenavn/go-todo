package api

import (
	"github.com/gin-gonic/gin"
	"go-todo/main/models"
	"net/http"
)

func NewUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
}
