package searchrepo

import (
	"context"
	domain "smartsearch/internal/domain/search"
	utils "smartsearch/internal/utils"

	"gorm.io/gorm"
)

type SearchRepository struct {
	db *gorm.DB
}

func NewSearchRepository(db *gorm.DB) *SearchRepository {
	return &SearchRepository{db: db}
}

func (ur *SearchRepository) Create(ctx context.Context, d *domain.SearchResponse) error {
	for _, item := range d.Results {
		m := Urls{
			Domain:  utils.ExtractDomain(item.URL),
			URL:     item.URL,
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
		Model(&Urls{}).
		Where("domain = ?", utils.ExtractDomain(website)).
		Pluck("url", &urls).Error

	if err != nil {
		return nil, err
	}

	return urls, nil
}
