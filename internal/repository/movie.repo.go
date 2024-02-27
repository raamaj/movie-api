package repository

import (
	"errors"
	"gorm.io/gorm"
	"movie-api/internal/domain"
	"movie-api/internal/model"
	"strconv"
	"time"
)

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *movieRepository {
	return &movieRepository{db: db}
}

func (m *movieRepository) List() (*[]model.Movie, error) {
	var movies []model.Movie

	db := m.db.Model(&movies)

	checkMovie := db.Where("deleted = ? OR deleted is null", false).Find(&movies)
	if checkMovie.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	return &movies, nil
}

func (m *movieRepository) View(id uint) (*model.Movie, error) {
	var movie model.Movie

	movie.ID = id

	db := m.db.Model(&movie)

	checkMovie := db.Debug().First(&movie, "deleted=false")

	if checkMovie.RowsAffected < 1 {
		return nil, errors.New("Movie with ID " + strconv.Itoa(int(id)) + " not found ")
	}

	return &movie, nil
}

func (m *movieRepository) Insert(input *domain.MovieRequest, timeNow time.Time) (*model.Movie, error) {
	var movie model.Movie

	db := m.db.Model(&movie)

	movie.Title = input.Title
	movie.Description = input.Description
	movie.Rating = input.Rating
	movie.Image = input.Image
	movie.CreatedAt = timeNow
	movie.UpdatedAt = timeNow
	movie.Deleted = false

	addMovie := db.Debug().Create(&movie)
	if addMovie.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	return &movie, nil
}

func (m *movieRepository) Update(input *domain.MovieRequest, timeNow time.Time) (*model.Movie, error) {
	var movie model.Movie

	movie.ID = input.ID

	db := m.db.Model(&movie)

	checkMovie := db.Debug().First(&movie, "deleted=false")
	if checkMovie.RowsAffected < 1 {
		return nil, errors.New("Movie with ID " + strconv.Itoa(int(input.ID)) + " not found ")
	}

	movie.Title = input.Title
	movie.Description = input.Description
	movie.Rating = input.Rating
	movie.Image = input.Image
	movie.UpdatedAt = timeNow

	updateMovie := db.Debug().Updates(&movie)
	if updateMovie.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	return &movie, nil
}

func (m *movieRepository) Delete(id uint, timeNow time.Time) (*model.Movie, error) {
	var movie model.Movie

	db := m.db.Model(&movie)

	checkExistMovie := db.Debug().First(&movie, "deleted=false")
	if checkExistMovie.RowsAffected < 1 {
		return nil, errors.New("Movie with ID " + strconv.Itoa(int(id)) + " not found ")
	}

	movie.Deleted = true
	movie.UpdatedAt = timeNow

	updateMovie := db.Debug().Updates(&movie)
	if updateMovie.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	return &movie, nil
}
