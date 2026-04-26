package pipeline

import (
	"context"
	"fmt"
	"log"
	domain "smartsearch/internal/entities"
	"smartsearch/internal/pipeline/llm"
	"smartsearch/internal/pipeline/scrape"
	"smartsearch/internal/pipeline/searx"
	repo "smartsearch/internal/repository"
	searchrepo "smartsearch/internal/repository/search"
	"sync"
)

type Pipeline struct {
	urlRepo     *searchrepo.SearchRepository
	websiteRepo *repo.WebsiteRepository
	summaryRepo *repo.SummaryRepository
}

func NewPipeline(urlRepo *searchrepo.SearchRepository, websiteRepo *repo.WebsiteRepository, summaryRepo *repo.SummaryRepository) *Pipeline {
	return &Pipeline{urlRepo: urlRepo,
		websiteRepo: websiteRepo, summaryRepo: summaryRepo}
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
			wm := domain.Website{
				Url:  url,
				Text: *scrapeText,
			}
			err := p.websiteRepo.Create(ctx, wm)
			if err != nil {
				log.Println("failed to save:", err)
			}

			systemPrompt := "You are a content summarizer. Create a summary of the given content in 50-100 words english paragraph Make sure all imprtant contents are mentioned in the summary"
			userPrompt := fmt.Sprintf("Summarize the following content \n ---CONTENT START ---\n %s \n ---CONTENT END---", *scrapeText)

			summaryText, err := llm.LlmCall(ctx, "deepseek-r1:1.5b", systemPrompt, userPrompt)
			fmt.Println(summaryText, "SUMMARY")
			sm := domain.Summary{
				Url:     url,
				Summary: summaryText,
			}
			err = p.summaryRepo.Create(ctx, sm)
			if err != nil {
				log.Println("failed to save:", err)
			}

		}(u.Url)
	}
	wg.Wait()

}
