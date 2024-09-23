package history

import (
	"errors"
	"github.com/sirupsen/logrus"
	"strconv"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/pkg/errs"
	"sygap_new_knowledge_management/backend/repository/search/history"
	"sygap_new_knowledge_management/backend/utils"
)

type SubmitService struct {
	repo *history.SubmitRepos
	log  *logrus.Logger
}

func NewSubmitService(repos *history.SubmitRepos, log *logrus.Logger) *SubmitService {
	return &SubmitService{repos, log}
}

func (s *SubmitService) SubmitRequestToUpdateKM(payload entities.RequestHistoryKnowledge, user string) error {
	if payload.Type != "workaround" && payload.Type != "fix-solution" && payload.Type != "reference" {
		return &errs.BadRequestError{
			Err: "Invalid value for field type",
		}
	}

	author, _ := strconv.Atoi(user)
	decodedKnowledgeContentId, _ := utils.GenerateDecoded(payload.KnowledgeContentID)
	knowledgeContentId, _ := strconv.Atoi(decodedKnowledgeContentId)

	var km entities.KnowledgeContent
	if err := s.repo.GetKnowledgeContentById(&km, knowledgeContentId); err != nil {
		s.log.Error(err.Error())
		return err
	}

	var kmHistory entities.HistoryKnowledge
	errGetLatestKMHistory := s.repo.GetLatestKnowledgeContentDetailHistoryByAuthorAndTypeAndKnowledgeContentId(&kmHistory, user, payload.Type, knowledgeContentId)

	if errGetLatestKMHistory != nil {
		var resourceNotFoundError *errs.ResourceNotFoundError
		if errors.As(errGetLatestKMHistory, &resourceNotFoundError) {
			historyKnowledge := &entities.HistoryKnowledge{
				KnowledgeContentId: knowledgeContentId,
				Note:               payload.Note,
				Type:               payload.Type,
				Value:              payload.Value,
				Status:             "Requested",
				Requestor:          user,
				CreatedBy:          author,
			}

			errSubmitToUpdateKnowledgeContent := s.repo.SubmitToUpdateKnowledgeContent(historyKnowledge)
			if errSubmitToUpdateKnowledgeContent != nil {
				return errSubmitToUpdateKnowledgeContent
			}

			return nil
		} else {
			s.log.Error(errGetLatestKMHistory.Error())
			return errGetLatestKMHistory
		}
	} else {
		kmHistory.Value = payload.Value
		kmHistory.Note = payload.Note

		if err := s.repo.UpdateKMHistory(&kmHistory); err != nil {
			s.log.Error(err.Error())
			return err
		}

		return nil

	}
}
