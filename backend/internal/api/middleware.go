package api

import (
	"chater/internal/domain/auth"
	"chater/internal/logging"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTAuthMiddleware - middleware для аутентификации через JWT
func JWTAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		logging.Logger.Debug("Check JWT")
		// Извлекаем JWT токен из Cookie
		tokenString, err := c.Cookie("token")
		if err != nil {
			logging.Logger.Error(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			logging.Logger.Error(err.Error())
			c.JSON(http.StatusUnauthorized, errorResponse{Error: "Invalid or expired token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*auth.Claims); ok && token.Valid {
			c.Set("user_id", claims.UserID)
		} else {
			logging.Logger.Error("Invalid token claims")
			c.JSON(http.StatusUnauthorized, errorResponse{Error: "Invalid token claims"})
			c.Abort()
			return
		}
		logging.Logger.Debug("JWT is valid. Next...")
		c.Next()
	}
}
