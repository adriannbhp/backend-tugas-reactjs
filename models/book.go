package models

import (
	"time"
)

// Book struct
type Book struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url" binding:"required,url"`
	ReleaseYear int       `json:"release_year" binding:"required,gte=1980,lte=2021"`
	Price       string    `json:"price" binding:"required"`
	TotalPage   int       `json:"total_page" binding:"required"`
	Thickness   string    `json:"thickness"`
	UserID      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
