package services

import (
	"context"
	"errors"
	"time"

	"go_portfolio/repository"
)

// AboutMeService définit l'interface pour le service de AboutMes
type AboutMeService interface {
	GetTheAboutMe(ctx context.Context, page, pageSize int) (*repository.AboutMeResponse, error)
}

// AboutMeService implémente l'interface AboutMeService
type aboutmeService struct {
	repo repository.AboutMeRepository
}

// NewAboutMeService crée une nouvelle instance de AboutMeService
func NewAboutMeService(repo repository.AboutMeRepository) AboutMeService {
	return &aboutmeService{repo: repo}
}

// GetTheAboutMes implémente la méthode pour récupérer les AboutMes
func (s *aboutmeService) GetTheAboutMe(
	ctx context.Context,
	page, pageSize int,
) (*repository.AboutMeResponse, error) {
	// Validation des paramètres
	if page < 1 || pageSize < 1 {
		return nil, errors.New("paramètres de pagination invalides")
	}

	// Création du context avec timeout
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Appel du repository
	return s.repo.GetTheAboutMe(ctx, page, pageSize)
}
