package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-bookstore/internal/user/repository"
	"go-bookstore/internal/user/service"
)

func Routes(r *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandlers := NewUserHandlers(userService)

	userRoutes := r.Group("/auth")
	{
		userRoutes.POST("/register", userHandlers.Register)
		userRoutes.POST("/signin", userHandlers.SignIn)
		// userRoutes.POST("/refresh", userHandlers.RefreshToken)
	}
	// r.Handle("POST", "auth/signin", userHandlers.SignIn)
}
