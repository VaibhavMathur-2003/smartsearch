package repo

import (
	"context"
	domain "smartsearch/internal/entities"
	"smartsearch/internal/models"

	"gorm.io/gorm"
)

type SummaryRepository struct {
	db *gorm.DB
}

func NewSummaryRepository(db *gorm.DB) *SummaryRepository {
	return &SummaryRepository{db: db}
}

func (wr *SummaryRepository) Create(ctx context.Context, d domain.Summary) error {
	m := models.Summary{
		Url:     d.Url,
		Summary: d.Summary,
	}
	if err := wr.db.WithContext(ctx).Create(&m).Error; err != nil {
		return err
	}
	return nil
}
