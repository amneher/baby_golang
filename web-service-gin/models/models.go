package models

import (
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	Title  string  `json:"title"`
	Artist string  `json:"artist" gorm:"index"`
	Price  float64 `json:"price"`
}
