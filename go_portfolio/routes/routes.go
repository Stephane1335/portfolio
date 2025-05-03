package routes

import (
	"go_portfolio/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.GET("/citations", controllers.GetAllCitations())
}
