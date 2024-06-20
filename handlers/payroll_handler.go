package handlers

import (
    "github.com/gin-gonic/gin"
    "company-management-api/models"
    "company-management-api/db"
    "net/http"
)

func CreatePayroll(c *gin.Context) {
    var payroll models.Payroll
    if err := c.ShouldBindJSON(&payroll); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.DB.Create(&payroll).Error; err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to create payroll"})
        return
    }
    c.JSON(http.StatusOK, payroll)
}

func GetPayrolls(c *gin.Context) {
    var payrolls []models.Payroll
    db.DB.Find(&payrolls)
    c.JSON(http.StatusOK, payrolls)
}
