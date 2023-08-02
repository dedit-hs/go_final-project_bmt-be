package models

import (
	"time"
)

type Admin struct {
	Id        int       `json:"id" form:"id" gorm:"primaryKey"`
	Username  string    `json:"username" form:"username" gorm:"unique"`
	Password  string    `json:"password" form:"password"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type AdminResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type AdminLoginResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
