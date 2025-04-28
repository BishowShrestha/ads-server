package ads

type Ad struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ImageURL  string `gorm:"column:image_url" json:"image_url"`
	TargetURL string `gorm:"column:target_url" json:"target_url"`
}
