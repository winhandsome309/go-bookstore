package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id         int            `json:"id" gorm:"unique,not null;index;primaryKey"`
	Email      string         `json:"email"`
	Password   string         `json:"password"`
	Role       string         `json:"role"`
	Balance    int            `json:"balance"`
	Created_at time.Time      `json:"created_at"`
	Updated_at time.Time      `json:"updated_at"`
	Deleted_at gorm.DeletedAt `json:"deleted_at"`
}

type UserReq struct {
	Email    string `form:"email,omitempty"`
	Password string `form:"password,omitempty"`
	Role     string `form:"role,omitempty"`
}

type UserLogin struct {
	Email    string `form:"email,omitempty"`
	Password string `form:"password,omitempty"`
}
