package clicks

type ClickService struct {
	repo *ClickRepository
}

func NewClickService(repo *ClickRepository) *ClickService {
	return &ClickService{repo: repo}
}

func (s *ClickService) SaveClick(click ClickEvent) error {
	return s.repo.SaveClick(click)
}
