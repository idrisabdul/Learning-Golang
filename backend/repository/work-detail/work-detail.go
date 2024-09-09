package workdetail

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/utils"

	// "github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WorkDetailRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewWorkDetail(db *gorm.DB, logger *logrus.Logger) *WorkDetailRepository {
	return &WorkDetailRepository{db, logger}
}

func (r *WorkDetailRepository) GetWorkDetail(idKM string) ([]entities.WorkDetailChangeResponse, error) {
	var work_detail []entities.WorkDetailChangeResponse

	if errGetWorkDetail := r.db.Table(utils.TABLE_WORKDETAIL_KNOWLEDGE_MANAGEMENT).Select(`
		wkm.id,
		wkm.type,
		wkm.note,
		wkm.created_at,
		file_hash,
		file_ori,
		e.employee_name
	`).
		Joins("LEFT JOIN "+ utils.TABLE_WORK_DETAIL_HAS_DOCUMENT + " ON wkm.id = wdhd.id_work_detail AND wdhd.type = 'knowledge'").
		Joins("LEFT JOIN "+utils.TABLE_EMPLOYEE+" ON wkm.created_by = e.id").
		Where("wkm.id_parent = ? OR (wdhd.type = 'knowledge' AND wdhd.file_hash IS NOT NULL)", idKM).
		Order("wkm.created_at DESC").
		Find(&work_detail).Error; errGetWorkDetail != nil {
		return nil, errGetWorkDetail
	}

	// CASE 
	// 		WHEN wkm.note NOT LIKE '%Status Updated To :%' THEN wdhd.file_hash
	// 		ELSE NULL
	// 	END AS file_hash,
	// 	CASE 
	// 		WHEN wkm.note NOT LIKE '%Status Updated To :%' THEN wdhd.file_ori
	// 		ELSE NULL
	// 	END AS file_ori,

	return work_detail, nil
}

func (r *WorkDetailRepository) SubmitToWorkdetailKnowledgeManagement(data entities.WorkdetailKnowledgeManagement) (int, error) {
	if errSubmitToWorkdetailKnowledgeManagement := r.db.Create(&data).Error; errSubmitToWorkdetailKnowledgeManagement != nil {
		return 0, errSubmitToWorkdetailKnowledgeManagement
	}
	return data.ID, nil
}

func (r *WorkDetailRepository) SubmitToWorkDetailHasDocument(data []entities.WorkDetailHasDocument) error {
	if errSubmitToWorkDetailHasDocument := r.db.CreateInBatches(&data, len(data)).Error; errSubmitToWorkDetailHasDocument != nil {
		return errSubmitToWorkDetailHasDocument
	}
	return nil
}
