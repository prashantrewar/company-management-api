package models

import (
    "github.com/jinzhu/gorm"
)

type Billing struct {
    gorm.Model
    CustomerName string `json:"customer_name" binding:"required"`
    Amount     float64 `json:"amount" binding:"required"`
    Status     string `json:"status" binding:"required"`
}
