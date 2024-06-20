package handlers

import (
    "github.com/gin-gonic/gin"
    "company-management-api/models"
    "company-management-api/db"
    "net/http"
)

func CreateBilling(c *gin.Context) {
    var billing models.Billing
    if err := c.ShouldBindJSON(&billing); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.DB.Create(&billing).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create billing"})
        return
    }

    c.JSON(http.StatusOK, billing)
}

func GetBillings(c *gin.Context) {
    var billings []models.Billing
    db.DB.Find(&billings)
    c.JSON(http.StatusOK, billings)
}
