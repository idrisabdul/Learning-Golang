package options

import (
	"sygap_new_knowledge_management/backend/model"
	masterrepo "sygap_new_knowledge_management/backend/repository/options"

	"github.com/sirupsen/logrus"
)

type RelationTypeService struct {
	repo *masterrepo.RelationTypeRepo
	log  *logrus.Logger
}

func NewRelationTypeService(repo *masterrepo.RelationTypeRepo, log *logrus.Logger) *RelationTypeService {
	return &RelationTypeService{repo, log}
}

func (s *RelationTypeService) GetListRelationType() ([]model.ListRelationType, error) {
	s.log.Println("Execute function GetListRelationType")

	return s.repo.GetListRelationType()
}