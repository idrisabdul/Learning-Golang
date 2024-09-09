package document

import (
	"sygap_new_knowledge_management/backend/repository/km/document"

	"github.com/sirupsen/logrus"
)

type DeleteService struct {
	repo *document.DeleteDocument
	log  *logrus.Logger
}

func NewDeleteService(repo *document.DeleteDocument, log *logrus.Logger) *DeleteService {
	return &DeleteService{repo, log}
}

func (s *DeleteService) DeleteDocumentKM(idFile, author string) error {
	return s.repo.DeleteDocumentKM(idFile, author)
}
