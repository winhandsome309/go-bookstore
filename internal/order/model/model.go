package model

import (
	"database/sql/driver"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

// type listOrderLine []OrderLine

// func (list *listOrderLine) Scan(value interface{}) error {
// 	bytes, ok := value.(string)
// 	if !ok {
// 		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
// 	}
// 	err := json.Unmarshal([]byte(bytes), list)
// 	// err := json.Unmarshal([]byte(value.(string)), list)
// 	return err
// }

// func (list listOrderLine) Value() (driver.Value, error) {
// 	data, _ := json.Marshal(list)
// 	return string(data), nil
// }

type listOrderLine []int

func (list *listOrderLine) Scan(value interface{}) error {
	s := value.(string)
	if s == "{}" {
		*list = []int{}
		return nil
	}
	s = strings.Trim(s, "{}")
	arr := strings.Split(s, ",")
	for _, val := range arr {
		temp, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		*list = append(*list, temp)
	}
	return nil
}

func (list listOrderLine) Value() (driver.Value, error) {
	s := "{"
	for i, val := range list {
		s += strconv.Itoa(val)
		if i < len(list)-1 {
			s += ","
		}
	}
	s += "}"
	return s, nil
}

type Order struct {
	Id         int `json:"id" gorm:"unique,not null;index;primaryKey"`
	CustomerID int `json:"customer_id"`
	TotalItem  int `json:"total_item"`
	TotalPrice int `json:"total_price"`
	// Lines      listOrderLine  `json:"lines" gorm:"type:text"`
	Lines     listOrderLine  `json:"lines" gorm:"type:integer[]"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type OrderLine struct {
	Id        int            `form:"id" json:"id"`
	ProductId int            `form:"product_id" json:"product_id"`
	OrderId   int            `form:"order_id" json:"order_id"`
	Quantity  int            `form:"quantity" json:"quantity"`
	Price     int            `form:"price" json:"price"`
	CreatedAt time.Time      `form:"created_at" json:"created_at"`
	UpdatedAt time.Time      `form:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `form:"deleted_at" json:"deleted_at"`
}

type OrderLineProduct struct {
	Id            int            `json:"id"`
	ProductId     int            `json:"product_id"`
	ProductTitle  string         `json:"product_title"`
	ProductImgUrl string         `json:"product_img_url"`
	OrderId       int            `json:"order_id"`
	Quantity      int            `json:"quantity"`
	Price         int            `json:"price"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

type OrderResponse struct {
	Id         int          `json:"id" gorm:"unique,not null;index;primaryKey"`
	CustomerID int          `json:"customer_id"`
	TotalItem  int          `json:"total_item"`
	TotalPrice int          `json:"total_price"`
	Lines      []*OrderLine `json:"lines"`
}
