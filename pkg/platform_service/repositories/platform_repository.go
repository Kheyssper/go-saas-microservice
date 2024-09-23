package repositories

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kheyssper/go-saas-microservice/pkg/platform_service/models"
)

// PlatformRepository handles database operations for the Platform entity
type PlatformRepository struct {
	db *pgxpool.Pool
}

// NewPlatformRepository creates a new instance of PlatformRepository.
func NewPlatformRepository(db *pgxpool.Pool) *PlatformRepository {
	return &PlatformRepository{db: db}
}

// Create inserts a new platform into the database
func (r *PlatformRepository) Create(ctx context.Context, p *models.Platform) error {
	query := `INSERT INTO platforms (platform_name, platform_slug, creator_id, status) 
              VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	return r.db.QueryRow(ctx, query, p.PlatformName, p.PlatformSlug, p.CreatorID, p.Status).
		Scan(&p.ID, &p.CreatedAt)
}

// FindAll returns all platforms.
func (r *PlatformRepository) FindAll(ctx context.Context) ([]models.Platform, error) {
	var platforms []models.Platform
	rows, err := r.db.Query(ctx, "SELECT id, platform_name, platform_slug, creator_id, status, created_at FROM platforms")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Platform
		if err := rows.Scan(&p.ID, &p.PlatformName, &p.PlatformSlug, &p.CreatorID, &p.Status, &p.CreatedAt); err != nil {
			return nil, err
		}
		platforms = append(platforms, p)
	}
	return platforms, nil
}

// UpdateStatus updates the status of a platform.
func (r *PlatformRepository) UpdateStatus(ctx context.Context, id int, status string) error {
	_, err := r.db.Exec(ctx, "UPDATE platforms SET status=$1 WHERE id=$2", status, id)
	return err
}

// Delete deletes a platform by ID.
func (r *PlatformRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "DELETE FROM platforms WHERE id=$1", id)
	return err
}

// FindByID searches for a platform by ID.
func (r *PlatformRepository) FindByID(ctx context.Context, id int) (*models.Platform, error) {
	var platform models.Platform
	err := r.db.QueryRow(ctx, "SELECT id, platform_name, platform_slug, creator_id, status, created_at FROM platforms WHERE id=$1", id).
		Scan(&platform.ID, &platform.PlatformName, &platform.PlatformSlug, &platform.CreatorID, &platform.Status, &platform.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &platform, nil
}
