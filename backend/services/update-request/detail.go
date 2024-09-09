package updaterequest

import (
	"sygap_new_knowledge_management/backend/entities"
	updaterequest "sygap_new_knowledge_management/backend/repository/update-request"

	"github.com/sirupsen/logrus"
)

type DetailService struct {
	repo *updaterequest.DetailRepos
	log  *logrus.Logger
}

func NewDetailService(repo *updaterequest.DetailRepos, log *logrus.Logger) *DetailService {
	return &DetailService{repo, log}
}

func (s *DetailService) GetListUpdateRequest(idKM string) ([]entities.ListUpdateRequest, error) {
	return s.repo.GetListUpdateRequest(idKM)
}

func (s *DetailService) GetDetailUpdateRequest(idUpdateRequest string) (entities.DetailUpdateRequest, error) {
	detailUpdateRequest, errDetailUpdateRequest := s.repo.GetDetailUpdateRequest(idUpdateRequest)
	if errDetailUpdateRequest != nil {
		return entities.DetailUpdateRequest{}, errDetailUpdateRequest
	}

	attachmentUpdateRequest, errAttachmentUpdateRequest := s.repo.GetUpdateRequestAttachment(idUpdateRequest)
	if errAttachmentUpdateRequest != nil {
		return entities.DetailUpdateRequest{}, nil
	}

	detailUpdateRequest.File = attachmentUpdateRequest

	return detailUpdateRequest, nil
}
