package history

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"strconv"
	"sygap_new_knowledge_management/backend/entities"
	repository "sygap_new_knowledge_management/backend/repository/history"
	"sygap_new_knowledge_management/backend/repository/search"
	"sygap_new_knowledge_management/backend/utils"
)

type HistoryService struct {
	repo       *repository.HistoryRepository
	log        *logrus.Logger
	searchRepo *search.SearchRepos
}

func NewHistoryService(repo *repository.HistoryRepository, searchRepo *search.SearchRepos, logger *logrus.Logger) *HistoryService {
	return &HistoryService{
		repo:       repo,
		log:        logger,
		searchRepo: searchRepo,
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

func (s *HistoryService) GetContentDetail(c *fiber.Ctx, userId string) (interface{}, error) {
	contentId, _ := utils.GenerateDecoded(c.Params("content_id"))
	historyId, _ := utils.GenerateDecoded(c.Params("history_id"))
	data, err := s.searchRepo.GetSearchDetail(contentId, userId)
	s.GetUpsertVisitor(contentId, userId)

	idKMHistory, _ := strconv.Atoi(historyId)
	idKMContent, _ := strconv.Atoi(contentId)
	var kmHistory entities.HistoryKnowledge
	if errKmDetailHistory := s.repo.GetHistoryKnowledgeByIdAndKMContentId(&kmHistory, idKMHistory, idKMContent); errKmDetailHistory != nil {
		s.log.Errorln(errKmDetailHistory.Error())
		return nil, errKmDetailHistory
	}

	dataMap, _ := data.(entities.SearchDetailResponse)
	dataMapContent, _ := dataMap.Content.(entities.SearchDetailChildResponse)

	if kmHistory.Type == "workaround" {
		dataMapContent.Workaround = kmHistory.Value
	} else if kmHistory.Type == "fix-solution" {
		dataMapContent.FixSolution = kmHistory.Value
	} else if kmHistory.Type == "reference" {
		dataMapContent.Reference = kmHistory.Value
	}

	dataMap.Content = dataMapContent

	result := entities.SearchPreviewDetailResponse{
		ID:            dataMap.ID,
		Type:          dataMap.Type,
		Title:         dataMap.Title,
		LastVisitor:   dataMap.LastVisitor,
		Keywords:      dataMap.Keywords,
		Content:       dataMapContent,
		Sidebar:       dataMap.Sidebar,
		Attachment:    dataMap.Attachment,
		HistoryStatus: kmHistory.Status,
	}

	return result, err
}

func (s *HistoryService) GetUpsertVisitor(contentId string, userId string) bool {
	userIdInt, _ := strconv.Atoi(userId)
	contentIdInt, _ := strconv.Atoi(contentId)
	code := userId + "-" + contentId
	return s.searchRepo.GetUpsertVisitor(contentIdInt, userIdInt, code)
}
