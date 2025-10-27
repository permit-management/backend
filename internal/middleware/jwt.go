package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/permit-management/backend/pkg/setting"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid Authorization header"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(setting.Conf.App.JWTSecret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if exp, ok := claims["exp"].(float64); ok {
				if int64(exp) < time.Now().Unix() {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
					return
				}
			}

			// Save user info to context for next handlers
			c.Set("user_id", claims["user_id"])
			c.Set("email", claims["email"])
			c.Set("role_id", claims["role_id"])

			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
		}
	}
}
