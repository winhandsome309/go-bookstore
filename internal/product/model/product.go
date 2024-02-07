package model

import (
	"time"
)

type Product struct {
	Id          string    `form:"id" json:"id" gorm:"unique,not null;index;primary_key"`
	Title       string    `form:"title" json:"title" `
	Description string    `form:"description" json:"description" `
	Price       int       `form:"price" json:"price" `
	Img_url     string    `form:"img_url" json:"img_url" `
	Quantity    int       `form:"quantity" json:"quantity" `
	LanguageId  int       `form:"language_id" json:"language_id" `
	CreatedAt   time.Time `form:"created_at" json:"created_at" `
	UpdatedAt   time.Time `form:"updated_at" json:"updated_at" `
	DeletedAt   time.Time `form:"deleted_at" json:"deleted_at" `
}

type UpdateProductReq struct {
	Title       *string `form:"title,omitempty"`
	Description *string `form:"description,omitempty"`
	Price       *int    `form:"price,omitempty"`
	Img_url     *string `form:"img_url,omitempty"`
	Quantity    *int    `form:"quantity,omitempty"`
}
