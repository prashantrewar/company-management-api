package handlers

import (
    "net/http"
    "company-management-api/db"
    "company-management-api/models"
	"company-management-api/utils"

    "github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
        return
    }
    user.Password = hashedPassword

    if err := db.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
    var users []models.User
    if err := db.DB.Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, users)
}
