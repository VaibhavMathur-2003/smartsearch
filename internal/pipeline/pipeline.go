package pipeline

import (
	"context"
	"fmt"
	"log"
	"smartsearch/internal/pipeline/searx"
	repo "smartsearch/internal/repository"
)

type Pipeline struct {
	urlRepo     *repo.UrlRepository
	websiteRepo *repo.WebsiteRepository
}

func NewPipeline(urlRepo *repo.UrlRepository, websiteRepo *repo.WebsiteRepository) *Pipeline {
	return &Pipeline{urlRepo: urlRepo,
		websiteRepo: websiteRepo}
}

func (p *Pipeline) RunPipeline(ctx context.Context, website string) {
	urlData, err := searx.CallSearx(website)
	if err != nil {
		fmt.Println("No urls ", err)
	}
	err = p.urlRepo.Create(ctx, urlData)
	if err != nil {
		log.Println("failed to save:", err)
	}

}
