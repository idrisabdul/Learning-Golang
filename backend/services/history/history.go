package history

import (
	"github.com/sirupsen/logrus"
	"sygap_new_knowledge_management/backend/entities"
	repository "sygap_new_knowledge_management/backend/repository/history"
	"sygap_new_knowledge_management/backend/utils"
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

func (s *HistoryService) GetHistoryListApprove(encodedIDKM string) ([]entities.HistoryKnowledgeList, error) {
	decodedIDKM, _ := utils.GenerateDecoded(encodedIDKM)
	return s.repo.GetHistoryApprove(decodedIDKM)
}

func (s *HistoryService) GetHistoryListApproveReject(encodedIDKM string) ([]entities.HistoryKnowledgeList, error) {
	decodedIDKM, _ := utils.GenerateDecoded(encodedIDKM)
	return s.repo.GetHistoryApproveReject(decodedIDKM)
}

func (s *HistoryService) GetHistoryListRequested(encodedIDKM string) ([]entities.HistoryNotif, error) {
	decodedIDKM, _ := utils.GenerateDecoded(encodedIDKM)
	return s.repo.GetHistoryRequested(decodedIDKM)
}
