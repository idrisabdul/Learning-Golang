package updaterequest

import (
	"sygap_new_knowledge_management/backend/entities"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UpdateRepos struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewUpdateRepos(db *gorm.DB, log *logrus.Logger) *UpdateRepos {
	return &UpdateRepos{db, log}
}

func (r *UpdateRepos) UpdateToKnowledgeContentUpdateRequest(data entities.KnowledgeContentUpdateRequest) error {

	if errUpdateToKnowledgeContentUpdateRequest := r.db.Where("id = ?", data.ID).Updates(&data).Error; errUpdateToKnowledgeContentUpdateRequest != nil {
		return errUpdateToKnowledgeContentUpdateRequest
	}
	return nil
}

func (r *UpdateRepos) UpdateToKnowledgeContentReportAttachment(data entities.KnowledgeContentReportAttachment) error {

	if errUpdateToKnowledgeContentReportAttachment := r.db.Where("id = ?", data.ID).Updates(&data).Error; errUpdateToKnowledgeContentReportAttachment != nil {
		return errUpdateToKnowledgeContentReportAttachment
	}
	return nil
}

func (r *UpdateRepos) InstanceSubmitToKnowledgeContentReportAttachment(data []entities.KnowledgeContentReportAttachment) error {
	return NewSubmitRepos(r.db, r.log).SubmitToKnowledgeUpdateReportAttachment(data)
}
