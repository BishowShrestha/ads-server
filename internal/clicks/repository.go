package clicks

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ClickRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewClickRepository(db *gorm.DB, log *zap.Logger) *ClickRepository {
	return &ClickRepository{db: db, log: log}
}
func (r *ClickRepository) SaveClick(click ClickEvent) error {
	return r.db.Create(&click).Error
}

func (r *ClickRepository) GetClickCounts() (map[uint]int, error) {
	type Result struct {
		AdID  uint
		Count int
	}
	var results []Result

	err := r.db.Model(&ClickEvent{}).
		Select("ad_id, COUNT(*) as count").
		Group("ad_id").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	counts := make(map[uint]int)
	for _, res := range results {
		counts[res.AdID] = res.Count
	}
	return counts, nil
}
