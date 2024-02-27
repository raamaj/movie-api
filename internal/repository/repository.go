package repository

import "gorm.io/gorm"

type Repositories struct {
	Movie MovieRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		Movie: NewMovieRepository(db),
	}
}
