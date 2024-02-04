package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	productRoute := r.Group("/products")
	{
		productRoute.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
}
