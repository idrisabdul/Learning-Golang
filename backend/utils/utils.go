package utils

import (
	"errors"
	"net/http"
	"sygap_new_knowledge_management/backend/pkg/errs"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ResponseWithError(c *fiber.Ctx, err error) error {
	var badRequestError *errs.BadRequestError
	var resourceNotFoundError *errs.ResourceNotFoundError
	var forbiddenError *errs.ForbiddenError

	if errors.As(err, &badRequestError) {
		return c.Status(http.StatusBadRequest).JSON(ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Error:      "BAD_REQUEST",
		})
	} else if errors.As(err, &resourceNotFoundError) {
		return c.Status(http.StatusNotFound).JSON(ResponseError{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Error:      "NOT_FOUND",
		})
	} else if errors.As(err, &forbiddenError) {
		return c.Status(http.StatusForbidden).JSON(ResponseError{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Error:      "FORBIDDEN",
		})
	} else {
		return c.Status(http.StatusInternalServerError).JSON(ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Error:      "INTERNAL_SERVER_ERROR",
		})
	}
}

func ResponseRequestValidationError(c *fiber.Ctx, err error) error {
	var slice []string

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseValidator{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid validation error",
			Error:      []string{err.Error()},
		})
	}

	for _, err := range err.(validator.ValidationErrors) {
		slice = append(slice, err.Field())
	}
	return c.Status(fiber.StatusBadRequest).JSON(ResponseValidator{
		StatusCode: fiber.StatusBadRequest,
		Message:    "Fill The Required Fields",
		Error:      slice,
	})
}
