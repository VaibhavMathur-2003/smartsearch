package repo

import (
	"context"
	"smartsearch/internal/domain"
	"smartsearch/internal/models"
	utils "smartsearch/internal/utils"

	"gorm.io/gorm"
)

type UrlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) *UrlRepository {
	return &UrlRepository{db: db}
}

func (ur *UrlRepository) Create(ctx context.Context, d *domain.Searx) error {
	for _, item := range d.UrlData {
		m := models.Urls{
			Domain:  utils.ExtractDomain(item.Url), // optional helper
			Url:     item.Url,
			Title:   item.Title,
			Content: item.Content,
		}

		if err := ur.db.WithContext(ctx).Create(&m).Error; err != nil {
			return err
		}
	}
	return nil
}

func (ur *UrlRepository) Get(ctx context.Context, website string) ([]string, error) {
	var urls []string

	err := ur.db.WithContext(ctx).
		Model(&models.Urls{}).
		Where("domain = ?", utils.ExtractDomain(website)).
		Pluck("url", &urls).Error

	if err != nil {
		return nil, err
	}

	return urls, nil
}
