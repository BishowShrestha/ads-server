package repository

import (
	"ad-server/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IClicksRepository interface {
	SaveClick(click model.ClickEvent) error
	GetClickCounts() (map[uint]int, error)
	GetHourlyAnalytics() ([]model.HourAnalytics, error)
}

type ClicksRepository struct {
	DB  *gorm.DB
	log *zap.Logger
}

func NewClicksRepository(db *gorm.DB, log *zap.Logger) IClicksRepository {
	return &ClicksRepository{DB: db, log: log}
}

func (s *ClicksRepository) SaveClick(click model.ClickEvent) error {
	return s.DB.Create(&click).Error
}

func (s *ClicksRepository) GetClickCounts() (map[uint]int, error) {
	type Result struct {
		AdID  uint
		Count int
	}
	var results []Result

	err := s.DB.Model(&model.ClickEvent{}).
		Select("ad_id, COUNT(*) as count").
		Group("ad_id").
		Scan(&results).Error

	if err != nil {
		s.log.Error("Error fetching click counts", zap.Error(err))
		return nil, err
	}

	counts := make(map[uint]int)
	for _, res := range results {
		counts[res.AdID] = res.Count
	}
	return counts, nil
}

func (s *ClicksRepository) GetHourlyAnalytics() ([]model.HourAnalytics, error) {
	var results []model.HourAnalytics
	err := s.DB.Raw(`
        SELECT date_trunc('hour', created_at) AS hour, COUNT(*) 
        FROM click_events 
        GROUP BY hour
        ORDER BY hour
    `).Scan(&results).Error
	if err != nil {
		s.log.Error("Error fetching hourly analytics", zap.Error(err))
		return nil, err
	}
	return results, nil
}
