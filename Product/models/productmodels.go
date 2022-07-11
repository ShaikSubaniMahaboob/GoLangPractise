package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	//ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Number      string `json:"number"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
