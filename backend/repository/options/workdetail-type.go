package options

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WorkDetailTypeRepos struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewWorkDetailTypeRepos(db *gorm.DB, log *logrus.Logger) *WorkDetailTypeRepos {
	return &WorkDetailTypeRepos{db, log}
}

func (r *WorkDetailTypeRepos) GetWorkDetailType(id int) ([]entities.WorkDetailType, error) {

	var response []entities.WorkDetailType
	if errGetWorkDetailType := r.db.Table(utils.TABLE_WORK_DETAIL_TYPE).Where("wdt.parent_id = ?", id).Find(&response).Error; errGetWorkDetailType != nil {
		return nil, errGetWorkDetailType
	}

	return response, nil
}
