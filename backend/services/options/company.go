package options

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/repository/options"

	"github.com/sirupsen/logrus"
)

type CompanyService struct {
	repo *options.CompanyRepos
	log *logrus.Logger
}

func NewCompanyService(repo *options.CompanyRepos, log *logrus.Logger) *CompanyService {
	return &CompanyService{repo, log}
}

func (s *CompanyService) GetCompanies() ([]entities.Company, error) {
	return s.repo.GetCompanies()
}