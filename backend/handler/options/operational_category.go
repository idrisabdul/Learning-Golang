package options

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/services/options"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type OperationalCategorysHandler struct {
	service *options.OperationalCategoryService
	log     *logrus.Logger
}

func NewOperationalCategorysHandler(service *options.OperationalCategoryService, log *logrus.Logger) *OperationalCategorysHandler {
	return &OperationalCategorysHandler{service: service, log: log}
}

func (h *OperationalCategorysHandler) GetOpCat(c *fiber.Ctx) error {

	var ListOpCat []entities.OperationCategory
	var errListOpCat error

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	idParent := c.Params("idParent")
	if idParent != "" {
		ListOpCat, errListOpCat = h.service.GetOpCat2(idParent)
		if errListOpCat != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "Error while GetOpCat2",
				Error:      errListOpCat.Error(),
			})
		}
	} else {
		ListOpCat, errListOpCat = h.service.GetOpCat1()
		if errListOpCat != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "Error while GetOpCat1",
				Error:      errListOpCat.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieve operational category",
		Data:       ListOpCat,
	})
}
