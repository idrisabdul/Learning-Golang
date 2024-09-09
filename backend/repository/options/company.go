package options

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CompanyRepos struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewCompanyRepos(db *gorm.DB, log *logrus.Logger) *CompanyRepos {
	return &CompanyRepos{db, log}
}

func (r *CompanyRepos) GetCompanies() ([]entities.Company, error) {
	var companies []entities.Company
	if errGetCompanies := r.db.Table(utils.TABLE_COMPANY).Find(&companies).Error; errGetCompanies != nil {
		r.log.Errorln("Error retrieving list companies: ", errGetCompanies)
		return nil, errGetCompanies
	}

	return companies, nil
}
