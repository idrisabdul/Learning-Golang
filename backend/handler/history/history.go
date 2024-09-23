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

	encodedIDKM := c.Params("id")

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	getHistoryDetail, err := h.service.GetHistoryListApprove(encodedIDKM)
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

	encodedIDKM := c.Params("id")

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	getHistoryDetail, err := h.service.GetHistoryListApproveReject(encodedIDKM)
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

	encodedIDKM := c.Params("id")

	getHistoryDetail, err := h.service.GetHistoryListRequested(encodedIDKM)
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

func (h *HistoryHandler) GetContentDetail(c *fiber.Ctx) error {
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	users := permission["user"].(map[string]interface{})
	userId := users["id"].(string)

	content, err := h.service.GetContentDetail(c, userId)
	if err != nil {
		h.log.Printf("Execute GetContentDetail Function: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to retrieve content detail",
			Error:      err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Success Retrieve Content Detail",
		Data:       content,
	})
}
