package options

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/repository/options"

	"github.com/sirupsen/logrus"
)

type ProductService struct {
	repo *options.ProductRepos
	log *logrus.Logger
}

func NewProductService(repo *options.ProductRepos, log *logrus.Logger) *ProductService {
	return &ProductService{repo, log}
}

func (s *ProductService) GetListProduct(idCompany, search string) ([]entities.ListProduct, error) {
	return s.repo.GetListProduct(idCompany,search)
}

func (s *ProductService) GetProductParentCategory(idCategory string) (entities.Product, error) {
	return s.repo.GetProductParentCategory(idCategory)
}