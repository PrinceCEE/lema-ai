package seeder

import (
	"context"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/princecee/lema-ai/internal/db/models"
)

type userRepo interface {
	CreateUser(ctx context.Context, u *models.User) error
}

type postRepo interface {
	CreatePost(ctx context.Context, p *models.Post) error
}

// Seed the database with 50 random users and 5 posts each
func Seed(userRepo userRepo, postRepo postRepo) {
	for i := 0; i < 50; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		user := &models.User{
			Name:     gofakeit.Name(),
			Username: gofakeit.Username(),
			Phone:    gofakeit.Phone(),
			Email:    gofakeit.Email(),
			Address: models.Address{
				Street:  gofakeit.StreetName(),
				City:    gofakeit.City(),
				State:   gofakeit.State(),
				Zipcode: gofakeit.Zip(),
			},
		}

		if err := userRepo.CreateUser(ctx, user); err != nil {
			panic(err)
		}

		for j := 0; j < 5; j++ {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			post := models.Post{
				Title:  gofakeit.Sentence(7),
				Body:   gofakeit.Paragraph(3, 7, 5, " "),
				UserID: user.ID,
			}

			if err := postRepo.CreatePost(ctx, &post); err != nil {
				panic(err)
			}
		}
	}

}
