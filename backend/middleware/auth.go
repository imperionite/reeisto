package middleware

import (
	"errors"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

<<<<<<< HEAD
	"github.com/imperionite/reeisto/backend/config"
	"github.com/imperionite/reeisto/backend/utils"
=======
	"github.com/imperionite/reetis/backend/config"
	"github.com/imperionite/reetis/backend/utils"
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.Error(c, http.StatusUnauthorized, "Missing token")
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &utils.JWTClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
<<<<<<< HEAD
			// ✅ FIX: Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}
=======
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
			return []byte(config.App.JWTSecret), nil
		})

		if err != nil {
<<<<<<< HEAD
=======
			// 🔥 Explicit expiration handling
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
			if errors.Is(err, jwt.ErrTokenExpired) {
				utils.Error(c, http.StatusUnauthorized, "Token expired")
				c.Abort()
				return
			}

			utils.Error(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		if !token.Valid {
			utils.Error(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

<<<<<<< HEAD
=======
		// 🔥 EXTRA SAFETY (manual expiration check)
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			utils.Error(c, http.StatusUnauthorized, "Token expired")
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")

		if slices.Contains(roles, role) {
			c.Next()
			return
		}

		utils.Error(c, http.StatusForbidden, "Forbidden")
		c.Abort()
	}
}