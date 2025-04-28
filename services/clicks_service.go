package services

import (
	"ad-server/model"
	"ad-server/repository"
)

type IClicksService interface {
	SaveClick(click model.ClickEvent) error
}

type ClicksService struct {
	repo repository.IClicksRepository
}

func NewClicksService(repo repository.IClicksRepository) IClicksService {
	return &ClicksService{repo: repo}
}

func (s *ClicksService) SaveClick(click model.ClickEvent) error {
	return s.repo.SaveClick(click)
}
