package options

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/repository/options"

	"github.com/sirupsen/logrus"
)

type UpdateRequestService struct {
	repo *options.UpdateRequestRepos
	log  *logrus.Logger
}

func NewUpdateRequestService(repo *options.UpdateRequestRepos, log *logrus.Logger) *UpdateRequestService {
	return &UpdateRequestService{repo, log}
}

func (s *UpdateRequestService) GetListUpdateRequestStatus() ([]entities.ListupdateRequestStatus, error) {
	return s.repo.GetListUpdateRequestStatus()
}
