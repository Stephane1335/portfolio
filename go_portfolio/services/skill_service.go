package services

import (
	"context"
	"errors"
	"time"

	"go_portfolio/repository"
)

// SkillService définit l'interface pour le service de Skills
type SkillService interface {
	GetTheSkill(ctx context.Context, page, pageSize int) (*repository.SkillResponse, error)
}

// SkillService implémente l'interface SkillService
type skillService struct {
	repo repository.SkillRepository
}

// NewSkillService crée une nouvelle instance de SkillService
func NewSkillService(repo repository.SkillRepository) SkillService {
	return &skillService{repo: repo}
}

// GetTheSkills implémente la méthode pour récupérer les Skills
func (s *skillService) GetTheSkill(
	ctx context.Context,
	page, pageSize int,
) (*repository.SkillResponse, error) {
	// Validation des paramètres
	if page < 1 || pageSize < 1 {
		return nil, errors.New("paramètres de pagination invalides")
	}

	// Création du context avec timeout
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Appel du repository
	return s.repo.GetTheSkill(ctx, page, pageSize)
}
