package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"go_portfolio/services"

	"github.com/gin-gonic/gin"
)

// ProjectController gère les requêtes HTTP pour les Projects
type ProjectController struct {
	service services.ProjectService
}

// NewProjectController crée une nouvelle instance de ProjectController
func NewProjectController(service services.ProjectService) *ProjectController {
	return &ProjectController{service: service}
}

// GetTheProjects gère la requête GET /Projects
func (c *ProjectController) GetTheProject(ctx *gin.Context) {
	// Logging de la requête
	fmt.Printf("Requête GET /Projects reçue\n")
	fmt.Printf("Adresse IP du serveur: %s\n", ctx.ClientIP())

	// Récupération des paramètres de pagination
	pageStr := ctx.Query("page")
	pageSizeStr := ctx.Query("page_size")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// Appel du service
	response, err := c.service.GetTheProject(ctx, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Configuration des headers
	ctx.Header("X-Total-Count", fmt.Sprintf("%d", response.Metadata.Total))

	// Envoi de la réponse
	ctx.JSON(http.StatusOK, response)
}
