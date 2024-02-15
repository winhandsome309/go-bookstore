package service

import (
	"go-bookstore/internal/location/model"
	"go-bookstore/internal/location/repository"

	"github.com/gin-gonic/gin"
)

type ILocationService interface {
	GetAllProvince(c *gin.Context) (*[]model.Provinces, error)
	GetDistrictByProvinceID(c *gin.Context, provinceId string) (*[]model.Districts, error)
	GetWardByDistrictID(c *gin.Context, districtId string) (*[]model.Wards, error)
}

type LocationService struct {
	repo repository.ILocationRepository
}

func NewLocationService(repo repository.ILocationRepository) *LocationService {
	return &LocationService{repo: repo}
}

func (r *LocationService) GetAllProvince(c *gin.Context) (*[]model.Provinces, error) {
	provinces, err := r.repo.GetAllProvince(c)
	if err != nil {
		return nil, err
	}
	return provinces, nil
}

func (r *LocationService) GetDistrictByProvinceID(c *gin.Context, provinceId string) (*[]model.Districts, error) {
	districts, err := r.repo.GetDistrictByProvinceID(c, provinceId)
	if err != nil {
		return nil, err
	}
	return districts, nil
}

func (r *LocationService) GetWardByDistrictID(c *gin.Context, districtId string) (*[]model.Wards, error) {
	wards, err := r.repo.GetWardByDistrictID(c, districtId)
	if err != nil {
		return nil, err
	}
	return wards, nil
}
