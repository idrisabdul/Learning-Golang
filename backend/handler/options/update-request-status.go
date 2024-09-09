package options

import (
	"sygap_new_knowledge_management/backend/services/options"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UpdateRequestHandler struct {
	service *options.UpdateRequestService
	log     *logrus.Logger
}

func NewUpdateRequest(service *options.UpdateRequestService, log *logrus.Logger) *UpdateRequestHandler {
	return &UpdateRequestHandler{service, log}
}

func (h *UpdateRequestHandler) GetListUpdateRequestStatus(c *fiber.Ctx) error {
	h.log.Print("Execute GetListUpdateRequestStatus Function in Handler")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	updateRequestStatus, errUpdateRequestStatus := h.service.GetListUpdateRequestStatus()

	if errUpdateRequestStatus != nil {
		h.log.Printf("Failed to retrieve data: %v", errUpdateRequestStatus.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to retrieve status List",
			Error:      errUpdateRequestStatus.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Success Retrieve status List",
		Data:       updateRequestStatus,
	})
}
