package services

import (
	"context"
	"errors"
	"time"

	"go_portfolio/repository"
)

// ProjectService définit l'interface pour le service de Projects
type ProjectService interface {
	GetTheProject(ctx context.Context, page, pageSize int) (*repository.ProjectResponse, error)
}

// ProjectService implémente l'interface ProjectService
type projectService struct {
	repo repository.ProjectRepository
}

// NewProjectService crée une nouvelle instance de ProjectService
func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

// GetTheProjects implémente la méthode pour récupérer les Projects
func (s *projectService) GetTheProject(
	ctx context.Context,
	page, pageSize int,
) (*repository.ProjectResponse, error) {
	// Validation des paramètres
	if page < 1 || pageSize < 1 {
		return nil, errors.New("paramètres de pagination invalides")
	}

	// Création du context avec timeout
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Appel du repository
	return s.repo.GetTheProject(ctx, page, pageSize)
}
