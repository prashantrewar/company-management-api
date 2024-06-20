package handlers

import (
	"net/http"
	"company-management-api/db"
	"company-management-api/models"
	"company-management-api/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
    var loginRequest models.LoginRequest
    if err := c.ShouldBindJSON(&loginRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := db.DB.Where("username = ?", loginRequest.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    if !utils.CheckPasswordHash(loginRequest.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    token, err := utils.GenerateJWT(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token": token,
        "role": user.Role,
    })
}
