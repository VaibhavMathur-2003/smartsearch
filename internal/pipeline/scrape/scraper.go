package scrape

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func Scrape(url string) *string {
	c := colly.NewCollector()

	var result *string

	c.OnResponse(func(r *colly.Response) {
		cleaned := clean(string(r.Body))
		result = &cleaned
	})

	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return result
}

func clean(html string) string {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	doc.Find("script, style").Remove()
	return doc.Text()
}
