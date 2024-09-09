package options

import (
	"strconv"
	"strings"
	"sygap_new_knowledge_management/backend/entities"
	masterrepo "sygap_new_knowledge_management/backend/repository/options"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
)

type ProductTypeService struct {
	repo *masterrepo.ProductTypeRepo
	log  *logrus.Logger
}

func NewProductTypeService(repo *masterrepo.ProductTypeRepo, log *logrus.Logger) *ProductTypeService {
	return &ProductTypeService{repo, log}
}

func (s *ProductTypeService) GetListProductType(IDProductName string, isAll string) ([]entities.ListServiceType, error) {
	s.log.Println("Execute function GetListProductType")

	IDProductNameDecoded, err := utils.GenerateDecoded(IDProductName)
	if err != nil {
		s.log.Error("Failed to decoded id product name in GetListProductType", err)
		return nil, err
	}

	var IDProdName int
	if IDProductName != "" {
		IDProductNameInt, err := strconv.Atoi(IDProductNameDecoded)
		if err != nil {
			s.log.Error("Failed to convert int id product name in GetListProductType", err)
			return nil, err
		}
		IDProdName = IDProductNameInt
	}

	var listProduct []entities.ListServiceType
	if isAll == "true" {
		listProdType, _ := s.repo.GetAllDetailProductType(IDProdName)
		listProduct = listProdType
	} else {
		listProdType, _ := s.repo.GetDetailProductType(IDProdName)
		listProduct = listProdType
	}
	return listProduct, nil
}

func (s *ProductTypeService) GetListProductTypeRelation(requestType string, search string, isAll string) ([]entities.ListServiceType, error) {
	s.log.Println("Execute function GetListProductType")

	var moduleId int
	if strings.ToLower(requestType) == "incident" {
		moduleId = 1
	} else if strings.ToLower(requestType) == "problem investigation" {
		moduleId = 3
	} else if strings.ToLower(requestType) == "infrastructure change" {
		moduleId = 4
	} else if strings.ToLower(requestType) == "known error" {
		moduleId = 3
	} else if strings.ToLower(requestType) == "request fulfillment" {
		moduleId = 2
	} else {
		moduleId = 6
	}

	var result []entities.ListServiceType

	if isAll == "true" {
		result, _ = s.repo.GetAllDetailProductTypeRelation(search)
	} else {
		result, _ = s.repo.GetDetailProductTypeRelation(moduleId, search)
	}

	return result, nil
}
