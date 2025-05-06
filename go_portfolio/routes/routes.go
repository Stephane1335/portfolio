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

	// Création des dépendances citations
	repo_citation := repository.NewCitationRepository(db)
	service_citation := services.NewCitationService(repo_citation)
	controller_citation := controllers.NewCitationController(service_citation)

	// Creation des dependances About Me
	repo_aboutme := repository.NewAboutMeRepository(db)
	service_aboutme := services.NewAboutMeService(repo_aboutme)
	controller_aboutme := controllers.NewAboutMeController(service_aboutme)

	// Creation des dependances projects
	repo_project := repository.NewProjectRepository(db)
	service_project := services.NewProjectService(repo_project)
	controller_project := controllers.NewProjectController(service_project)

	// Creation des dependances Skill
	repo_skill := repository.NewSkillRepository(db)
	service_skill := services.NewSkillService(repo_skill)
	controller_skill := controllers.NewSkillController(service_skill)

	// Creation des dependances Timeline
	repo_timeline := repository.NewTimelineRepository(db)
	service_timeline := services.NewTimelineService(repo_timeline)
	controller_timeline := controllers.NewTimelineController(service_timeline)

	// Creation des dependances Why Hire Me
	repo_whyhireme := repository.NewWhyHireMeRepository(db)
	service_whyhireme := services.NewWhyHireMeService(repo_whyhireme)
	controller_whyhireme := controllers.NewWhyHireMeController(service_whyhireme)

	// Déclaration des routes
	router.GET("/citations", controller_citation.GetTheCitation)
	router.GET("/aboutme", controller_aboutme.GetTheAboutMe)
	router.GET("/projects", controller_project.GetTheProject)
	router.GET("/skills", controller_skill.GetTheSkill)
	router.GET("/timelines", controller_timeline.GetTheTimeline)
	router.GET("/whyhireme", controller_whyhireme.GetTheWhyHireMe)
	router.PUT("/citations/:id", controller_citation.UpdateCitation)
}
