package service

import (
	"movie-api/internal/domain"
)

//go:generate mockgen -source=interface.go -destination=mocks/mock.go

type MovieService interface {
	List() (*[]domain.MovieResponse, error)
	View(id uint) (*domain.MovieResponse, error)
	Insert(input *domain.MovieRequest) (*domain.MovieResponse, error)
	Update(input *domain.MovieRequest) (*domain.MovieResponse, error)
	Delete(id uint) (*domain.MovieResponse, error)
}
