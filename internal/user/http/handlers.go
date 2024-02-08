package http

import (
	"go-bookstore/internal/user/service"

	"go-bookstore/internal/user/model"

	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	tokenString, expTime, err := h.service.SignIn(c, &userLogin)
	if err != nil {
		log.Error(err)
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expTime,
	})
}

func (h *UserHandlers) RefreshToken(c *gin.Context) {

}
