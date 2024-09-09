package options

import (
	"sygap_new_knowledge_management/backend/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrganizationRepo struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewOrganizationRepos(db *gorm.DB, log *logrus.Logger) *OrganizationRepo {
	return &OrganizationRepo{db, log}
}

func (r *OrganizationRepo) GetDetailOrganization(search string, idCompany string) ([]model.ListOrganization, error) {
	r.log.Println("Execute function GetDetailOrganization")
	var organization []model.ListOrganization
	var err error
	if search != "" {
		err = r.db.Table("organization").
			Where("deleted_at IS NULL AND parent_id = ? AND id_company = ?", 0, idCompany).
			Where("organization.organization_name LIKE ? COLLATE utf8_general_ci", "%"+search+"%"). // search query for lowwer/upper text
			Group("organization.id").
			Limit(20).
			Find(&organization).Error
	} else {
		err = r.db.Table("organization").
			Where("deleted_at IS NULL AND parent_id = ? AND id_company = ?", 0, idCompany).
			Group("organization.id").
			Limit(20).
			Find(&organization).Error
	}
	if err != nil {
		r.log.Error("Failed to get data for organization in GetDetailOrganization", err)
		return nil, err
	}

	return organization, nil
}

func (r *OrganizationRepo) GetAllDetailOrganization(search string) ([]model.ListOrganization, error) {
	r.log.Println("Execute function GetAllDetailOrganization")
	var organization []model.ListOrganization
	var err error
	if search != "" {
		err = r.db.Table("organization").
			Where("deleted_at IS NULL AND parent_id = ?", 0).
			Where("organization.organization_name LIKE ? COLLATE utf8_general_ci", "%"+search+"%"). // search query for lowwer/upper text
			Group("organization.organization_name").
			Limit(20).
			Find(&organization).Error
	} else {
		err = r.db.Table("organization").
			Where("deleted_at IS NULL AND parent_id = ?", 0).
			Group("organization.organization_name").
			Limit(20).
			Find(&organization).Error
	}
	if err != nil {
		r.log.Error("Failed to get data for organization in GetAllDetailOrganization", err)
		return nil, err
	}

	return organization, nil
}
