package options

import (
	"sygap_new_knowledge_management/backend/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type StatusRepo struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewStatusRepos(db *gorm.DB, log *logrus.Logger) *StatusRepo {
	return &StatusRepo{db, log}
}

func (r *StatusRepo) GetActiveStatus(reqType string) ([]model.ListStatus, error) {
	r.log.Println("Execute function GetActiveStatus")

	var status []model.ListStatus
	if reqType == "" {
		err1 := r.db.Table("status").Select("UPPER(status.status) AS status").Where("deleted_at IS NULL AND type = ? OR type = ?", "problem", "known_error").Group("status").
			Find(&status).Error
		if err1 != nil {
			r.log.Error("Failed get detail status in GetActiveStatus", err1)
			return nil, err1
		}
	} else {
		err := r.db.Table("status").Select("UPPER(status.status) AS status").Where("deleted_at IS NULL AND type = ?", reqType).
			Find(&status).Error
		if err != nil {
			r.log.Error("Failed get detail status in GetActiveStatus", err)
			return nil, err
		}
	}

	return status, nil
}

func (r *StatusRepo) GetAllActiveStatus() ([]model.ListStatus, error) {
	r.log.Println("Execute function GetActiveStatus")

	var status []model.ListStatus
	err := r.db.Table("status").Select("UPPER(status.status) AS status").Where("deleted_at IS NULL").Group("status.status").
		Find(&status).Error
	if err != nil {
		r.log.Error("Failed get detail status in GetActiveStatus", err)
		return nil, err
	}

	return status, nil
}
