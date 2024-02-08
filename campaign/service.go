package campaign

type Service interface {
	GetCampaign(UserID int) ([]Campaign, error)
	GetCampaignDetail(input GetCampaignDetailInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaign(UserID int) ([]Campaign, error) {
	if UserID != 0 {
		campaign, err := s.repository.FindByUserID(UserID)

		if err != nil {
			return campaign, err
		}

		return campaign, nil
	}

	campaign, err := s.repository.FindAll()

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) GetCampaignDetail(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.Id)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
