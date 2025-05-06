package services

import (
	"context"
	"errors"
	"time"

	"go_portfolio/repository"
)

// WhyHireMeService définit l'interface pour le service de WhyHireMes
type WhyHireMeService interface {
	GetTheWhyHireMe(ctx context.Context, page, pageSize int) (*repository.WhyHireMeResponse, error)
}

// WhyHireMeService implémente l'interface WhyHireMeService
type whyhiremeService struct {
	repo repository.WhyHireMeRepository
}

// NewWhyHireMeService crée une nouvelle instance de WhyHireMeService
func NewWhyHireMeService(repo repository.WhyHireMeRepository) WhyHireMeService {
	return &whyhiremeService{repo: repo}
}

// GetTheWhyHireMes implémente la méthode pour récupérer les WhyHireMes
func (s *whyhiremeService) GetTheWhyHireMe(
	ctx context.Context,
	page, pageSize int,
) (*repository.WhyHireMeResponse, error) {
	// Validation des paramètres
	if page < 1 || pageSize < 1 {
		return nil, errors.New("paramètres de pagination invalides")
	}

	// Création du context avec timeout
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Appel du repository
	return s.repo.GetTheWhyHireMe(ctx, page, pageSize)
}
