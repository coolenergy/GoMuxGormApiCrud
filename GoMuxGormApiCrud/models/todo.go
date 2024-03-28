package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"default:null" json:"description"`
	Completed   bool   `gorm:"not null;default:false" json:"completed"`
}
