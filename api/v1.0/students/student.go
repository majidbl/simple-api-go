package students

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	students := r.Group("/students")
	{
		//students.GET("/register", register)
		students.POST("/", create)
		students.GET("/", read)
		students.GET("/:id", readByID)
		students.DELETE("/:id", deleteByID)
		students.PATCH("/:id", update)
	}
}
