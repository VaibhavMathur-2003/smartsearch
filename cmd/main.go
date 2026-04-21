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

	err := db.AutoMigrate(&models.Urls{})
	if err != nil {
		log.Fatal("migration failed:", err)
	}
	err = db.AutoMigrate(&models.Website{})
	if err != nil {
		log.Fatal("migration failed:", err)
	}
	urlRpo := repo.NewUrlRepository(db)
	websiteRepo := repo.NewWebsiteRepository(db)

	pipeline := pipeline.NewPipeline(urlRpo, websiteRepo)
	args := os.Args[1]
	pipeline.RunPipeline(ctx, args)

}
