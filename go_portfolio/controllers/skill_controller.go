package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"go_portfolio/services"

	"github.com/gin-gonic/gin"
)

// SkillController gère les requêtes HTTP pour les Skills
type SkillController struct {
	service services.SkillService
}

// NewSkillController crée une nouvelle instance de SkillController
func NewSkillController(service services.SkillService) *SkillController {
	return &SkillController{service: service}
}

// GetTheSkills gère la requête GET /Skills
func (c *SkillController) GetTheSkill(ctx *gin.Context) {
	// Logging de la requête
	fmt.Printf("Requête GET /Skills reçue\n")

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
	response, err := c.service.GetTheSkill(ctx, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Configuration des headers
	ctx.Header("X-Total-Count", fmt.Sprintf("%d", response.Metadata.Total))

	// Envoi de la réponse
	ctx.JSON(http.StatusOK, response)
}
