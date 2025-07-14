package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"

	"Backend-POS/utils/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header id requird"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			return
		}

		token := parts[1]
		claims, err := jwt.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		type staff struct {
			ID           int    `json:"id"`
			Username     string `json:"username"`
			PasswordHash string `json:"password_hash"`
		}

		tt, _ := json.Marshal(claims["sub"])
		var stf *staff
		_ = json.Unmarshal(tt, &stf)
		ctx.Set("staff_id", stf.ID)

		ctx.Next()
	}
}
