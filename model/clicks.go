package model

import (
	"gorm.io/gorm"
	"time"
)

type ClickEvent struct {
	gorm.Model
	AdID         uint   `gorm:"index" json:"ad_id"`
	IPAddress    string `json:"ip_address"`
	PlaybackTime int    `json:"playback_time"`
}

type HourAnalytics struct {
	Hour  time.Time `json:"hour"`
	Count int       `json:"count"`
}
