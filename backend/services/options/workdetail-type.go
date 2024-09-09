package options

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/repository/options"

	"github.com/sirupsen/logrus"
)

type WorkDetailTypeService struct {
	repo *options.WorkDetailTypeRepos
	log  *logrus.Logger
}

func NewworkDetailTypeService(repo *options.WorkDetailTypeRepos, log *logrus.Logger) *WorkDetailTypeService {
	return &WorkDetailTypeService{repo, log}
}

func (s *WorkDetailTypeService) GetWorkDetailType() ([]entities.ListWorkDetailType, error) {

	var response []entities.ListWorkDetailType
	parentType, errGetParentType := s.repo.GetWorkDetailType(0)
	if errGetParentType != nil {
		return nil, errGetParentType
	}

	for _, v := range parentType {
		childList, errGetChildList := s.repo.GetWorkDetailType(v.ID)
		if errGetChildList != nil || len(childList) == 0 {
			continue
		}

		response = append(response, entities.ListWorkDetailType{
			Label:   v.Type,
			Options: childList,
		})
	}

	return response, nil
}
