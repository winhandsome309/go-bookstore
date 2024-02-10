package middleware

import (
	"fmt"
	"go-bookstore/internal/user/model"
	"go-bookstore/pkg/config"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
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
			return []byte("30092002"), nil
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
			err = config.GetDB().Where("email = ?", claims["payload"].(map[string]interface{})["email"]).First(&user).Error
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
			tokenString, err := token.SignedString([]byte("30092002"))
			if err != nil {
				c.Next()
				return
			}
			maxAge := time.Now().Unix() + int64(60)
			c.SetCookie("Authorization", tokenString, int(maxAge), "/", "", false, false)
		} else {
			// c.AbortWithStatus(http.StatusUnauthorized)
			c.Next()
			return
		}
		c.Next()
	}
}

// func RefreshToken() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString, ok := c.Cookie("Authorization")
// 		if ok != nil {
// 			c.Next()
// 			return
// 		}
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			// Don't forget to validate the alg is what you expect:
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 			}
// 			return "30092002", nil
// 		})
// 		if err != nil {
// 			c.Next()
// 			return
// 		}
// 		if claims, ok := token.Claims.(jwt.MapClaims); ok {
// 			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 				"payload": claims["payload"],
// 				"exp":     time.Now().Add(20 * time.Minute).Unix(),
// 			})
// 			tokenString, err := token.SignedString("30092002")
// 			if err != nil {
// 				c.Next()
// 				return
// 			} else {
// 				maxAge := time.Now().Unix() + int64(60)
// 				c.SetCookie("Authorization", tokenString, int(maxAge), "/", "", false, false)
// 				c.Next()
// 			}
// 		} else {
// 			c.Next()
// 			return
// 		}
// 	}
// }
