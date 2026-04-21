package pipeline

import (
	"context"
	"fmt"
	"log"
	"smartsearch/internal/domain"
	"smartsearch/internal/pipeline/scrape"
	"smartsearch/internal/pipeline/searx"
	repo "smartsearch/internal/repository"
	"sync"
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
	var wg sync.WaitGroup
	urlData, err := searx.CallSearx(website)
	if err != nil {
		fmt.Println("No urls ", err)
	}
	err = p.urlRepo.Create(ctx, urlData)
	if err != nil {
		log.Println("failed to save:", err)
	}
	for _, u := range urlData.UrlData {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			scrapeText := scrape.Scrape(url)
			if scrapeText == nil {
				log.Println("empty scrape:", url)
				return
			}
			m := domain.Website{
				Url:  url,
				Text: *scrapeText,
			}
			err := p.websiteRepo.Create(ctx, m)
			if err != nil {
				log.Println("failed to save:", err)
			}
		}(u.Url)
	}
	wg.Wait()

}
