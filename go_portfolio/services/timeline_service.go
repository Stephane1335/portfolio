package services

import (
	"context"
	"errors"
	"time"

	"go_portfolio/repository"
)

// TimelineService définit l'interface pour le service de Timelines
type TimelineService interface {
	GetTheTimeline(ctx context.Context, page, pageSize int) (*repository.TimelineResponse, error)
}

// TimelineService implémente l'interface TimelineService
type timelineService struct {
	repo repository.TimelineRepository
}

// NewTimelineService crée une nouvelle instance de TimelineService
func NewTimelineService(repo repository.TimelineRepository) TimelineService {
	return &timelineService{repo: repo}
}

// GetTheTimelines implémente la méthode pour récupérer les Timelines
func (s *timelineService) GetTheTimeline(
	ctx context.Context,
	page, pageSize int,
) (*repository.TimelineResponse, error) {
	// Validation des paramètres
	if page < 1 || pageSize < 1 {
		return nil, errors.New("paramètres de pagination invalides")
	}

	// Création du context avec timeout
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Appel du repository
	return s.repo.GetTheTimeline(ctx, page, pageSize)
}
