package feedback_hdl

import (
	"strconv"
	services "sygap_new_knowledge_management/backend/services/feedback"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type FeedbackHandler struct {
	service *services.FeedbackService
	log  *logrus.Logger
}

func NewFeedbackHandler(service *services.FeedbackService, log *logrus.Logger) *FeedbackHandler {
	return &FeedbackHandler{
		service: service,
		log:  log,
	}
}

func (h *FeedbackHandler) GetFeedbackList(c *fiber.Ctx) error {
	h.log.Print("Execute GetFeedbackList Function")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	getFeedback, err := h.service.GetFeedbackList(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
			StatusCode: fiber.StatusBadRequest,
			Message:     "Failed to parse request body",
			Error:       err.Error(),
		})
	}

	data, metaPagination := utils.GetPaginated(c, page, limit, getFeedback)

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieve data",
		Data:       data,
		Meta:       utils.Meta{Pagination: metaPagination},
	})
}

// SearchRelationList handles POST requests to search Problem Relation List.
func (h *FeedbackHandler) ExportFeedbackList(c *fiber.Ctx) error {

	// Validate Auth Token
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	//parse payload data
	// var payload model.ExportFeedbackParams
	// if err := c.BodyParser(&payload); err != nil {
	// 	h.log.Println("Failed parsed payload data")
	// 	return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
	// 		StatusCode: fiber.StatusBadRequest,
	// 		Message:     "Error parsed payload data",
	// 		Error:       err.Error(),
	// 	})
	// }

	// retrieves related incident from the service
	feedbacks, err := h.service.ExportFeedbackList(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:     "Failed to retrieve data",
			Error:       err.Error(),
		})

	}
	
	// get header excel
	getHeader, _ := h.service.GetHeaderTitleExcelSvc(c)
	file_excell := utils.ExportExcellCustom(feedbacks, getHeader)
	c.Set(fiber.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set(fiber.HeaderContentDisposition, "attachment; filename=km-feedback.xls")
	return c.Send(file_excell)
}
