package http

import (
	"go-bookstore/internal/location/service"

	_ "go-bookstore/internal/location/model"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type LocationHandlers struct {
	service service.ILocationService
}

func NewLocationHandlers(service service.ILocationService) *LocationHandlers {
	return &LocationHandlers{service: service}
}

// GetAllProvince godoc
//
//	@Summary	get all provinces in VN
//	@Tags		location
//	@Produce	json
//	@Success	200	{object}	model.Provinces
//	@Router		/location/provinces [get]
func (h *LocationHandlers) GetAllProvince(c *gin.Context) {
	provinces, err := h.service.GetAllProvince(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "request fail",
		})
		return
	}
	c.JSON(http.StatusOK, provinces)
}

// GetDistrictByProvinceID godoc
//
//	@Summary	get districts by province id
//	@Tags		location
//	@Produce	json
//	@Param		provinceId	query		string	true	"Query"
//	@Success	200			{object}	model.Districts
//	@Router		/location/districts/:provinceId [get]
func (h *LocationHandlers) GetDistrictByProvinceID(c *gin.Context) {
	provinceId := c.Query("provinceId")
	districts, err := h.service.GetDistrictByProvinceID(c, provinceId)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "request fail",
		})
		return
	}
	c.JSON(http.StatusOK, districts)
}

// GetWardByDistrictID godoc
//
//	@Summary	get wards by district id
//	@Tags		location
//	@Produce	json
//	@Param		districtId	query		string	true	"Query"
//	@Success	200			{object}	model.Wards
//	@Router		/location/wards/:districtId [get]
func (h *LocationHandlers) GetWardByDistrictID(c *gin.Context) {
	districtId := c.Query("districtId")
	wards, err := h.service.GetWardByDistrictID(c, districtId)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "request fail",
		})
		return
	}
	c.JSON(http.StatusOK, wards)
}
