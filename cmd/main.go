package main

import (
	"context"
	"log"
	"os"
	"smartsearch/internal/core"
	"smartsearch/internal/models"
	"smartsearch/internal/pipeline"
	repo "smartsearch/internal/repository"
)

func main() {
	ctx := context.Background()
	db := core.InitDB()

	err := db.AutoMigrate(
		&models.Urls{},
		&models.Website{},
		&models.Summary{},
	)
	if err != nil {
		log.Fatal("migration failed:", err)
	}
	urlRpo := repo.NewUrlRepository(db)
	websiteRepo := repo.NewWebsiteRepository(db)
	summaryRepo := repo.NewSummaryRepository(db)

	pipeline := pipeline.NewPipeline(urlRpo, websiteRepo, summaryRepo)
	args := os.Args[1]
	pipeline.RunPipeline(ctx, args)

}
