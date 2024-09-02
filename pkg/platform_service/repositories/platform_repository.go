package repositories

import (
	"context"
	"github.com/kheyssper/go-saas-microservice/pkg/platform_service/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

// PlatformRepository lida com as operações do banco de dados para a entidade Platform
type PlatformRepository struct {
	db *pgxpool.Pool
}

// NewPlatformRepository cria uma nova instância de PlatformRepository.
func NewPlatformRepository(db *pgxpool.Pool) *PlatformRepository {
	return &PlatformRepository{db: db}
}

// Create insere uma nova plataforma no banco de dados
func (r *PlatformRepository) Create(ctx context.Context, p *Platform) error {
	query := `INSERT INTO platforms (platform_name, platform_slug, creator_id, status) 
              VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	return r.db.QueryRow(ctx, query, p.PlatformName, p.PlatformSlug, p.CreatorID, p.Status).
		Scan(&p.ID, &p.CreatedAt)
}

// FindAll retorna todas as plataformas.
func (r *PlatformRepository) FindAll(ctx context.Context) ([]Platform, error) {
	var platforms []Platform
	rows, err := r.db.Query(ctx, "SELECT id, platform_name, platform_slug, creator_id, status, created_at FROM platforms")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Platform
		if err := rows.Scan(&p.ID, &p.PlatformName, &p.PlatformSlug, &p.CreatorID, &p.Status, &p.CreatedAt); err != nil {
			return nil, err
		}
		platforms = append(platforms, p)
	}
	return platforms, nil
}

// UpdateStatus atualiza o status de uma plataforma.
func (r *PlatformRepository) UpdateStatus(ctx context.Context, id int, status string) error {
	_, err := r.db.Exec(ctx, "UPDATE platforms SET status=$1 WHERE id=$2", status, id)
	return err
}
