package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(inp GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(id GetCampaignDetailInput, data CreateCampaignInput) (Campaign, error)
}

type service struct{ repository Repository }

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userId int) ([]Campaign, error) {
	if userId != 0 {
		campaigns, err := s.repository.FindByUserId(userId)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) GetCampaignById(inp GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindById(inp.Id)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {

	stringConcat := fmt.Sprintf("%s %d", input.Name, input.User.Id)

	campaign := Campaign{
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Perks:            input.Perks,
		GoalAmount:       input.GoalAmount,
		UserId:           input.User.Id,
		Slug:             slug.Make(stringConcat),
	}

	newCampaign, err := s.repository.Save(campaign)

	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *service) UpdateCampaign(id GetCampaignDetailInput, data CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindById(id.Id)

	if err != nil {
		return campaign, err
	}

	if campaign.UserId != data.User.Id {
		return campaign, errors.New("not owner of the campaign")
	}

	campaign.Name = data.Name
	campaign.ShortDescription = data.ShortDescription
	campaign.Description = data.Description
	campaign.Perks = data.Perks
	campaign.GoalAmount = data.GoalAmount

	update, err := s.repository.Update(campaign)

	if err != nil {
		return update, err
	}

	return update, nil
}
