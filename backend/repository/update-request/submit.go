package updaterequest

import (
	"sygap_new_knowledge_management/backend/entities"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SubmitRepos struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewSubmitRepos(db *gorm.DB, log *logrus.Logger) *SubmitRepos {
	return &SubmitRepos{db, log}
}

// Submit Function
func (r *SubmitRepos) SubmitToKnowledgeContentUpdateRequest(data entities.KnowledgeContentUpdateRequest) (int, error) {

	if errSubmitToKnowledgeUpdateRequest := r.db.Create(&data).Error; errSubmitToKnowledgeUpdateRequest != nil {
		return 0, errSubmitToKnowledgeUpdateRequest
	}

	return data.ID, nil
}

func (r *SubmitRepos) SubmitToKnowledgeUpdateReportAttachment(data []entities.KnowledgeContentReportAttachment) error {

	if errSubmitToKnowledgeUpdateReportAttachment := r.db.CreateInBatches(&data, len(data)).Error; errSubmitToKnowledgeUpdateReportAttachment != nil {
		return errSubmitToKnowledgeUpdateReportAttachment
	}

	return nil
}

func (r *SubmitRepos) SubmitToKnowledgeContentUpdateRequestComment() {

}
