package repository

import (
	"go-bookstore/internal/location/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ILocationRepository interface {
	GetAllProvince(c *gin.Context) (*[]model.Provinces, error)
	GetDistrictByProvinceID(c *gin.Context, provinceId string) (*[]model.Districts, error)
	GetWardByDistrictID(c *gin.Context, districtId string) (*[]model.Wards, error)
}

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{db: db}
}

func (r *LocationRepository) GetAllProvince(c *gin.Context) (*[]model.Provinces, error) {
	var provinces []model.Provinces
	err := r.db.Raw("SELECT DISTINCT(province), province_code FROM locations").Find(&provinces).Error
	return &provinces, err
}

func (r *LocationRepository) GetDistrictByProvinceID(c *gin.Context, provinceId string) (*[]model.Districts, error) {
	var districts []model.Districts
	err := r.db.Raw("SELECT DISTINCT(district), district_code FROM locations WHERE province_code = ?", provinceId).Find(&districts).Error
	return &districts, err
}

func (r *LocationRepository) GetWardByDistrictID(c *gin.Context, districtId string) (*[]model.Wards, error) {
	var wards []model.Wards
	err := r.db.Raw("SELECT DISTINCT(ward), ward_code FROM locations WHERE district_code = ?", districtId).Find(&wards).Error
	return &wards, err
}
