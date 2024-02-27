package repository

import (
	"movie-api/internal/domain"
	"movie-api/internal/model"
	"time"
)

//go:generate mockgen -source=interface.go -destination=mocks/mock.go

type MovieRepository interface {
	List() (*[]model.Movie, error)
	View(id uint) (*model.Movie, error)
	Insert(input *domain.MovieRequest, timeNow time.Time) (*model.Movie, error)
	Update(input *domain.MovieRequest, timeNow time.Time) (*model.Movie, error)
	Delete(id uint, timeNow time.Time) (*model.Movie, error)
}
