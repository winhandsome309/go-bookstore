package main

import (
	productHttp "go-bookstore/internal/product/http"
	"go-bookstore/internal/product/model"
	"go-bookstore/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init database
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&model.Product{})

	// Init gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Create server
	productHttp.Routes(r, db)

	// Init routes
	// routes.InitRoutes(r)
	r.Run(":8080")
}
