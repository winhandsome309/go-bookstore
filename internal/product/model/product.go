package model

import (
	"time"
)

type Product struct {
	Id          string    `json:"id" gorm:"unique,not null;index;primary_key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Img_url     string    `json:"img_url"`
	Quantity    int       `json:"quantity"`
	LanguageId  int       `json:"language_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
