package api

import (
	"github.com/gin-gonic/gin"
	"github.com/majidzarephysics/university/api/v1.0"
	
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong api",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", ping)
		
		apiv1.ApplyRoutes(api)
	}
}
