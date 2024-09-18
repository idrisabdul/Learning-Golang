package history

import (
	"strconv"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/services/km/history"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SubmitHistoryHandler struct {
	submit *history.SubmitService
	log    *logrus.Logger
}

func NewHistoryhandler(submit *history.SubmitService, log *logrus.Logger) *SubmitHistoryHandler {
	return &SubmitHistoryHandler{submit, log}
}

func (h *SubmitHistoryHandler) ApprovalKM(c *fiber.Ctx) error {
	h.log.Print("Execute ApprovalKM function on Handler")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	author, _ := strconv.Atoi(permission["user"].(map[string]any)["id"].(string))

	idKMHistory, _ := strconv.Atoi(c.Params("id"))

	var request entities.ApprovalKM
	validate := validator.New()

	if errParseToPayload := c.BodyParser(&request); errParseToPayload != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed Submit data",
			Error:      errParseToPayload.Error(),
		})
	}

	if err := validate.Struct(request); err != nil {
		return utils.ResponseRequestValidationError(c, err)
	}

	errSubmit := h.submit.ApprovalKM(author, idKMHistory, request.ApprovedStatus)
	if errSubmit != nil {
		return utils.ResponseWithError(c, errSubmit)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully Submit data",
		Data:       "success",
	})
}
