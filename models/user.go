package models

import "time"

type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Username  string    `json:"username" binding:"required"`
    Password  string    `json:"password" binding:"required"`
    Role      string    `json:"role" binding:"required"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}
