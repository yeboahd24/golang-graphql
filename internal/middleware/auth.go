package middleware

import (
	"auth-system/internal/auth"
	"context"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService *auth.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			c.Next()
			return
		}

		token, err := authService.ValidateToken(bearerToken[1])
		if err != nil {
			c.Next()
			return
		}

		// Add token claims to context
		ctx := context.WithValue(c.Request.Context(), "token", token.Claims)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
