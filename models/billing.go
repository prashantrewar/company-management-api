package models

import (
    "github.com/jinzhu/gorm"
)

type Billing struct {
    gorm.Model
    CustomerID uint   `json:"customer_id" binding:"required"`
    Amount     float64 `json:"amount" binding:"required"`
    Status     string `json:"status" binding:"required"`
}
