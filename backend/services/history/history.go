package history

import (
	"github.com/sirupsen/logrus"
	"sygap_new_knowledge_management/backend/entities"
	repository "sygap_new_knowledge_management/backend/repository/history"
)

type HistoryService struct {
	repo *repository.HistoryRepository
	log  *logrus.Logger
}

func NewHistoryService(repo *repository.HistoryRepository, logger *logrus.Logger) *HistoryService {
	return &HistoryService{
		repo: repo,
		log:  logger,
	}
}

func (s *HistoryService) GetHistoryListApprove(idKM string) ([]entities.HistoryKnowledgeList, error) {
	return s.repo.GetHistoryApprove(idKM)
}

func (s *HistoryService) GetHistoryListApproveReject(idKM string) ([]entities.HistoryKnowledgeList, error) {
	return s.repo.GetHistoryApproveReject(idKM)
}

func (s *HistoryService) GetHistoryListRequested(idKM string) ([]entities.HistoryNotif, error) {
	return s.repo.GetHistoryRequested(idKM)
}
