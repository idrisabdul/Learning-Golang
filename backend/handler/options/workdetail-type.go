package options

import (
	"sygap_new_knowledge_management/backend/services/options"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type WorkDetailTypeHandler struct {
	service *options.WorkDetailTypeService
	log     *logrus.Logger
}

func NewWorkDetailTypeHandler(service *options.WorkDetailTypeService, log *logrus.Logger) *WorkDetailTypeHandler {
	return &WorkDetailTypeHandler{service, log}
}

func (h *WorkDetailTypeHandler) GetWorkDetailType(c *fiber.Ctx) error {
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	workDetailType, errWorkDetailType := h.service.GetWorkDetailType()

	if errWorkDetailType != nil {
		h.log.Printf("Failed to retrieve data: %v", errWorkDetailType.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to retrieve status List",
			Error:      errWorkDetailType.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Success Retrieve status List",
		Data:       workDetailType,
	})
}
