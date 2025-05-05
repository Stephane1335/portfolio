package routes

import (
	"go_portfolio/controllers"
	"go_portfolio/repository"
	"go_portfolio/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// SetupRoutes configure les routes de l'application
func SetupRoutes(router *gin.Engine, db *mongo.Database) {
	// Création des dépendances
	repo := repository.NewCitationRepository(db)
	service := services.NewCitationService(repo)
	controller := controllers.NewCitationController(service)

	// Déclaration des routes
	router.GET("/citations", controller.GetTheCitation)
	router.PUT("/citations/:id", controller.UpdateCitation)
}
