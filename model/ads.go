package model

import "gorm.io/gorm"

type Ad struct {
	gorm.Model
	ImageURL  string `gorm:"column:image_url" json:"image_url"`
	TargetURL string `gorm:"column:target_url" json:"target_url"`
}
