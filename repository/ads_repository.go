package repository

import (
	"ad-server/model"
	"gorm.io/gorm"
)

type IAdsRepository interface {
	CreateAds(ad model.Ad) error
	GetAllAds() ([]model.Ad, error)
}

type AdsRepository struct {
	DB *gorm.DB
}

func NewAdsRepository(db *gorm.DB) IAdsRepository {
	return &AdsRepository{DB: db}
}

func (r *AdsRepository) CreateAds(ad model.Ad) error {
	if err := r.DB.Create(&ad).Error; err != nil {
		return err
	}
	return nil
}

func (r *AdsRepository) GetAllAds() ([]model.Ad, error) {
	var ads []model.Ad
	if err := r.DB.Find(&ads).Error; err != nil {
		return nil, err
	}
	return ads, nil
}
