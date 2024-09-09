package workdetail

import (
	"fmt"
	"strconv"
	"sygap_new_knowledge_management/backend/entities"
	services "sygap_new_knowledge_management/backend/services/work-detail"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type WorkDetailHandler struct {
	service *services.WorkDetailService
	log     *logrus.Logger
}

func NewWorkDetailHandler(service *services.WorkDetailService, log *logrus.Logger) *WorkDetailHandler {
	return &WorkDetailHandler{
		service: service,
		log:     log,
	}
}

func (h *WorkDetailHandler) GetWorkDetail(c *fiber.Ctx) error {
	h.log.Print("Execute GetWorkDetail Function")

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

	getWorkDetail, err := h.service.GetWorkDetail(idKM)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to parse request body",
			Error:      err.Error(),
		})
	}

	data, metaPagination := utils.GetPaginated(c, page, limit, getWorkDetail)

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieve data",
		Data:       data,
		Meta:       utils.Meta{Pagination: metaPagination},
	})
}

func (h *WorkDetailHandler) AddWorkDetail(c *fiber.Ctx) error {
	h.log.Print("Execute GetWorkDetail Function")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	var payload entities.SubmitWorkDetail
	validate := validator.New()
	if err := c.BodyParser(&payload); err != nil {
		h.log.WithError(err).Error("failed to parse request body")
		h.log.Println("Error parsing JSON:", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to parse request body",
			Error:      err.Error(),
		})
	}

	if err := validate.Struct(&payload); err != nil {
		var slice []string
		for _, err := range err.(validator.ValidationErrors) {
			h.log.WithError(err).Error("Validation error:")
			h.log.Println("Validation error:", err.Field(), err.Tag(), err.Param())
			slice = append(slice, err.Field())
		}
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseValidator{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Fill The Required Fields",
			Error:      slice,
		})
	}

	payload.Submitter, _ = strconv.Atoi(permission["user"].(map[string]interface{})["id"].(string))
	payload.IDParent, _ = c.ParamsInt("id")

	fmt.Printf("payload.IDParent: %v\n", payload.IDParent)
	form, _ := c.MultipartForm()
	payload.Attachment = form.File["FileUpload"]

	fmt.Printf("payload.Attachment: %v\n", payload.Attachment)
	if errSubmitWorkdetail := h.service.SubmitWorkDetail(payload); errSubmitWorkdetail != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to add work detail",
			Error:      errSubmitWorkdetail.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully add work detail",
	})
}
