package services

import (
	"context"
	"github.com/kheyssper/go-saas-microservice/pkg/platform_service/models"
	"github.com/kheyssper/go-saas-microservice/pkg/platform_service/repositories"
)

// PlatformService handles business logic for platforms.
type PlatformService struct {
	repo *repositories.PlatformRepository
}

// NewPlatformService creates a new instance of PlatformService.
func NewPlatformService(repo *repositories.PlatformRepository) *PlatformService {
	return &PlatformService{repo: repo}
}

// CreatePlatform creates a new platform.
func (s *PlatformService) CreatePlatform(ctx context.Context, platformName, platformSlug string, creatorID int) (*models.Platform, error) {
	p := &models.Platform{
		PlatformName: platformName,
		PlatformSlug: platformSlug,
		CreatorID:    creatorID,
		Status:       "active", // initial status

	}
	err := s.repo.Create(ctx, p)
	return p, err
}

// ListPlatforms returns all platforms.
func (s *PlatformService) ListPlatforms(ctx context.Context) ([]models.Platform, error) {
	return s.repo.FindAll(ctx)
}

// RunOrStopPlatform updates the status of a platform.
func (s *PlatformService) RunOrStopPlatform(ctx context.Context, id int, status string) error {
	return s.repo.UpdateStatus(ctx, id, status)
}

// DeletePlatform deletes a platform by ID.
func (s *PlatformService) DeletePlatform(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// GetPlatformByID retrieves a platform by ID.
func (s *PlatformService) GetPlatformByID(ctx context.Context, id int) (*models.Platform, error) {
	return s.repo.FindByID(ctx, id)
}
