package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"go_portfolio/services"

	"github.com/gin-gonic/gin"
)

// TimelineController gère les requêtes HTTP pour les Timelines
type TimelineController struct {
	service services.TimelineService
}

// NewTimelineController crée une nouvelle instance de TimelineController
func NewTimelineController(service services.TimelineService) *TimelineController {
	return &TimelineController{service: service}
}

// GetTheTimelines gère la requête GET /Timelines
func (c *TimelineController) GetTheTimeline(ctx *gin.Context) {
	// Logging de la requête
	fmt.Printf("Requête GET /Timelines reçue\n")

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
	response, err := c.service.GetTheTimeline(ctx, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Envoi de la réponse
	ctx.JSON(http.StatusOK, response)
}
