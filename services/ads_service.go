package services

import (
	"ad-server/model"
	"ad-server/repository"
)

type IAdsService interface {
	CreateAds(ad model.Ad) error
	GetAllAds() ([]model.Ad, error)
}

type AdsService struct {
	adsRepo repository.IAdsRepository
}

func NewAdsService(repo repository.IAdsRepository) IAdsService {
	return &AdsService{adsRepo: repo}
}

func (s *AdsService) CreateAds(ad model.Ad) error {
	return s.adsRepo.CreateAds(ad)
}

func (s *AdsService) GetAllAds() ([]model.Ad, error) {
	return s.adsRepo.GetAllAds()
}
