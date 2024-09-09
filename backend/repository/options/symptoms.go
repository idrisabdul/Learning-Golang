package options

import (
	"sygap_new_knowledge_management/backend/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SymptomsRepo struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewSymptomsRepos(db *gorm.DB, log *logrus.Logger) *SymptomsRepo {
	return &SymptomsRepo{db, log}
}

func (r *SymptomsRepo) GetDetailSymptoms(IDProductName int, IDProductType int, IDCompany int, search string) ([]model.ListSymptoms, error) {
	r.log.Println("Execute function GetDetailSymptoms")

	var listSymptoms []model.ListSymptoms
	var err error

	if search != "" {
		err = r.db.Table("symptoms").
			Where("deleted_at IS NULL and id_service = ? and id_service_type = ?", IDProductName, IDProductType).
			Where("symptoms.symptom_name LIKE ? COLLATE utf8_general_ci", "%"+search+"%").
			Group("symptoms.id").
			Limit(20).
			Find(&listSymptoms).Error
	} else {
		err = r.db.Table("symptoms").
			Where("deleted_at IS NULL and id_service = ? and id_service_type = ?", IDProductName, IDProductType).
			Group("symptoms.id").
			Limit(20).
			Find(&listSymptoms).Error
	}
	if err != nil {
		r.log.Error("Failed to get listSymptoms in GetDetailSymptoms", err)
		return nil, err
	}

	return listSymptoms, nil
}

func (r *SymptomsRepo) GetDetailSymptomsRelation(IDProductType string, search string) ([]model.ListSymptoms, error) {
	r.log.Println("Execute function GetDetailSymptomsRelation")

	var listSymptoms []model.ListSymptoms
	var err error

	if search != "" {
		err = r.db.Table("symptoms").
			Where("deleted_at IS NULL").
			Where("id_service_type = ?", IDProductType).
			Where("symptoms.symptom_name LIKE ? COLLATE utf8_general_ci", "%"+search+"%").
			Group("symptoms.symptom_name").
			Limit(20).
			Find(&listSymptoms).Error
	} else {
		err = r.db.Table("symptoms").
			Where("deleted_at IS NULL").
			Where("id_service_type = ?", IDProductType).
			Group("symptoms.symptom_name").
			Limit(20).
			Find(&listSymptoms).Error
	}
	if err != nil {
		r.log.Error("Failed to get listSymptoms in GetDetailSymptomsRelation", err)
		return nil, err
	}

	return listSymptoms, nil
}
