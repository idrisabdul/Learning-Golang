package history

import (
	"errors"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/pkg/errs"
	"sygap_new_knowledge_management/backend/utils"

	// "github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HistoryRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewHistoryRepo(db *gorm.DB, logger *logrus.Logger) *HistoryRepository {
	return &HistoryRepository{db, logger}
}

func (r *HistoryRepository) GetHistoryApprove(decodedIDKM string) ([]entities.HistoryKnowledgeList, error) {
	var history []entities.HistoryKnowledgeList

	if errGetHistoryDetail := r.db.Table(utils.TABLE_HISTORY_KNOWLEDGE).Select(`
		hk.id,
		hk.knowledge_content_id,
		hk.note,
		hk.type,
		hk.value,
		hk.status,
		hk.date,
		hk.updated_at,
		hk.created_at,
		hk.created_by,
		hk.deleted_at,
		hk.deleted,
		e1.employee_name AS requestor,
		e2.employee_name AS updated_by
	`).
		Joins("LEFT JOIN "+utils.TABLE_EMPLOYEE+"1 ON hk.requestor = e1.id").
		Joins("LEFT JOIN "+utils.TABLE_EMPLOYEE+"2 ON hk.updated_by = e2.id").
		Where("hk.knowledge_content_id = ?", decodedIDKM).
		Where("hk.status = 'Approved'").
		Order("hk.updated_at DESC").
		Find(&history).Error; errGetHistoryDetail != nil {
		return nil, errGetHistoryDetail
	}
	return history, nil
}

func (r *HistoryRepository) GetHistoryApproveReject(decodedIDKM string) ([]entities.HistoryKnowledgeList, error) {
	var history []entities.HistoryKnowledgeList

	if errGetHistoryDetail := r.db.Table(utils.TABLE_HISTORY_KNOWLEDGE).Select(`
		hk.id,
		hk.knowledge_content_id,
		hk.note,
		hk.type,
		hk.value,
		hk.status,
		hk.date,
		hk.updated_at,
		hk.created_at,
		hk.created_by,
		hk.deleted_at,
		hk.deleted,
		e1.employee_name AS requestor,
		e2.employee_name AS updated_by
	`).
		Joins("LEFT JOIN "+utils.TABLE_EMPLOYEE+"1 ON hk.requestor = e1.id").
		Joins("LEFT JOIN "+utils.TABLE_EMPLOYEE+"2 ON hk.updated_by = e2.id").
		Where("hk.knowledge_content_id = ?", decodedIDKM).
		Where("hk.status in ('Approved', 'Rejected')").
		Order("hk.updated_at DESC").
		Find(&history).Error; errGetHistoryDetail != nil {
		return nil, errGetHistoryDetail
	}
	return history, nil
}

func (r *HistoryRepository) GetHistoryRequested(decodedIDKM string) ([]entities.HistoryNotif, error) {
	var history []entities.HistoryNotif

	if errGetHistoryDetail := r.db.Table(utils.TABLE_HISTORY_KNOWLEDGE).Select(`
			hk.id,
			concat(ifnull(e.employee_name, '')," Requested to update ",ifnull(hk.type, ''),", ",ifnull(hk.note, '')) as notif
		`).
		Joins("LEFT JOIN "+utils.TABLE_EMPLOYEE+" ON hk.requestor = e.id").
		Where("hk.status = 'Requested'").
		Where("hk.knowledge_content_id = ?", decodedIDKM).
		Find(&history).Error; errGetHistoryDetail != nil {
		return nil, errGetHistoryDetail
	}
	return history, nil
}

func (r *HistoryRepository) GetHistoryKnowledgeByIdAndKMContentId(data *entities.HistoryKnowledge, kmDetailHistoryId int, kmContentId int) error {
	if err := r.db.
		Where("id = ?", kmDetailHistoryId).
		Where("knowledge_content_id = ?", kmContentId).
		First(&data, kmDetailHistoryId).Error; err != nil {
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
