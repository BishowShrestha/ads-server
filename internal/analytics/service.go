package analytics

import "ad-server/internal/clicks"

type AnalyticsService struct {
	clickRepo *clicks.ClickRepository
}

func NewAnalyticsService(clickRepo *clicks.ClickRepository) *AnalyticsService {
	return &AnalyticsService{clickRepo: clickRepo}
}

func (s *AnalyticsService) GetClickCounts() (map[uint]int, error) {
	return s.clickRepo.GetClickCounts()
}
