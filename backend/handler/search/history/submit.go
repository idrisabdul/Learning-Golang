package history

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/services/search/history"
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

func (h *SubmitHistoryHandler) SubmitRequestToUpdateKM(c *fiber.Ctx) error {
	h.log.Print("Execute SubmitRequestToUpdateKM function on Handler")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	author := permission["user"].(map[string]any)["id"].(string)

	validate := validator.New()

	var request entities.RequestHistoryKnowledge

	// parse request from Front End
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

	errSubmit := h.submit.SubmitRequestToUpdateKM(request, author)
	if errSubmit != nil {
		return utils.ResponseWithError(c, errSubmit)
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully Submit data",
		Data:       "success",
	})
}
