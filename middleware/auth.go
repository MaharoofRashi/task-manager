package middleware

import (
	"errors"
	"github.com/MaharoofRashi/task-manager/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTMiddleware(jwtUtil *utils.JWTUtil) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.RespondError(c, http.StatusUnauthorized, errors.New("authorization header is missing"))
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			utils.RespondError(c, http.StatusUnauthorized, errors.New("bearer token missing"))
			c.Abort()
			return
		}

		userID, err := jwtUtil.ValidateToken(tokenString)
		if err != nil {
			utils.RespondError(c, http.StatusUnauthorized, errors.New("invalid or expired token"))
			c.Abort()
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}
