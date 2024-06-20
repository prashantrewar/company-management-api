package models

import (
    "github.com/jinzhu/gorm"
)

type Payroll struct {
    gorm.Model
    EmployeeName string  `json:"employee_name" binding:"required"`
    Amount       float64 `json:"amount" binding:"required"`
    Status       string  `json:"status" binding:"required"`
}
