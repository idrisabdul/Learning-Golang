package updaterequest

import (
	"strconv"
	"sygap_new_knowledge_management/backend/entities"
	updaterequest "sygap_new_knowledge_management/backend/services/update-request"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UpdateRequestHandler struct {
	submit *updaterequest.SubmitService
	update *updaterequest.UpdateService
	detail *updaterequest.DetailService
	log    *logrus.Logger
}

func NewUpdateRequestHandler(submit *updaterequest.SubmitService, update *updaterequest.UpdateService, detail *updaterequest.DetailService, log *logrus.Logger) *UpdateRequestHandler {
	return &UpdateRequestHandler{submit, update, detail, log}
}

// Get Function
func (h *UpdateRequestHandler) GetListUpdateRequest(c *fiber.Ctx) error {

	h.log.Print("Execute GetListUpdateRequest Function")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	idKM := c.Params("id")

	listUpdateRequest, errListUpdateRequest := h.detail.GetListUpdateRequest(idKM)
	if errListUpdateRequest != nil {
		h.log.Printf("Execute ListUpdateRequest Function: %v", errListUpdateRequest.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to retrieve data",
			Error:      errListUpdateRequest.Error(),
		})
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("per_page", "10"))

	data, metaPagination := utils.GetPaginated(c, page, limit, listUpdateRequest)

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieve data",
		Data:       data,
		Meta:       utils.Meta{Pagination: metaPagination},
	})
}

func (h *UpdateRequestHandler) GetDetailUpdateRequest(c *fiber.Ctx) error {
	h.log.Print("Execute GetDetailUpdateRequest Function")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	idUpdateRequest := c.Params("id")

	detailUpdateRequest, errDetailUpdateRequest := h.detail.GetDetailUpdateRequest(idUpdateRequest)
	if errDetailUpdateRequest != nil {
		h.log.Printf("Failed Execute DetailUpdateRequest Function: %v", errDetailUpdateRequest.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to retrieve data",
			Error:      errDetailUpdateRequest.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieve data",
		Data:       detailUpdateRequest,
	})
}

// Submit Function
func (h *UpdateRequestHandler) SubmitUpdateRequest(c *fiber.Ctx) error {

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	author, _ := strconv.Atoi(permission["user"].(map[string]any)["id"].(string))

	validate := validator.New()
	form, errForm := c.MultipartForm()
	if errForm != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to Submit",
			Error:      errForm.Error(),
		})
	}
	var payload entities.SubmitUpdateRequest
	if errParseToPayload := c.BodyParser(&payload); errParseToPayload != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to Submit",
			Error:      errParseToPayload.Error(),
		})
	}

	if err := validate.Struct(payload); err != nil {
		var slice []string

		if _, ok := err.(*validator.InvalidValidationError); ok {
			return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseValidator{
				StatusCode: fiber.StatusBadRequest,
				Message:    "Invalid validation error",
				Error:      []string{err.Error()},
			})
		}

		for _, err := range err.(validator.ValidationErrors) {
			slice = append(slice, err.Field())
		}
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseValidator{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Fill The Required Fields",
			Error:      slice,
		})
	}
	payload.Attachment = form.File["attachment"]
	payload.Submitter = author

	if errSubmitUpdateRequest := h.submit.SubmitUpdateRequest(payload); errSubmitUpdateRequest != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to Submit",
			Error:      errSubmitUpdateRequest.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully Submit data",
	})
}

// Update/Delete Function
func (h *UpdateRequestHandler) UpdateUpdateRequest(c *fiber.Ctx) error {
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	author, _ := strconv.Atoi(permission["user"].(map[string]any)["id"].(string))

	validate := validator.New()
	form, errForm := c.MultipartForm()
	if errForm != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to Update",
			Error:      errForm.Error(),
		})
	}
	var payload entities.SubmitUpdateRequest
	if errParseToPayload := c.BodyParser(&payload); errParseToPayload != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to Update",
			Error:      errParseToPayload.Error(),
		})
	}

	if err := validate.Struct(payload); err != nil {
		var slice []string

		if _, ok := err.(*validator.InvalidValidationError); ok {
			return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseValidator{
				StatusCode: fiber.StatusBadRequest,
				Message:    "Invalid validation error",
				Error:      []string{err.Error()},
			})
		}

		for _, err := range err.(validator.ValidationErrors) {
			slice = append(slice, err.Field())
		}
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseValidator{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Fill The Required Fields",
			Error:      slice,
		})
	}

	payload.ID, _ = c.ParamsInt("id")
	payload.Attachment = form.File["attachment"]
	payload.Submitter = author

	if errUpdateUpdateRequest := h.update.UpdateUpdateRequest(payload); errUpdateUpdateRequest != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to Update",
			Error:      errUpdateUpdateRequest.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully Update data",
	})
}

func (h *UpdateRequestHandler) DeleteUpdateRequest(c *fiber.Ctx) error {
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	author, _ := strconv.Atoi(permission["user"].(map[string]any)["id"].(string))
	id, _ := c.ParamsInt("id")
	if errDeleteUpdateRequest := h.update.DeleteUpdateRequest(id, author); errDeleteUpdateRequest != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to Update",
			Error:      errDeleteUpdateRequest.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully Delete data",
	})
}

func (h *UpdateRequestHandler) DeleteUpdateRequestAttachment(c *fiber.Ctx) error {

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	author, _ := permission["user"].(map[string]any)["id"].(string)
	idFile, _ := c.ParamsInt("id")

	if errDeleteUpdateRequestAttachment := h.update.DeleteUpdateRequestAttachment(idFile, author); errDeleteUpdateRequestAttachment != nil {
		h.log.Printf("Execute DeleteUpdateRequsetAttachment Function: %v", errDeleteUpdateRequestAttachment.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to retrieve KM List",
			Error:      errDeleteUpdateRequestAttachment.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully delete data",
	})
}
