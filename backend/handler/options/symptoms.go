package options

import (
	options "sygap_new_knowledge_management/backend/services/options"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"

	"github.com/sirupsen/logrus"
)

type SymptomsHandler struct {
	service *options.SymptomsService
	log     *logrus.Logger
}

func NewSymptomsHandler(service *options.SymptomsService, log *logrus.Logger) *SymptomsHandler {
	return &SymptomsHandler{service, log}
}

func (h *SymptomsHandler) GetSymptomsHandler(c *fiber.Ctx) error {
	h.log.Println("Executtye function GetSymptomsHandler")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	IDProductName := c.Query("id_product_name")
	IDProductType := c.Query("id_product_type")
	IDCompany := c.Query("id_company")
	search := c.Query("search")

	listSymptoms, err := h.service.GetListSymptoms(IDProductName, IDProductType, IDCompany, search)
	if err != nil {
		h.log.Error("Failed get list symptoms in GetSymptomsHandler")

		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:     "Error",
			Error:       err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:     "Success",
		Data:        listSymptoms,
	})

}

func (h *SymptomsHandler) GetSymptomsRelationHandler(c *fiber.Ctx) error {
	h.log.Println("Executtye function GetSymptomsRelationHandler")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	listSymptoms, err := h.service.GetListSymptomsRelation(c)
	if err != nil {
		h.log.Error("Failed get list symptoms in GetSymptomsRelationHandler")

		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:     "Error",
			Error:       err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:     "Success",
		Data:        listSymptoms,
	})

}
