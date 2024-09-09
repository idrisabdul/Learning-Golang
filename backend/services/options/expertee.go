package options

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/repository/options"

	"github.com/sirupsen/logrus"
)

type ExperteeService struct {
	repo *options.ExperteeRepos
	log  *logrus.Logger
}

func NewExperteeService(repo *options.ExperteeRepos, log *logrus.Logger) *ExperteeService {
	return &ExperteeService{repo, log}
}

func (s *ExperteeService) GetExperteeGroup(idCompany string) ([]entities.Organization, error) {
	return s.repo.GetExperteeGroup(idCompany)
}

func (s *ExperteeService) GetExpertees(idOrganization string) ([]entities.ListExpertees, error) {
	return s.repo.GetExpertees(idOrganization)
}
