package service

import (
	"movie-api/internal/domain"
	"movie-api/internal/model"
	"movie-api/internal/repository"
	"time"
)

type MovieServiceImpl struct {
	Repository repository.MovieRepository
}

func NewMovieService(movieRepo repository.MovieRepository) *MovieServiceImpl {
	return &MovieServiceImpl{
		Repository: movieRepo,
	}
}

func (m *MovieServiceImpl) List() (*[]domain.MovieResponse, error) {
	results, err := m.Repository.List()
	if err != nil {
		return nil, err
	}
	return convertToResponseList(results), nil
}

func (m *MovieServiceImpl) View(id uint) (*domain.MovieResponse, error) {
	result, err := m.Repository.View(id)
	if err != nil {
		return nil, err
	}
	return convertToResponse(result), nil
}

func (m *MovieServiceImpl) Insert(input *domain.MovieRequest) (*domain.MovieResponse, error) {
	result, err := m.Repository.Insert(input, time.Now())
	if err != nil {
		return nil, err
	}
	return convertToResponse(result), nil
}

func (m *MovieServiceImpl) Update(input *domain.MovieRequest) (*domain.MovieResponse, error) {
	result, err := m.Repository.Update(input, time.Now())
	if err != nil {
		return nil, err
	}
	return convertToResponse(result), nil
}

func (m *MovieServiceImpl) Delete(id uint) (*domain.MovieResponse, error) {
	result, err := m.Repository.Delete(id, time.Now())
	if err != nil {
		return nil, err
	}
	return convertToResponse(result), nil
}

func convertToResponse(input *model.Movie) *domain.MovieResponse {
	return &domain.MovieResponse{
		ID:          input.ID,
		Title:       input.Title,
		Description: input.Description,
		Rating:      input.Rating,
		Image:       input.Image,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
	}
}

func convertToResponseList(inputs *[]model.Movie) *[]domain.MovieResponse {
	var result []domain.MovieResponse
	for _, input := range *inputs {
		result = append(result, domain.MovieResponse{
			ID:          input.ID,
			Title:       input.Title,
			Description: input.Description,
			Rating:      input.Rating,
			Image:       input.Image,
			CreatedAt:   input.CreatedAt,
			UpdatedAt:   input.UpdatedAt,
		})
	}

	return &result
}
