package options

import (
	"sygap_new_knowledge_management/backend/services/options"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ExperteeHandler struct {
	service *options.ExperteeService
	log     *logrus.Logger
}

func NewExperteeHandler(service *options.ExperteeService, log *logrus.Logger) *ExperteeHandler {
	return &ExperteeHandler{service, log}
}

func (h *ExperteeHandler) GetExperteeGroup(c *fiber.Ctx) error {

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	idCompany := c.Query("company")
	experteeGroup, errGetExperteeGroup := h.service.GetExperteeGroup(idCompany)
	if errGetExperteeGroup != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Error while GetExperteeGroup",
			Error:      errGetExperteeGroup.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieve expertee group",
		Data:       experteeGroup,
	})
}

func (h *ExperteeHandler) GetExpertees(c *fiber.Ctx) error {

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	idOrganization := c.Query("id_organization")
	expertees, errGetExpertees := h.service.GetExpertees(idOrganization)
	if errGetExpertees != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Error while GetExpertees",
			Error:      errGetExpertees.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieve expertee list",
		Data:       expertees,
	})
}
