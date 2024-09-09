package options

import (
	"strconv"
	"sygap_new_knowledge_management/backend/model"
	options "sygap_new_knowledge_management/backend/repository/options"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SymptomsService struct {
	repo *options.SymptomsRepo
	log  *logrus.Logger
}

func NewSymptomsService(repo *options.SymptomsRepo, log *logrus.Logger) *SymptomsService {
	return &SymptomsService{repo, log}
}

func (s *SymptomsService) GetListSymptoms(IDProductName string, IDProductType string, IDCompany string, search string) ([]model.ListSymptoms, error) {
	s.log.Println("Execute function GetListSymptoms")

	IDProductNameDecoded, err := utils.GenerateDecoded(IDProductName)
	if err != nil {
		s.log.Error("Failed to decoded id product name in GetListSymptoms", err)
		return nil, err
	}

	var IDProdName int
	if IDProductName != "" {
		IDProductNameInt, err := strconv.Atoi(IDProductNameDecoded)
		if err != nil {
			s.log.Error("Failed to convert int id product name in GetListSymptoms", err)
			return nil, err
		}
		IDProdName = IDProductNameInt
	}

	IDProductTypeDecoded, err := utils.GenerateDecoded(IDProductType)
	if err != nil {
		s.log.Error("Failed to decoded id product type in GetListSymptoms", err)
		return nil, err
	}

	var IDProdType int
	if IDProductType != "" {
		IDProductTypeInt, err := strconv.Atoi(IDProductTypeDecoded)
		if err != nil {
			s.log.Error("Failed to convert int id product type in GetListSymptoms", err)
			return nil, err
		}
		IDProdType = IDProductTypeInt
	}

	IDCompanyDecoded, err := utils.GenerateDecoded(IDCompany)
	if err != nil {
		s.log.Error("Failed to decoded id company in GetListSymptoms", err)
		return nil, err
	}

	var IDCompanys int
	if IDCompany != "" {
		IDCompanyInt, err := strconv.Atoi(IDCompanyDecoded)
		if err != nil {
			s.log.Error("Failed to convert int id company", err)
			return nil, err
		}
		IDCompanys = IDCompanyInt
	}

	return s.repo.GetDetailSymptoms(IDProdName, IDProdType, IDCompanys, search)
}

func (s *SymptomsService) GetListSymptomsRelation(c *fiber.Ctx) ([]model.ListSymptoms, error) {
	s.log.Println("Execute function GetListSymptomsRelation")

	return s.repo.GetDetailSymptomsRelation(c.Query("id_product_type"), c.Query("search"))
}
