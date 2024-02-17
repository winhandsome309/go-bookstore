package model

import (
	"time"

	"gorm.io/gorm"
)

type Shipping struct {
	Id               int            `form:"id" json:"id"`
	OrderId          int            `form:"order_id" json:"order_id"`
	FirstName        string         `form:"first_name" json:"first_name"`
	LastName         string         `form:"last_name" json:"last_name"`
	Email            string         `form:"email" json:"email"`
	Phone            string         `form:"phone" json:"phone"`
	Address          string         `form:"address" json:"address"`
	ShippingProvince string         `form:"shipping_province" json:"shipping_province"`
	ShippingDistrict string         `form:"shipping_district" json:"shipping_district"`
	ShippingWard     string         `form:"shipping_ward" json:"shipping_ward"`
	CreatedAt        time.Time      `form:"created_at" json:"created_at"`
	UpdatedAt        time.Time      `form:"updated_at" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `form:"deleted_at" json:"deleted_at"`
}
