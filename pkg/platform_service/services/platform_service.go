package service

import "context"

// PlatformService lida com a lógica de negócios da plataforma.
type PlatformService struct {
	repo *PlatformRepository
}

// NewPlatformService cria uma nova instância de PlatformService.
func NewPlatformService(repo *PlatformRepository) *PlatformService {
	return &PlatformService{repo: repo}
}

// CreatePlatform cria uma nova plataforma.
func (s *PlatformService) CreatePlatform(ctx context.Context, platformName, platformSlug string, creatorID int) (*Platform, error) {
	p := &Platform{
		PlatformName: platformName,
		PlatformSlug: platformSlug,
		CreatorID:    creatorID,
		Status:       "active", // status inicial

	}
	err := s.repo.Create(ctx, p)
	return p, err
}

// ListPlatforms retorna todas as plataformas.
func (s *PlatformService) ListPlatforms(ctx context.Context) ([]Platform, error) {
	return s.repo.FindAll(ctx)
}

// RunOrStopPlatform atualiza o status de uma plataforma.
func (s *PlatformService) RunOrStopPlatform(ctx context.Context, id int, status string) error {
	return s.repo.UpdateStatus(ctx, id, status)
}
