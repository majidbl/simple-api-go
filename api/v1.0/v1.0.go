package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/majidzarephysics/university/api/v1.0/students"
	"github.com/majidzarephysics/university/api/v1.0/teachers"
	"github.com/majidzarephysics/university/api/v1.0/class"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong api v1.0",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/ping", ping)

		students.ApplyRoutes(v1)
		teachers.ApplyRoutes(v1)
		class.ApplyRoutes(v1)
	}
}
