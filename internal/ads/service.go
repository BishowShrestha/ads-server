package ads

type AdService struct {
	repo *AdRepository
}

func NewAdService(repo *AdRepository) *AdService {
	return &AdService{repo: repo}
}

func (s *AdService) ListAds() ([]Ad, error) {
	return s.repo.GetAllAds()
}
