package api

import (
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"movie-api/internal/repository"
	"movie-api/internal/service"
)

type Repository struct {
	movieRepo repository.MovieRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{movieRepo: repository.NewMovieRepository(db)}
}

type Services struct {
	userService service.MovieService
}

func NewServices(repo *Repository) *Services {
	return &Services{
		userService: service.NewMovieService(repo.movieRepo),
	}
}

func NewApi(app *fiber.App, services *Services) {
	app.Use(recover2.New())
	app.Use(requestid.New())

	userHandler := NewMovieHandler(services.userService)

	app.Get("/Movies", userHandler.movieListHandler)
	app.Get("/Movies/:id", userHandler.movieViewHandler)
	app.Post("/Movies", userHandler.movieInsertHandler)
	app.Patch("/Movies/:id", userHandler.movieUpdateHandler)
	app.Delete("/Movies/:id", userHandler.movieDeleteHandler)

	log.Fatal(app.Listen(viper.GetString("APP_ADDRESS")))
}
