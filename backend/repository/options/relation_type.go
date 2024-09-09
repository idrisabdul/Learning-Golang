package options

import (
	"sygap_new_knowledge_management/backend/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RelationTypeRepo struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewRelationTypeRepos(db *gorm.DB, log *logrus.Logger) *RelationTypeRepo {
	return &RelationTypeRepo{db, log}
}

func (r *RelationTypeRepo) GetListRelationType() ([]model.ListRelationType, error) {
	r.log.Println("Execute function GetDetailRelationType")

	var detailRelationType []model.ListRelationType
	err := r.db.Table("relation_type").Where("deleted_by IS NULL").Find(&detailRelationType).Error
	if err != nil {
		r.log.Error("Failed to get detail relation type in GetDetailRelationType", err)
		return nil, err
	}

	return detailRelationType, nil
}
