package searchrepo

import (
	"context"
	domain "smartsearch/internal/entities"
	"smartsearch/internal/models"
	utils "smartsearch/internal/utils"

	"gorm.io/gorm"
)

type SearchRepository struct {
	db *gorm.DB
}

func NewSearchRepository(db *gorm.DB) *SearchRepository {
	return &SearchRepository{db: db}
}

func (ur *SearchRepository) Create(ctx context.Context, d *domain.Searx) error {
	for _, item := range d.UrlData {
		m := models.Urls{
			Domain:  utils.ExtractDomain(item.Url),
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

func (ur *SearchRepository) Get(ctx context.Context, website string) ([]string, error) {
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
