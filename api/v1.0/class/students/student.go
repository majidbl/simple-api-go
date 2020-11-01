package students

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	students := r.Group("/students")
	{
		//define class teacher route
		students.POST("/:sid/:cid", create)
		students.GET("/:sid", read)
		students.GET("/:sid/:cid", readByID)
		students.DELETE("/:id", deleteByID)
		
	}
}
