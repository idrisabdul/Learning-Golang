package history

import (
	"errors"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/pkg/errs"
	"time"

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

func (r *SubmitRepos) SubmitToUpdateKnowledgeContent(data *entities.HistoryKnowledge) error {
	r.log.Print("Execute SubmitToUpdateKnowledgeContent Function")

	if errSubmitToUpdateKnowledgeContent := r.db.Create(&data).Error; errSubmitToUpdateKnowledgeContent != nil {
		r.log.Errorln("Failed to SubmitToUpdateKnowledgeContent: ", errSubmitToUpdateKnowledgeContent.Error())
		return errSubmitToUpdateKnowledgeContent
	}

	return nil
}

func (r *SubmitRepos) GetKnowledgeContentById(data *entities.KnowledgeContent, knowledgeContentId int) error {
	if err := r.db.First(&data, knowledgeContentId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &errs.ResourceNotFoundError{
				Err: "Data Not Found",
			}
		} else {
			return err
		}
	}

	return nil
}

func (r *SubmitRepos) GetLatestKnowledgeContentDetailHistoryByAuthorAndTypeAndKnowledgeContentId(data *entities.HistoryKnowledge, author string, typeValue string, knowledgeContentId int) error {
	err := r.db.
		Where("requestor = ?", author).
		Where("type = ?", typeValue).
		Where("knowledge_content_id = ?", knowledgeContentId).
		Where("status = ?", "Requested").
		Order("created_at DESC").
		First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &errs.ResourceNotFoundError{
				Err: "Data Not Found",
			}
		} else {
			return err
		}
	}

	return nil
}

func (r *SubmitRepos) UpdateKMHistory(data *entities.HistoryKnowledge) error {
	now := time.Now()
	if err := r.db.Model(&data).UpdateColumns(map[string]interface{}{
		"value":      data.Value,
		"note":       data.Note,
		"created_at": now,
	}).Error; err != nil {
		return err
	}

	return nil
}
