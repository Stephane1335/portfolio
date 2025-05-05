package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"go_portfolio/models"
	"go_portfolio/services"

	"github.com/gin-gonic/gin"
)

// CitationController gère les requêtes HTTP pour les citations
type CitationController struct {
	service services.CitationService
}

// NewCitationController crée une nouvelle instance de CitationController
func NewCitationController(service services.CitationService) *CitationController {
	return &CitationController{service: service}
}

// GetTheCitations gère la requête GET /citations
func (c *CitationController) GetTheCitation(ctx *gin.Context) {
	// Logging de la requête
	fmt.Printf("Requête GET /citations reçue\n")

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
	response, err := c.service.GetTheCitation(ctx, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Configuration des headers
	ctx.Header("X-Total-Count", fmt.Sprintf("%d", response.Metadata.Total))

	// Envoi de la réponse
	ctx.JSON(http.StatusOK, response)
}

// UpdateCitation gère la requête pour mettre à jour une citation
func (c *CitationController) UpdateCitation(ctx *gin.Context) {
	// Validation de la méthode
	if ctx.Request.Method != http.MethodPut {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "méthode non autorisée"})
		return
	}

	// Récupération de l'ID depuis les paramètres
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID requis"})
		return
	}

	// Lecture du corps de la requête
	var citation models.Citation
	if err := ctx.ShouldBindJSON(&citation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "format de requête invalide"})
		return
	}
	// Appel au service
	response, err := c.service.UpdateCitation(id, &citation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Envoi de la réponse
	ctx.JSON(http.StatusOK, response)
}
