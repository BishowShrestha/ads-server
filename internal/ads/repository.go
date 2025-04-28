package ads

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AdRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewAdRepository(db *gorm.DB, log *zap.Logger) *AdRepository {
	return &AdRepository{db: db, log: log}
}

func (r *AdRepository) GetAllAds() ([]Ad, error) {
	var ads []Ad
	if err := r.db.Find(&ads).Error; err != nil {
		return nil, err
	}
	return ads, nil
}
