package teachers

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	teachers := r.Group("/teachers")
	{
		//students.GET("/register", register)
		teachers.POST("/", create)
		teachers.GET("/", read)
		teachers.GET("/:id", readByID)
		teachers.DELETE("/:id", deleteByID)
		teachers.PATCH("/:id", update)
	}
}
