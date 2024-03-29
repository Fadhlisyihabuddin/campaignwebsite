package campaign

import (
	"fmt"
	"strconv"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaign(UserID int) ([]Campaign, error)
	GetCampaignDetail(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
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

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID

	slugString := fmt.Sprintf("%s %s", input.Name, strconv.Itoa(input.User.ID))
	campaign.Slug = slug.Make(slugString)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}
