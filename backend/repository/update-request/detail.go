package updaterequest

import (
	"fmt"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DetailRepos struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewDetailRepos(db *gorm.DB, log *logrus.Logger) *DetailRepos {
	return &DetailRepos{db, log}
}

func (r *DetailRepos) GetListUpdateRequest(idKM string) ([]entities.ListUpdateRequest, error) {
	var updateRequests []entities.ListUpdateRequest

	subQuery := r.db.Table(utils.TABLE_STATUS).Select("s.id").Where("LOWER(s.status) = LOWER(kcur.status)").Where("s.type = 'knowledge_update_req'")

	if errGetUpdateRequest := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_UPDATE_REQUEST).
		Select(`
		kcur.*,
		(?) AS status_id, 
		e.username AS submitter_username, 
		e.employee_name AS submitter_name
		`, subQuery).
		Joins("LEFT JOIN " + utils.TABLE_EMPLOYEE + " ON kcur.submitter_id = e.id").
		Joins("LEFT JOIN " + utils.TABLE_KNOWLEDGE_CONTENT_LOG_VERSION + " ON kcur.id = kclv.key_content").
		Where(r.db.Where("kcur.knowledge_id = ?", idKM).Or("kclv.key_content = ?", idKM)).
		Where("kcur.deleted_at IS NULL").
		Order("kcur.submit_date DESC").
		Find(&updateRequests).Error; errGetUpdateRequest != nil {
		r.log.Errorln("Error on GetListUpdateRequest in Repo: ", errGetUpdateRequest)
		return nil, errGetUpdateRequest
	}

	return updateRequests, nil
}

func (r *DetailRepos) GetDetailUpdateRequest(idUpdateRequest string) (entities.DetailUpdateRequest, error) {

	var detailUpdateRequest entities.DetailUpdateRequest

	subQuery := r.db.Table(utils.TABLE_STATUS).Select("s.id").Where("LOWER(s.status) = LOWER(kcur.status)").Where("s.type = 'knowledge_update_req'")

	fmt.Printf("subQuery: %v\n", subQuery)

	if errGetDetailUpdateRequest := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_UPDATE_REQUEST).
		Select(`
		kcur.*,
		(?) AS status_id, 
		e.employee_name AS submitter
		`, subQuery).
		Joins("LEFT JOIN "+utils.TABLE_EMPLOYEE+" ON kcur.submitter_id = e.id").
		Joins("LEFT JOIN "+utils.TABLE_KNOWLEDGE_CONTENT_LOG_VERSION+" ON kcur.id = kclv.key_content").
		Where("kcur.id = ?", idUpdateRequest).
		Find(&detailUpdateRequest).Error; errGetDetailUpdateRequest != nil {
		r.log.Errorln("Error on GetListUpdateRequest in Repo: ", errGetDetailUpdateRequest)
		return entities.DetailUpdateRequest{}, errGetDetailUpdateRequest
	}

	return detailUpdateRequest, nil
}

func (r *DetailRepos) GetUpdateRequestAttachment(idUpdateRequest string) ([]entities.DetailDocument, error) {

	var response []entities.DetailDocument

	if errGetUpdateRequestAttachment := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_REPORT_ATTACHMENT).
		Where("kcra.knowledge_content_update_request_id = ?", idUpdateRequest).
		Where("deleted_at IS NULL").
		Find(&response).Error; errGetUpdateRequestAttachment != nil {
		return nil, errGetUpdateRequestAttachment
	}

	return response, nil
}
