package summary

import (
	"context"
	domain "smartsearch/internal/domain/summary"

	"gorm.io/gorm"
)

type SummaryRepository struct {
	db *gorm.DB
}

func NewSummaryRepository(db *gorm.DB) *SummaryRepository {
	return &SummaryRepository{db: db}
}

func (wr *SummaryRepository) Create(ctx context.Context, d *domain.SummaryResponse) error {
	m := Summary{
		URL:     d.URL,
		Summary: d.Summary,
	}
	if err := wr.db.WithContext(ctx).Create(&m).Error; err != nil {
		return err
	}
	return nil
}
