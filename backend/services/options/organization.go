package options

import (
	"sygap_new_knowledge_management/backend/model"
	options "sygap_new_knowledge_management/backend/repository/options"

	"github.com/sirupsen/logrus"
)

type OrganizationService struct {
	repo *options.OrganizationRepo
	log  *logrus.Logger
}

func NewOrganizationService(repo *options.OrganizationRepo, log *logrus.Logger) *OrganizationService {
	return &OrganizationService{repo, log}
}

func (s *OrganizationService) GetActiveOrganization(search string, idCompany string, isAll string) ([]model.ListOrganization, error) {
	s.log.Println("Execute function GetActiveOrganization")

	var result []model.ListOrganization

	if isAll == "true" {
		data, _ := s.repo.GetAllDetailOrganization(search)
		result = data
	} else {
		data, _ := s.repo.GetDetailOrganization(search, idCompany)
		result = data
	}
	return result, nil
}
