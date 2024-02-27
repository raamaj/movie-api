package domain

import (
	"errors"
	"movie-api/pkg/validator"
	"time"
)

type MovieRequest struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Rating      float64 `json:"rating" validate:"min=0,max=10"`
	Image       string  `json:"image" validate:"required"`
}

func (m *MovieRequest) validate() []error {
	errs := validator.ValidateStruct(m)

	return errs
}

func (m *MovieRequest) ValidateInsert() []error {
	return m.validate()
}

func (m *MovieRequest) ValidateUpdate() []error {
	if m.ID == 0 {
		return []error{errors.New("Required ID field!")}
	}

	return m.validate()
}

type MovieResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rating      float64   `json:"rating"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
