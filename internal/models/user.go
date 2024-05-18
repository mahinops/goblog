package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email" gorm:"unique"`
	Password string `form:"password" json:"password" binding:"required"`
}
