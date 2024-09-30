package km

import (
	"strconv"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/services/km/form"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type FormHandler struct {
	submit *form.SubmitService
	detail *form.DetailService
	update *form.UpdateService
	log    *logrus.Logger
}

func NewFormhandler(submit *form.SubmitService, detail *form.DetailService, update *form.UpdateService, log *logrus.Logger) *FormHandler {
	return &FormHandler{submit, detail, update, log}
}

func (h *FormHandler) SubmitKM(c *fiber.Ctx) error {
	h.log.Print("Execute SubmitKM function on Handler")

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
	isDecisionTree := c.QueryBool("isdt", false)
	toDraft := c.QueryBool("tdrft", false)
	var response string

	if isDecisionTree {
		var request entities.RequestSubmitKMDecisionTree
		// parse request from Front End
		if errParseToPayload := c.BodyParser(&request); errParseToPayload != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
				StatusCode: fiber.StatusBadRequest,
				Message:    "Failed Submit data",
				Error:      errParseToPayload.Error(),
			})
		}

		// validate payload
		if toDraft {
			if err := validate.Struct(request); err != nil {
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

		}
		// to submit
		IDKM, errSubmit := h.submit.SubmitKMDecisionTree(request, author, toDraft)
		if errSubmit != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "Failed Submit data",
				Error:      errSubmit.Error(),
			})
		}
		response = IDKM
	} else {
		var request entities.SubmitKMNonDecisionTree

		// parse request from Front End
		if errParseToPayload := c.BodyParser(&request); errParseToPayload != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
				StatusCode: fiber.StatusBadRequest,
				Message:    "Failed Submit data",
				Error:      errParseToPayload.Error(),
			})
		}

		// validate payload
		if toDraft {
			if err := validate.Struct(request); err != nil {
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

		}
		// to submit
		IDKM, errSubmit := h.submit.SubmitKM(request, author, toDraft)
		if errSubmit != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "Failed Submit data",
				Error:      errSubmit.Error(),
			})
		}

		response = IDKM
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully Submit data",
		Data:       response,
	})
}

func (h *FormHandler) DetailKM(c *fiber.Ctx) error {
	h.log.Print("Execute DetailKM function on Handler")

	var response any
	var err error

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	idKM := c.Params("id")
	isDecisionTree := c.QueryBool("isdt", false)
	if isDecisionTree {
		detailKMDecisionTree, errDetailKMDecisionTree := h.detail.DetailKMDecisionTree(idKM, permission)
		response = detailKMDecisionTree
		err = errDetailKMDecisionTree
	} else {
		detailKm, errDetailKM := h.detail.DetailKM(idKM, permission)
		response = detailKm
		err = errDetailKM
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed Retrieve data",
			Data:       err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully Retrieve data",
		Data:       response,
	})
}

func (h *FormHandler) UpdateKM(c *fiber.Ctx) error {
	h.log.Print("Execute UpdateKM function on Handler")

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

	isDecisionTree := c.QueryBool("isdt", false)
	step := c.Params("step", "none")

	if isDecisionTree {
		var request entities.SubmitKMDecisionTree
		// parse request from Front End
		if errParseToPayload := c.BodyParser(&request); errParseToPayload != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
				StatusCode: fiber.StatusBadRequest,
				Message:    "Failed Submit data",
				Error:      errParseToPayload.Error(),
			})
		}

		// validate payload
		// if errValidateRequest := utils.ValidateRequest(request, c); errValidateRequest != nil {
		// 	return errValidateRequest
		// }
		if err := validate.Struct(request); err != nil {
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

		// to submit
		if errUpdate := h.update.UpdateKMDecisionTree(request, author, step); errUpdate != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "Failed Update data",
				Error:      errUpdate.Error(),
			})
		}
	} else {
		var request entities.SubmitKMNonDecisionTree

		// parse request from Front End
		if errParseToPayload := c.BodyParser(&request); errParseToPayload != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
				StatusCode: fiber.StatusBadRequest,
				Message:    "Failed Submit data",
				Error:      errParseToPayload.Error(),
			})
		}

		// validate payload
		// if errValidateRequest := utils.ValidateRequest(request, c); errValidateRequest != nil {
		// 	return errValidateRequest
		// }
		if err := validate.Struct(request); err != nil {
			var slice []string
			// validate payload
			// if errValidateRequest := utils.ValidateRequest(request, c); errValidateRequest != nil {
			// 	return errValidateRequest
			// }
			if err := validate.Struct(request); err != nil {
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
			for _, err := range err.(validator.ValidationErrors) {
				slice = append(slice, err.Field())
			}
			return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseValidator{
				StatusCode: fiber.StatusBadRequest,
				Message:    "Fill The Required Fields",
				Error:      slice,
			})
		}

		// to submit
		if errUpdate := h.update.UpdateKM(request, author, step); errUpdate != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "Failed Submit data",
				Error:      errUpdate.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully Update data",
	})
}

func (h *FormHandler) SetClosedVersion(c *fiber.Ctx) error {
	h.log.Print("Execute UpdateKM function on Handler")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	author, _ := strconv.Atoi(permission["user"].(map[string]any)["id"].(string))

	var payload entities.SetClosedVersion

	if errParseToPayload := c.BodyParser(&payload); errParseToPayload != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed Set Closed Version data",
			Error:      errParseToPayload.Error(),
		})
	}

	if errSetClosedVersion := h.update.SetClosedVersion(payload.ID, author, payload.Note); errSetClosedVersion != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed Set Closed Version data",
			Error:      errSetClosedVersion.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully Set Closed Version data",
	})
}
