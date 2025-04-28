package clicks

type ClickEvent struct {
	ID           uint   `gorm:"primaryKey"`
	AdID         uint   `gorm:"index" json:"ad_id"`
	IPAddress    string `json:"ip_address"`
	PlaybackTime int    `json:"playback_time"`
	CreatedAt    int64  `gorm:"autoCreateTime:milli"`
}
