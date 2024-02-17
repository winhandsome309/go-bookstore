package middleware

import (
	"fmt"
	"go-bookstore/internal/user/model"
	"go-bookstore/pkg/dbs"

	"time"

	"go-bookstore/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.GetConfig()
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			// c.AbortWithStatus(http.StatusUnauthorized)
			c.Next()
			return
		}
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(cfg.AuthSecret), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Check the expiry time
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				// c.AbortWithStatus(http.StatusUnauthorized)
				c.Next()
				return
			}
			// Find the user with token
			var user model.User
			err = dbs.GetDB().Where("email = ?", claims["payload"].(map[string]interface{})["email"]).First(&user).Error
			if err != nil {
				// c.AbortWithStatus(http.StatusUnauthorized)
				c.Next()
				return
			}
			c.Set("user", user)
			// Refresh token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"payload": claims["payload"].(map[string]interface{}),
				"exp":     time.Now().Add(20 * time.Minute).Unix(),
			})
			tokenString, err := token.SignedString([]byte(cfg.AuthSecret))
			if err != nil {
				c.Next()
				return
			}
			maxAge := time.Now().Unix() + int64(60)
			c.SetCookie("Authorization", tokenString, int(maxAge), "/", "", true, false)
		} else {
			// c.AbortWithStatus(http.StatusUnauthorized)
			c.Next()
			return
		}
		c.Next()
	}
}
