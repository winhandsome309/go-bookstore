package main

import (
	_ "go-bookstore/docs"
	locationHttp "go-bookstore/internal/location/http"
	orderHttp "go-bookstore/internal/order/http"
	paymentHttp "go-bookstore/internal/payment/http"
	productHttp "go-bookstore/internal/product/http"
	"go-bookstore/internal/product/model"
	shippingHttp "go-bookstore/internal/shipping/http"
	userHttp "go-bookstore/internal/user/http"
	"go-bookstore/pkg/config"
	"go-bookstore/pkg/dbs"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	// c.Header("Access-Control-Allow-Origin", "https://bookstore-fe-v8ch.onrender.com")
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
}

//	@title			Go Bookstore Application
//	@description	This is a bookstore e-commerce application
//	@version		1.0
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	WinHandsome
//	@contact.url	https://web.facebook.com/winhandsomee/
//	@contact.email	xuanthangnguyen2002@gmail.com

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/

func main() {
	cfg := config.LoadConfig()

	// Init database
	dbs.Connect(cfg.DatabaseURI)
	db := dbs.GetDB()
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
	paymentHttp.Routes(r, db)

	// Init swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Port to run
	// r.Run(":8080")
	r.Run(":" + cfg.HttpPort)
}
