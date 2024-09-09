package km

import (
	"strconv"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/services/km"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type KMListHandler struct {
	service *km.KMListService
	log     *logrus.Logger
}

func NewKMListHandler(service *km.KMListService, log *logrus.Logger) *KMListHandler {
	return &KMListHandler{service, log}
}

func (h *KMListHandler) GetListKM(c *fiber.Ctx) error {

	h.log.Print("Execute GetListKM Function")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}


	var request entities.SearchListKM
	if errRequest := c.BodyParser(&request); errRequest != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to retrieve KM List",
			Error:      errRequest.Error(),
		})
	}

	listKM, errListKM := h.service.GetListKM(request, permission)
	if errListKM != nil {
		h.log.Printf("Execute ListKM Function: %v", errListKM.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to retrieve KM List",
			Error:      errListKM.Error(),
		})
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("per_page", "10"))

	data, metaPagination := utils.GetPaginated(c, page, limit, listKM)

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieve data",
		Data:       data,
		Meta:       utils.Meta{Pagination: metaPagination},
	})
}
