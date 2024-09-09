package options

import (
	"strings"
	"sygap_new_knowledge_management/backend/model"
	options "sygap_new_knowledge_management/backend/repository/options"

	"github.com/sirupsen/logrus"
)

type StatusService struct {
	repo *options.StatusRepo
	log  *logrus.Logger
}

func NewStatusService(repo *options.StatusRepo, log *logrus.Logger) *StatusService {
	return &StatusService{repo, log}
}

func (s *StatusService) GetActiveStatus(reqType string, isAll string) ([]model.ListStatus, error) {

	var Type string

	if strings.ToLower(reqType) == "incident" {
		Type = "incident"
	} else if strings.ToLower(reqType) == "infrastructure change" {
		Type = "change"
	} else if strings.ToLower(reqType) == "problem investigation" {
		Type = "problem"
	} else if strings.ToLower(reqType) == "known error" {
		Type = "known_error"
	} else if strings.ToLower(reqType) == "request fullfillment" {
		Type = "request"
	} else if strings.ToLower(reqType) == "knowledge" {
		Type = "knowledge"
	} else {
		Type = ""
	}

	var finalStatus []model.ListStatus

	if isAll == "true" {
		allStatus, _ := s.repo.GetAllActiveStatus()
		finalStatus = allStatus
	} else {
		allStatus, _ := s.repo.GetActiveStatus(Type)
		finalStatus = allStatus
	}

	return finalStatus, nil
}
