package options

import (
	"sygap_new_knowledge_management/backend/services/options"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CompanyHandler struct {
	service *options.CompanyService
	log     *logrus.Logger
}

func NewCompanyHandler(service *options.CompanyService, log *logrus.Logger) *CompanyHandler {
	return &CompanyHandler{service, log}
}

func (h *CompanyHandler) GetCompanies(c *fiber.Ctx) error {
	h.log.Print("Execute GetCompanies Function")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	companies, errGetCompanies := h.service.GetCompanies()
	if errGetCompanies != nil {
		h.log.Printf("Execute GetCompanies Function: %v", errGetCompanies.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to retrieve company List",
			Error:      errGetCompanies.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Success Retrieve Company List",
		Data:      companies,
	})
}
