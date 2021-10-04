package campaign

type Service interface{
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(inp GetCampaignDetailInput) (Campaign,error)
}

type service struct {repository Repository}

func NewService(repository Repository) *service  {
	return &service{repository}
}

func (s *service)GetCampaigns(userId int) ([]Campaign, error)  {
	if userId != 0 {
		campaigns, err := s.repository.FindByUserId(userId)
		if err != nil {
			return campaigns, err
		}

		return campaigns,nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) GetCampaignById(inp GetCampaignDetailInput) (Campaign, error)  {
	campaign, err := s.repository.FindById(inp.Id)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}