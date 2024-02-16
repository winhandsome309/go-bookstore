package main

import (
	locationHttp "go-bookstore/internal/location/http"
	orderHttp "go-bookstore/internal/order/http"
	productHttp "go-bookstore/internal/product/http"
	"go-bookstore/internal/product/model"
	shippingHttp "go-bookstore/internal/shipping/http"
	userHttp "go-bookstore/internal/user/http"
	"go-bookstore/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Cross-Origin-Resource-Policy", "cross-origin")
	c.Header("Access-Control-Expose-Headers", "Authorization")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}

	// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")
	// if c.Request.Method == "OPTIONS" {
	// 	c.AbortWithStatus(204)
	// 	return
	// }
	// c.Next()
}

func main() {
	// Init database
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&model.Product{})

	// Init gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Use CORS
	// r.Use(cors.Default())

	r.Use(corsMiddleware)

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://127.0.0.1:3000"},
	// 	AllowMethods:     []string{"GET", "PUT", "PATCH", "OPTIONS"},
	// 	AllowHeaders:     []string{"Content-Type, Authorization", "Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge: 12 * time.Hour,
	// }))

	// Init servers
	productHttp.Routes(r, db)
	userHttp.Routes(r, db)
	orderHttp.Routes(r, db)
	locationHttp.Routes(r, db)
	shippingHttp.Routes(r, db)

	// Port to run
	r.Run(":8080")
}
