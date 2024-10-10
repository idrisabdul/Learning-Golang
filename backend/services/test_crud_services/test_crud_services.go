package testcrudservices

import (
	"sygap_new_knowledge_management/backend/model"
	testcrud "sygap_new_knowledge_management/backend/repository/test_crud"

	"github.com/sirupsen/logrus"
)

type TestCrudSvc struct {
	repo *testcrud.TestCrudRepository
	log  *logrus.Logger
}

func NewTestCrudSvc(repo *testcrud.TestCrudRepository, log *logrus.Logger) *TestCrudSvc {
	return &TestCrudSvc{repo, log}
}

func (s *TestCrudSvc) GetListTest() ([]model.ListCrudModel, error) {
	return s.repo.GetTestCrud()
}

func (s *TestCrudSvc) GetDetailTestSvc(idKnowledgeContent int) (model.GetDetailCrudModel, error) {
	detailTestSvc, errDetailTestSvc := s.repo.GetDetailCrud(idKnowledgeContent)
	if errDetailTestSvc != nil {
		return model.GetDetailCrudModel{}, errDetailTestSvc
	}

	return detailTestSvc, nil
}

func (s *TestCrudSvc) InsertCrudTestSvc(data model.AddKnowledgeContent) (model.AddKnowledgeContent, error) {
	return s.repo.InsertCrudTest(data)
}

func (s *TestCrudSvc) UpdateCrudTestSvc(data model.UpdateKnowledgeContent) (model.UpdateKnowledgeContent, error) {
	return s.repo.UpdateCrudTest(data)
}
