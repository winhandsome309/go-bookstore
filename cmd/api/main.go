package main

import (
	productHttp "go-bookstore/internal/product/http"
	"go-bookstore/internal/product/model"
	"go-bookstore/pkg/config"

	"github.com/gin-contrib/cors"
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

	// Use CORS
	r.Use(cors.Default())

	// Init servers
	productHttp.Routes(r, db)

	// Port to run
	r.Run(":8080")
}
