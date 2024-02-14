package http

import (
	"go-bookstore/internal/user/service"
	"time"

	"go-bookstore/internal/user/model"

	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserHandlers struct {
	service service.IUserService
}

func NewUserHandlers(service service.IUserService) *UserHandlers {
	return &UserHandlers{service: service}
}

func (h *UserHandlers) Register(c *gin.Context) {
	var userReq model.UserReq
	if err := c.ShouldBind(&userReq); c.Request.Body == nil || err != nil {
		log.Error("Invalid parameters", err)
		return
	}
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(userReq.Password), 14)
	userReq.Password = string(passwordHash)
	err := h.service.Register(c, &userReq)
	if err != nil {
		log.Error("Register Fail", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Register successfully",
	})
}

func (h *UserHandlers) SignIn(c *gin.Context) {
	var userLogin model.UserLogin
	if err := c.ShouldBind(&userLogin); c.Request.Body == nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters",
		})
		return
	}
	tokenString, _, err := h.service.SignIn(c, &userLogin)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}
	maxAge := int(time.Now().Unix() + int64(60))
	c.SetCookie("Authorization", tokenString, maxAge, "/", "", false, false)

	// http.SetCookie(c.Writer, &http.Cookie{
	// 	Name:     "Authorization",
	// 	Value:    tokenString,
	// 	Expires:  time.Now().Add(time.Hour * 24),
	// 	HttpOnly: true,
	// 	SameSite: http.SameSiteDefaultMode,
	// 	Secure:   false,
	// 	Path:     "/",
	// })

	c.JSON(http.StatusAccepted, gin.H{
		"message": "User signin successfully",
	})
}

func (h *UserHandlers) SignOut(c *gin.Context) {
	// tokenString, err := h.service.SignOut(c)
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "log out fail",
		})
		return
	}
	c.SetCookie("Authorization", tokenString, -1, "/", "", false, false)
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Sign out successfully",
	})
}
