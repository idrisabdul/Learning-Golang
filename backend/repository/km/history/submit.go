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

func (r *SubmitRepos) GetKnowledgeContentDetailHistoryById(data *entities.HistoryKnowledge, kmDetailHistoryId int) error {
	if err := r.db.First(&data, kmDetailHistoryId).Error; err != nil {
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

func (r *SubmitRepos) GetKnowledgeContentDetailByKnowledgeContentId(data *entities.KnowledgeContentDetail, knowledgeContentId int) error {
	if err := r.db.Where("knowledge_content_id = ?", knowledgeContentId).First(&data).Error; err != nil {
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

func (r *SubmitRepos) ApprovalKM(author int, approvedStatus string, dataKMHistory *entities.HistoryKnowledge, dataKMDetail *entities.KnowledgeContentDetail, dataKM *entities.KnowledgeContent) error {
	tx := r.db.Begin()

	if tx.Error != nil {
		r.log.Fatal("Failed to begin transaction:", tx.Error)
		return tx.Error
	}

	if err := updateContainKnowledgeContain(author, tx, approvedStatus, dataKMHistory, dataKMDetail, dataKM); err != nil {
		tx.Rollback()
		return tx.Error
	}

	if err := tx.Commit().Error; err != nil {
		r.log.Info("Success update content knowledge contain")
		return tx.Error
	}

	return nil
}

func updateContainKnowledgeContain(author int, tx *gorm.DB, approvedStatus string, dataKMHistory *entities.HistoryKnowledge, dataKMDetail *entities.KnowledgeContentDetail, dataKM *entities.KnowledgeContent) error {
	now := time.Now()
	if approvedStatus == "approved" {
		var kmDetail interface{}
		if dataKMHistory.Type == "workaround" {
			kmDetail = map[string]interface{}{"workaround": &dataKMHistory.Value}
		} else if dataKMHistory.Type == "fix-solution" {
			kmDetail = map[string]interface{}{"fix_solution": &dataKMHistory.Value}
		} else if dataKMHistory.Type == "reference" {
			kmDetail = map[string]interface{}{"reference": &dataKMHistory.Value}
		}

		version := dataKM.Version + 1

		kmHistory := map[string]interface{}{
			"status":     "Approved",
			"date":       now,
			"updated_at": now,
			"updated_by": author,
		}

		if err := tx.Model(&dataKM).Updates(map[string]interface{}{
			"version":    version,
			"updated_at": now,
			"updated_by": author,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Model(&dataKMDetail).Updates(kmDetail).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Model(&dataKMHistory).Updates(kmHistory).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if err := tx.Model(&dataKMHistory).Updates(map[string]interface{}{
			"status":     "Rejected",
			"date":       now,
			"updated_at": now,
			"updated_by": author,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}
