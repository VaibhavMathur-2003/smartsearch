package repo

import (
	"context"
	"smartsearch/internal/domain"
	"smartsearch/internal/models"

	"gorm.io/gorm"
)

type WebsiteRepository struct {
	db *gorm.DB
}

func NewWebsiteRepository(db *gorm.DB) *WebsiteRepository {
	return &WebsiteRepository{db: db}
}

func (wr *WebsiteRepository) Create(ctx context.Context, d domain.Website) error {
	m := models.Website{
		Url:  d.Url,
		Text: d.Text,
	}
	if err := wr.db.WithContext(ctx).Create(&m).Error; err != nil {
		return err
	}
	return nil
}
