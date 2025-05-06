package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"go_portfolio/services"

	"github.com/gin-gonic/gin"
)

// AboutMeController gère les requêtes HTTP pour les AboutMes
type AboutMeController struct {
	service services.AboutMeService
}

// NewAboutMeController crée une nouvelle instance de AboutMeController
func NewAboutMeController(service services.AboutMeService) *AboutMeController {
	return &AboutMeController{service: service}
}

// GetTheAboutMes gère la requête GET /AboutMes
func (c *AboutMeController) GetTheAboutMe(ctx *gin.Context) {
	// Logging de la requête
	fmt.Printf("Requête GET /AboutMes reçue\n")

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
	response, err := c.service.GetTheAboutMe(ctx, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Configuration des headers
	ctx.Header("X-Total-Count", fmt.Sprintf("%d", response.Metadata.Total))

	// Envoi de la réponse
	ctx.JSON(http.StatusOK, response)
}
