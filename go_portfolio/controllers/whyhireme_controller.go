package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"go_portfolio/services"

	"github.com/gin-gonic/gin"
)

// WhyHireMeController gère les requêtes HTTP pour les WhyHireMes
type WhyHireMeController struct {
	service services.WhyHireMeService
}

// NewWhyHireMeController crée une nouvelle instance de WhyHireMeController
func NewWhyHireMeController(service services.WhyHireMeService) *WhyHireMeController {
	return &WhyHireMeController{service: service}
}

// GetTheWhyHireMes gère la requête GET /WhyHireMes
func (c *WhyHireMeController) GetTheWhyHireMe(ctx *gin.Context) {
	// Logging de la requête
	fmt.Printf("Requête GET /WhyHireMes reçue\n")

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
	response, err := c.service.GetTheWhyHireMe(ctx, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Envoi de la réponse
	ctx.JSON(http.StatusOK, response)
}
