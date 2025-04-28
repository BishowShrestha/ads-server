package services

import (
	"ad-server/model"
	"ad-server/repository"
)

type IAnalyticsService interface {
	GetAdClickCounts() (map[uint]int, error)
	GetHourlyAnalytics() ([]model.HourAnalytics, error)
}

type AnalyticsService struct {
	clickRepo repository.IClicksRepository
}

func NewAnalyticsService(clickRepo repository.IClicksRepository) IAnalyticsService {
	return &AnalyticsService{clickRepo: clickRepo}
}
func (s *AnalyticsService) GetAdClickCounts() (map[uint]int, error) {
	return s.clickRepo.GetClickCounts()
}

func (s *AnalyticsService) GetHourlyAnalytics() ([]model.HourAnalytics, error) {
	return s.clickRepo.GetHourlyAnalytics()
}
