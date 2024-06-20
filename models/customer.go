package models

import (
    "github.com/jinzhu/gorm"
)

type Customer struct {
    gorm.Model
    Name    string `json:"name" binding:"required"`
    Address string `json:"address" binding:"required"`
    Email   string `json:"email" binding:"required"`
}
