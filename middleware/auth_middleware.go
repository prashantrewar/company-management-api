package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "company-management-api/utils"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.Request.URL.Path == "/login" {
            c.Next()
            return
        }

        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
            c.Abort()
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token format invalid"})
            c.Abort()
            return
        }

        token := parts[1]
        if _, err := utils.ValidateJWT(token); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
            c.Abort()
            return
        }

        c.Next()
    }
}
