package infratructure

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	Gin *gin.Engine
}
func NewGinRouter() Router {

	httpRouter := gin.Default()

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})
	return Router{
		Gin: httpRouter,
	}

}
