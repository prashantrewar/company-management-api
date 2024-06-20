package handlers

import (
    "github.com/gin-gonic/gin"
    "company-management-api/models"
    "company-management-api/db"
    "net/http"
)

func CreateCustomer(c *gin.Context) {
    var customer models.Customer
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.DB.Create(&customer).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
        return
    }

    c.JSON(http.StatusOK, customer)
}

func GetCustomers(c *gin.Context) {
    var customers []models.Customer
    if err := db.DB.Find(&customers).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve customers"})
        return
    }

    c.JSON(http.StatusOK, customers)
}
