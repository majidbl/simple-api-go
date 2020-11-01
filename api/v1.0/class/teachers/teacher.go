package teachers

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	teachers := r.Group("/teachers")
	{
		// define class/teachers route
		teachers.POST("/:tid", create)
		teachers.GET("/:tid", read)
		teachers.GET("/:tid/:cid", readByID)
		teachers.DELETE("/:tid/:cid", deleteByID)
		teachers.PATCH("/:tid", update)
	}
}
