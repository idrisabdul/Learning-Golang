package feedback_repo

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FeedbackRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewFeedbackRepository(db *gorm.DB, logger *logrus.Logger) *FeedbackRepository {
	return &FeedbackRepository{db, logger}
}

func (p *FeedbackRepository) GetFeedbackList(km_id string) ([]entities.FeedbackList, error) {
	var feedback []entities.FeedbackList
	if err := p.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_FEEDBACK).Where("kcf.knowledge_id = ?", km_id).Select(`
		kcf.id, kcf.knowledge_id, kcf.submitter_id, kcf.usefull, kcf.rating, kcf.comment, kcf.date_submit,
		kc.knowledge_id as code, e.employee_name as submitter
	`).Joins(`
		LEFT JOIN knowledge_content kc ON kcf.knowledge_id = kc.id
	`).Joins(`
	LEFT JOIN employee e ON kcf.submitter_id = e.id
	`).Order("kcf.id DESC").
		Find(&feedback).Error; err != nil {
		p.logger.Error("Error retrieving services: ", err)
		return nil, err
	}

	return feedback, nil
}

func (p *FeedbackRepository) GetDetailFeedback(km_id string) (entities.ListKM, error) {
	p.logger.Println("Execute function GetDetailFeedback")

	var detailFeedback entities.ListKM
	err := p.db.Table("feedback").Where("id = ?", km_id).Find(&detailFeedback).Error
	if err != nil {
		p.logger.Error("failed to get detail feedback", err)
		return entities.ListKM{}, err
	}

	return detailFeedback, nil
}

func (p *FeedbackRepository) GetKnowledgeContent(km_id string) (entities.KnowledgeContent, error) {
	var detailKC entities.KnowledgeContent
	err := p.db.Table("knowledge_content").Where("id = ?", km_id).
		Find(&detailKC).Error
	if err != nil {
		p.logger.Error("Failed to get detail knowledge content", err)
		return entities.KnowledgeContent{}, nil
	}

	return detailKC, nil
}
