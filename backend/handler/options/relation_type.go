package options

import (
	options "sygap_new_knowledge_management/backend/services/options"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type RelationTypeHandler struct {
	service *options.RelationTypeService
	log     *logrus.Logger
}

func NewRelationTypeHandler(service *options.RelationTypeService, log *logrus.Logger) *RelationTypeHandler {
	return &RelationTypeHandler{service, log}
}

func (h *RelationTypeHandler) GetRelationType(c *fiber.Ctx) error {
	h.log.Println("Executtye function GetRelationType")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	listRelationType, err := h.service.GetListRelationType()
	if err != nil {
		h.log.Error("Failed get list product type in GetRelationType")
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:     "Failed get list product type",
			Error:       err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:     "Success",
		Data:        listRelationType,
	})

}