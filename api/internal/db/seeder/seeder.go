package seeder

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/princecee/lema-ai/internal/db/models"
)

type userRepo interface {
	CreateUser(u *models.User) error
}

type postRepo interface {
	CreatePost(p *models.Post) error
}

// Seed the database with 50 random users and 5 posts each
func Seed(userRepo userRepo, postRepo postRepo) {
	for i := 0; i < 50; i++ {
		user := &models.User{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Username:  gofakeit.Username(),
			Phone:     gofakeit.Phone(),
			Email:     gofakeit.Email(),
			Address: models.Address{
				Street:  gofakeit.StreetName(),
				City:    gofakeit.City(),
				State:   gofakeit.State(),
				Zipcode: gofakeit.Zip(),
			},
		}

		if err := userRepo.CreateUser(user); err != nil {
			panic(err)
		}

		for j := 0; j < 5; j++ {
			post := models.Post{
				Title:  gofakeit.Sentence(7),
				Body:   gofakeit.Paragraph(3, 7, 5, " "),
				UserID: user.ID,
			}

			if err := postRepo.CreatePost(&post); err != nil {
				panic(err)
			}
		}
	}

}
