package options

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ExperteeRepos struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewExperteeRepos(db *gorm.DB, log *logrus.Logger) *ExperteeRepos {
	return &ExperteeRepos{db, log}
}

func (r *ExperteeRepos) GetExperteeGroup(idCompany string) ([]entities.Organization, error) {
	var experteeGroup []entities.Organization

	if errGetExperteeGroup := r.db.Table(utils.TABLE_ORGANIZATION).Where("deleted = 0").Where("id_company = ?",idCompany).Find(&experteeGroup).Error; errGetExperteeGroup != nil {
		r.log.Errorln("Error retrieving list expertee group: ", errGetExperteeGroup)
		return nil, errGetExperteeGroup
	}

	return experteeGroup, nil
}

func (r *ExperteeRepos) GetExpertees(idOrganization string) ([]entities.ListExpertees, error) {
	var expertees []entities.ListExpertees

	if errGetExpertees := r.db.Table(utils.TABLE_EMPLOYEE).Select("e.id,e.username,e.employee_name,oehr.id_organization, oehr.id_role").
		Joins("LEFT JOIN "+utils.TABLE_ORGANIZATION_EMPLOYEE_HAS_ROLE+" ON e.id = oehr.id_employee").
		Where("oehr.id_organization = ?", idOrganization).
		Where("oehr.id_role = 30"). // role = sme
		Find(&expertees).Error; errGetExpertees != nil {
		r.log.Errorln("Error retrieving list expertee group: ", errGetExpertees)
		return nil, errGetExpertees
	}

	return expertees, nil
}
