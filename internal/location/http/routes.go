package http

import (
	"go-bookstore/internal/location/repository"
	"go-bookstore/internal/location/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(c *gin.Engine, db *gorm.DB) {
	locationRepo := repository.NewLocationRepository(db)
	locationService := service.NewLocationService(locationRepo)
	locationHandlers := NewLocationHandlers(locationService)

	locationRoutes := c.Group("/location")
	{
		locationRoutes.GET("/provinces", locationHandlers.GetAllProvince)
		locationRoutes.GET("/districts", locationHandlers.GetDistrictByProvinceID)
		locationRoutes.GET("/wards", locationHandlers.GetWardByDistrictID)
	}
}
