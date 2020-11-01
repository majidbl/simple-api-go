package class

import (
	"github.com/gin-gonic/gin"
	"github.com/majidzarephysics/university/api/v1.0/class/students"
	"github.com/majidzarephysics/university/api/v1.0/class/teachers"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong class api v1.0",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	class := r.Group("/class")
	{
		class.GET("/ping", ping)

		students.ApplyRoutes(class)
		teachers.ApplyRoutes(class)
	}
}
