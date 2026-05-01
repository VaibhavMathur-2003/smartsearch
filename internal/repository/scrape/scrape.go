package scrape

import (
	"context"
	domain "smartsearch/internal/domain/scrape"

	"gorm.io/gorm"
)

type WebsiteRepository struct {
	db *gorm.DB
}

func NewWebsiteRepository(db *gorm.DB) *WebsiteRepository {
	return &WebsiteRepository{db: db}
}

func (wr *WebsiteRepository) Create(ctx context.Context, d *domain.ScrapeData) error {
	m := ScrapeData{
		URL:      d.URL,
		RawHTML:  d.RawHTML,
		Content:  d.Content,
		Metadata: d.Metadata,
	}
	if err := wr.db.WithContext(ctx).Create(&m).Error; err != nil {
		return err
	}
	return nil
}
