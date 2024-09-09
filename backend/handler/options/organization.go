package options

import (
	options "sygap_new_knowledge_management/backend/services/options"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type OrganizationHandler struct {
	service *options.OrganizationService
	log     *logrus.Logger
}

func NewOrganizationHandler(service *options.OrganizationService, log *logrus.Logger) *OrganizationHandler {
	return &OrganizationHandler{service, log}
}

func (h *OrganizationHandler) GetOrganizationHandler(c *fiber.Ctx) error {
	h.log.Println("Execute function GetOrganizationHandler")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	search := c.Query("search")
	idCompany := c.Query("id_company")
	isAll := c.Query(("is_all"))

	dataOrganization, err := h.service.GetActiveOrganization(search, idCompany, isAll)
	if err != nil {
		h.log.Println("Failed get data organization in GetActiveOrganization")

		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:     "Error",
			Error:       err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:     "Success",
		Data:        dataOrganization,
	})
}
