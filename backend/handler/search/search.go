package search

import (
	"fmt"
	"strconv"
	"sygap_new_knowledge_management/backend/model"
	"sygap_new_knowledge_management/backend/services/search"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SearchHandler struct {
	service *search.SearchService
	log     *logrus.Logger
}

func NewSearchHandler(service *search.SearchService, log *logrus.Logger) *SearchHandler {
	return &SearchHandler{service, log}
}

// Api search list page
func (h *SearchHandler) GetSearchList(c *fiber.Ctx) error {
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	// create page and limit for pagination
	Page, _ := strconv.Atoi(c.Query("page", "1"))
	Limit, _ := strconv.Atoi(c.Query("limit", "10"))

	users := permission["user"].(map[string]interface{})
	user_id := users["id"].(string)

	search_list, err := h.service.GetSearchList(c, user_id)
	if err != nil {
		h.log.Printf("Execute get Search List Function: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to retrieve search List",
			Error:      err.Error(),
		})
	}

	Data, MetaPagination := utils.GetPaginated(c, Page, Limit, search_list)

	return c.Status(200).JSON(utils.Response{
		StatusCode: 200,
		Message:    "Succes",
		Data:       Data,
		Meta:       utils.Meta{Pagination: MetaPagination},
	})
}

// Api search content detail
func (h *SearchHandler) GetContentDetail(c *fiber.Ctx) error {
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	users := permission["user"].(map[string]interface{})
	user_id := users["id"].(string)
	content, err := h.service.GetContentDetail(c, user_id)
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

// Api get comment / feedback
func (h *SearchHandler) GetContentFeedback(c *fiber.Ctx) error {
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	users := permission["user"].(map[string]interface{})
	user_id := users["id"].(string)
	content, err := h.service.GetContentFeedback(c, user_id)
	if err != nil {
		h.log.Printf("Execute GetContentFeedback Function: %v", err.Error())
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

// Api report content detail
func (h *SearchHandler) ReportContentDetail(c *fiber.Ctx) error {
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	users := permission["user"].(map[string]interface{})
	user_id := users["id"].(string)
	employee_name := users["employee_name"].(string)
	content, err := h.service.ReportContentDetail(c, user_id, employee_name)
	if err != nil {
		h.log.Printf("Execute ReportContentDetail Function: %v", err.Error())
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

// Api bookmark content detail
func (h *SearchHandler) BookmarkContentDetail(c *fiber.Ctx) error {
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	users := permission["user"].(map[string]interface{})
	user_id := users["id"].(string)
	content, err := h.service.BookmarkContentDetail(c, user_id)
	if err != nil {
		h.log.Printf("Execute BookmarkContentDetail Function: %v", err.Error())
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

// Api bookmark content detail
func (h *SearchHandler) FeedbackContentDetail(c *fiber.Ctx) error {
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	var payload model.FeedbackContent
	validate := validator.New()
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(utils.Response{
			StatusCode: 400,
			Message:    "Failed to parse request body",
			Error:      err.Error(),
		})
	}
	if err := validate.Struct(&payload); err != nil {
		var slice []string
		for _, err := range err.(validator.ValidationErrors) {
			slice = append(slice, err.Field())
		}
		message := fmt.Sprintf("Fill The Required Fields %v", slice)
		return c.Status(400).JSON(utils.ResponseValidator{
			StatusCode: 400,
			Message:    message,
			Error:      slice,
		})
	}

	users := permission["user"].(map[string]interface{})
	user_id := users["id"].(string)
	employee_name := users["employee_name"].(string)
	content, err := h.service.FeedbackContentDetail(c, user_id, payload, employee_name)
	if err != nil {
		h.log.Printf("Execute FeedbackContentDetail Function: %v", err.Error())
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
