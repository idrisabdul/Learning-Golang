package options

import (
	"sygap_new_knowledge_management/backend/services/options"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ContentTypeHandler struct {
	service *options.ContentTypeService
	log *logrus.Logger
}

func NewContentTypeHandler(service *options.ContentTypeService, log *logrus.Logger) *ContentTypeHandler {
	return &ContentTypeHandler{service, log}
}

func (h *ContentTypeHandler) GetContentType(c *fiber.Ctx) error {
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	ContentType, errGetContentType := h.service.GetContentType()
	if errGetContentType != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Error while GetContentType",
			Error:      errGetContentType.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieve expertee list",
		Data:       ContentType,
	})
}