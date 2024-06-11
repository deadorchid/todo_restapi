package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/todo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"todos": "Nothing complete",
		})
	})

	return r
}
