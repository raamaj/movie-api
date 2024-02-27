package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"movie-api/internal/handler/api"
	"movie-api/pkg/dotenv"
	"movie-api/pkg/logger"
	"movie-api/pkg/postgres"
)

// @title Movie APP
// @version 1.0
// @description MOVIE APP is a web application for film management. This application provides basic CRUD functionality.
// @contact.name Rama Jayapermana
// @contact.email jayapermanarama@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	err := dotenv.Viper()

	if err != nil {
		logger.Error("error", zap.Error(err))
	}

	db, err := postgres.NewClient()

	if err != nil {
		logger.Error("error", zap.Error(err))
	}

	repo := api.NewRepository(db)
	service := api.NewServices(repo)
	app := fiber.New()
	api.NewApi(app, service)
}
