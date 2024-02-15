package http

import (
	"go-bookstore/internal/location/service"

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
