package options

import (
	options "sygap_new_knowledge_management/backend/services/options"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type StatusHandler struct {
	service *options.StatusService
	log     *logrus.Logger
}

func NewStatusHandler(service *options.StatusService, log *logrus.Logger) *StatusHandler {
	return &StatusHandler{service, log}
}

func (h *StatusHandler) GetOptionStatus(c *fiber.Ctx) error {
	h.log.Println("Execute function GetOptionStatus")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	detailStatus, err := h.service.GetActiveStatus(c.Query("request_type"), c.Query("is_all"))
	if err != nil {
		h.log.Error("Failed to call function get active status in GetOptionStatus", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:     "Error",
			Error:       err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:     "Success get detail status",
		Data:        detailStatus,
	})

}
