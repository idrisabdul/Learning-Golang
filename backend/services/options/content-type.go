package options

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/repository/options"

	"github.com/sirupsen/logrus"
)

type ContentTypeService struct {
	repo *options.ContentTypeRepos
	log *logrus.Logger
}

func NewContentTypeService(repo *options.ContentTypeRepos, log *logrus.Logger) *ContentTypeService {
	return &ContentTypeService{repo, log}
}

func (s *ContentTypeService) GetContentType() ([]entities.ContentType, error) {
	return s.repo.GetContentType()
}