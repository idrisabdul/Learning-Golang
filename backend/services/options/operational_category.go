package options

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/repository/options"

	"github.com/sirupsen/logrus"
)

type OperationalCategoryService struct {
	repo *options.OperationalCategoryRepos
	log  *logrus.Logger
}

func NewOperationalCategorysService(repo *options.OperationalCategoryRepos, log *logrus.Logger) *OperationalCategoryService {
	return &OperationalCategoryService{repo: repo, log: log}
}

func (s *OperationalCategoryService) GetOpCat1() ([]entities.OperationCategory, error) {
	return s.repo.GetOpCat1()
}

func (s *OperationalCategoryService) GetOpCat2(idParent string) ([]entities.OperationCategory, error) {
	return s.repo.GetOpCat2(idParent)
}
