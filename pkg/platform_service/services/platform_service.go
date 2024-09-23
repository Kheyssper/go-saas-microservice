package services

import (
	"context"
	"github.com/kheyssper/go-saas-microservice/pkg/platform_service/models"
	"github.com/kheyssper/go-saas-microservice/pkg/platform_service/repositories"
)

// PlatformService lida com a lógica de negócios da plataforma.
type PlatformService struct {
	repo *repositories.PlatformRepository
}

// NewPlatformService cria uma nova instância de PlatformService.
func NewPlatformService(repo *repositories.PlatformRepository) *PlatformService {
	return &PlatformService{repo: repo}
}

// CreatePlatform cria uma nova plataforma.
func (s *PlatformService) CreatePlatform(ctx context.Context, platformName, platformSlug string, creatorID int) (*models.Platform, error) {
	p := &models.Platform{
		PlatformName: platformName,
		PlatformSlug: platformSlug,
		CreatorID:    creatorID,
		Status:       "active", // status inicial

	}
	err := s.repo.Create(ctx, p)
	return p, err
}

// ListPlatforms retorna todas as plataformas.
func (s *PlatformService) ListPlatforms(ctx context.Context) ([]models.Platform, error) {
	return s.repo.FindAll(ctx)
}

// RunOrStopPlatform atualiza o status de uma plataforma.
func (s *PlatformService) RunOrStopPlatform(ctx context.Context, id int, status string) error {
	return s.repo.UpdateStatus(ctx, id, status)
}

// DeletePlatform deleta uma plataforma por ID.
func (s *PlatformService) DeletePlatform(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// GetPlatformByID busca uma plataforma por ID.
func (s *PlatformService) GetPlatformByID(ctx context.Context, id int) (*models.Platform, error) {
	return s.repo.FindByID(ctx, id)
}
