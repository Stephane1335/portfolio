package services

import (
	"context"
	"errors"
	"time"

	"go_portfolio/models"
	"go_portfolio/repository"
)

// CitationService définit l'interface pour le service de citations
type CitationService interface {
	GetTheCitation(ctx context.Context, page, pageSize int) (*repository.CitationResponse, error)
	UpdateCitation(id string, citation *models.Citation) (*repository.CitationUpdateResponse, error)
}

// citationService implémente l'interface CitationService
type citationService struct {
	repo repository.CitationRepository
}

// NewCitationService crée une nouvelle instance de citationService
func NewCitationService(repo repository.CitationRepository) CitationService {
	return &citationService{repo: repo}
}

// GetTheCitations implémente la méthode pour récupérer les citations
func (s *citationService) GetTheCitation(
	ctx context.Context,
	page, pageSize int,
) (*repository.CitationResponse, error) {
	// Validation des paramètres
	if page < 1 || pageSize < 1 {
		return nil, errors.New("paramètres de pagination invalides")
	}

	// Création du context avec timeout
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Appel du repository
	return s.repo.GetTheCitation(ctx, page, pageSize)
}

// UpdateCitation implémente la méthode pour mettre à jour une citation
func (s *citationService) UpdateCitation(
	id string,
	citation *models.Citation,
) (*repository.CitationUpdateResponse, error) {
	// Validation des paramètres
	if id == "" {
		return nil, errors.New("ID requis")
	}
	if citation == nil {
		return nil, errors.New("citation requise")
	}

	// Appel au repository
	return s.repo.UpdateCitation(id, citation)
}
