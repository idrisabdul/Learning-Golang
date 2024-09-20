package history

import (
	"strconv"
	"strings"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/pkg/errs"
	"sygap_new_knowledge_management/backend/repository/km/history"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
)

type SubmitService struct {
	repo *history.SubmitRepos
	log  *logrus.Logger
}

func NewSubmitService(repos *history.SubmitRepos, log *logrus.Logger) *SubmitService {
	return &SubmitService{repos, log}
}

func (s *SubmitService) ApprovalKM(author int, encodedIdHistoryKM string, approvedStatus string) error {
	var kmHistory entities.HistoryKnowledge
	var kmDetail entities.KnowledgeContentDetail
	var km entities.KnowledgeContent

	approvedStatusLower := strings.ToLower(approvedStatus)

	decodedIdHistoryKM, _ := utils.GenerateDecoded(encodedIdHistoryKM)
	idKMHistory, _ := strconv.Atoi(decodedIdHistoryKM)

	if approvedStatusLower != "rejected" && approvedStatusLower != "approved" {
		return &errs.BadRequestError{
			Err: "Invalid value for Approved Status",
		}
	}

	if errKmDetailHistory := s.repo.GetKnowledgeContentDetailHistoryById(&kmHistory, idKMHistory); errKmDetailHistory != nil {
		s.log.Errorln(errKmDetailHistory.Error())
		return errKmDetailHistory
	}

	if strings.ToLower(kmHistory.Status) != "requested" {
		return &errs.ForbiddenError{
			Err: "History had been approved or rejected",
		}
	}

	if errKmDetail := s.repo.GetKnowledgeContentDetailByKnowledgeContentId(&kmDetail, kmHistory.KnowledgeContentId); errKmDetail != nil {
		s.log.Errorln(errKmDetail.Error())
		return errKmDetail
	}

	if errKm := s.repo.GetKnowledgeContentById(&km, kmHistory.KnowledgeContentId); errKm != nil {
		s.log.Errorln(errKm.Error())
		return errKm
	}

	errApproval := s.repo.ApprovalKM(author, approvedStatusLower, &kmHistory, &kmDetail, &km)
	if errApproval != nil {
		return errApproval
	}

	return nil

}
