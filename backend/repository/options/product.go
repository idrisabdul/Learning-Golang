package options

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepos struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewProductRepos(db *gorm.DB, log *logrus.Logger) *ProductRepos {
	return &ProductRepos{db, log}
}

func (r *ProductRepos) GetListProduct(idCompany, search string) ([]entities.ListProduct, error) {
	r.log.Print("Execute GetListProduct Function on Repo")

	var products []entities.ListProduct

	query := r.db.Table(utils.TABLE_PRODUCT_CATEGORIES).
		Select("pc.id, pc.name, pc.parent_id as category").
		Where("pc.tier = 3")

	if idCompany != "" {
		query = query.Where("pc.company_id = ?", idCompany)
	}

	if search != "" {
		query = query.Where("pc.name LIKE ?", "%"+search+"%")
	}

	if errGetListProduct := query.Find(&products).Error; errGetListProduct != nil {
		r.log.Errorln("Failed to execute GetListProduct on Repo: ", errGetListProduct)
		return nil, errGetListProduct
	}

	return products, nil
}

func (r *ProductRepos) GetProductParentCategory(idCategory string) (entities.Product, error) {
	r.log.Print("Execute GetProductParentCategory Function on Repo")

	var product entities.Product

	if errGetProduct := r.db.Table(utils.TABLE_PRODUCT_CATEGORIES).
		Where("pc.id = ?", idCategory).
		Find(&product).Error; errGetProduct != nil {
		r.log.Errorln("Failed to get product on GetProductParentCategory Function in Repo: ", errGetProduct)
		return entities.Product{}, errGetProduct
	}

	if errSetProductParentName := product.SetParentName(r.db); errSetProductParentName != nil {
		r.log.Errorln("Failed to set product on GetProductParentCategory Function in Repo: ", errSetProductParentName)
		return entities.Product{}, nil
	}

	return product, nil
}
