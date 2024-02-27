package model

import "time"

type Movie struct {
	ID          uint      `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"type:varchar; not null"`
	Description string    `json:"description" gorm:"type:varchar; not null"`
	Rating      float64   `json:"rating" gorm:"type:float; not null"`
	Image       string    `json:"image" gorm:"type:varchar; not null"`
	Deleted     bool      `json:"deleted" gorm:"type:boolean; default:null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
