package models

import "time"

type Platform struct {
	ID           int       `json:"id"`
	PlatformName string    `json:"platform_name"`
	PlatformSlug string    `json:"platform_slug"`
	CreatorID    int       `json:"creator_id"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}
