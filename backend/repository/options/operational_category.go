package options

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OperationalCategoryRepos struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewOperationalCategorysRepos(db *gorm.DB, log *logrus.Logger) *OperationalCategoryRepos {
	return &OperationalCategoryRepos{db, log}
}

func (r *OperationalCategoryRepos) GetOpCat1() ([]entities.OperationCategory, error) {
	r.log.Print("Execute GetOpCat1 Function on Repo")

	var OpCat1 []entities.OperationCategory

	if errGetOpCat1 := r.db.Table(utils.TABLE_OPERATIONAL_CATEGORIES).
		Where("oc.deleted_at IS NULL").
		Where("oc.parent_id = 0").
		Find(&OpCat1).Error; errGetOpCat1 != nil {
		r.log.Println("Error while GetOpCat1 on Repo: ", errGetOpCat1)
		return nil, errGetOpCat1
	}

	return OpCat1, nil
}

func (r *OperationalCategoryRepos) GetOpCat2(idParent string) ([]entities.OperationCategory, error) {
	r.log.Print("Execute GetOpCat2 Function on Repo")

	var OpCat2 []entities.OperationCategory

	if errGetOpCat2 := r.db.Table(utils.TABLE_OPERATIONAL_CATEGORIES).
		Where("oc.deleted_at IS NULL").
		Where("oc.parent_id = ?", idParent).
		Find(&OpCat2).Error; errGetOpCat2 != nil {
		r.log.Println("Error while GetOpCat2 on Repo: ", errGetOpCat2)
		return nil, errGetOpCat2
	}

	return OpCat2, nil
}
