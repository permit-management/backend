package seeder

import (
	"context"

	"github.com/permit-management/backend/internal/repository"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

type TagSeeder struct{}

func init() {
	RegisterSeeder(TagSeeder{})
}

func (TagSeeder) Name() string {
	return "TagSeeder"
}

func (TagSeeder) Seed(engine *gorm.DB) (err error) {
	ctx := context.WithValue(context.Background(), "username", "seeder")
	repo := repository.NewTagRepository(ctx, engine)
	_, err = repo.CreateTag(gofakeit.Gamertag())
	return
}
