package options

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UpdateRequestRepos struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewUpdateRequestRepos(db *gorm.DB, log *logrus.Logger) *UpdateRequestRepos {
	return &UpdateRequestRepos{db, log}
}

func (r UpdateRequestRepos) GetListUpdateRequestStatus() ([]entities.ListupdateRequestStatus, error) {
	r.log.Print("Execute GetListUpdateRequestStatus Function on Repo")

	var response []entities.ListupdateRequestStatus
	if errGetListUpdateRequestStatus := r.db.Table(utils.TABLE_STATUS).
		Where("s.type = 'knowledge_update_req'").
		Order("s.id ASC").
		Find(&response).Error; errGetListUpdateRequestStatus != nil {
		r.log.Errorln("Failed to execute GetListUpdateRequestStatus Function on Repo: ", errGetListUpdateRequestStatus)
		return nil, errGetListUpdateRequestStatus
	}

	return response, nil
}