package history

import (
	"strconv"
	services "sygap_new_knowledge_management/backend/services/history"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type HistoryHandler struct {
	service *services.HistoryService
	log     *logrus.Logger
}

func NewHistoryListHandler(service *services.HistoryService, log *logrus.Logger) *HistoryHandler {
	return &HistoryHandler{
		service: service,
		log:     log,
	}
}

func (h *HistoryHandler) GetHistoryListApproved(c *fiber.Ctx) error {
	h.log.Print("Execute GetHistory List Approved Function")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	idKM := c.Params("id")

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	getHistoryDetail, err := h.service.GetHistoryListApprove(idKM)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to parse request body",
			Error:      err.Error(),
		})
	}

	data, metaPagination := utils.GetPaginated(c, page, limit, getHistoryDetail)

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieve data",
		Data:       data,
		Meta:       utils.Meta{Pagination: metaPagination},
	})
}

func (h *HistoryHandler) GetHistoryListApprovedReject(c *fiber.Ctx) error {
	h.log.Print("Execute GetHistory Approve and Reject Function")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	idKM := c.Params("id")

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	getHistoryDetail, err := h.service.GetHistoryListApproveReject(idKM)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to parse request body",
			Error:      err.Error(),
		})
	}

	data, metaPagination := utils.GetPaginated(c, page, limit, getHistoryDetail)

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieve data",
		Data:       data,
		Meta:       utils.Meta{Pagination: metaPagination},
	})
}

func (h *HistoryHandler) GetHistoryListRequested(c *fiber.Ctx) error {
	h.log.Print("Execute GetHistory Approve and Reject Function")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	idKM := c.Params("id")

	getHistoryDetail, err := h.service.GetHistoryListRequested(idKM)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to parse request body",
			Error:      err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieve data",
		Data:       getHistoryDetail,
	})
}
