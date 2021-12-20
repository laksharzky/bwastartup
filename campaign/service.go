package campaign

type Service interface {
	GetUserByID(UserID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUserByID(userID int) ([]Campaign, error) {
	campaign, err := s.repository.FindUserByID(userID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
