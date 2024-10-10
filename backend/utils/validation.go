package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	// "github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

// var validate *validator.Validate = validator.New()

func ValidateToken(c *fiber.Ctx) map[string]interface{} {
	url := os.Getenv("URL_CI")
	authorize := fetchGetApi(url, c)

	return authorize
}

func fetchGetApi(url string, c *fiber.Ctx) map[string]interface{} {
	response, err := http.NewRequest("GET", url, nil)
	body := map[string]interface{}{
		"status": "200",
		"msg":    "Payload not found or Authorization Failed",
	}

	if err != nil {
		body["error"] = err
		body["status"] = "400"
		return body
	}

	response.Header.Add("Authorization", c.Get("Authorization"))
	responseData, err := http.DefaultClient.Do(response)
	if err != nil {
		body["error"] = err
		body["status"] = "500"
		return body
	}
	defer responseData.Body.Close()
	data, err := io.ReadAll(responseData.Body)
	if err != nil {
		body["error"] = err
		body["status"] = "400"
		return body
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(string(data)), &result)

	status := strconv.FormatFloat(result["status"].(float64), 'f', -1, 64)
	result["status"] = status
	return result
}

func ParseAndValidate[T any](c *fiber.Ctx, data []byte) (T, *Response) {
	// Create a new validator instance
	validate := validator.New()

	var request T

	if errParse := json.Unmarshal(data, &request); errParse != nil {
		return request, &Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid data sent. Please Check your data and try again",
			Error:      "Failed to parse data to struct",
		}
	}

	// Validate the request struct
	if errValidate := validate.Struct(request); errValidate != nil {
		// Handle validation errors
		var errors []string
		if _, ok := errValidate.(*validator.InvalidValidationError); ok {
			return request, &Response{
				StatusCode: fiber.StatusBadRequest,
				Message:    "Invalid data sent. Please check your data and try again",
				Error:      "Validation Error: " + errValidate.Error(),
			}
		}

		for _, err := range errValidate.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}

		return request, &Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Fill the required fields",
			Error:      strings.Join(errors, ", "),
		}
	}

	return request, nil
}

// not ready to use. DONT USE THIS
// func ValidateRequest(data FlexibleType, c *fiber.Ctx) error {
// 	if err := validate.Struct(&data); err != nil {
// 		var slice []string

// 		if _, ok := err.(*validator.InvalidValidationError); ok {
// 			return c.Status(fiber.StatusBadRequest).JSON(ResponseValidator{
// 				StatusCode: fiber.StatusBadRequest,
// 				Message:    "Invalid validation error",
// 				Error:      []string{err.Error()},
// 			})
// 		}

// 		for _, err := range err.(validator.ValidationErrors) {
// 			slice = append(slice, err.Field())
// 		}
// 		return c.Status(fiber.StatusBadRequest).JSON(ResponseValidator{
// 			StatusCode: fiber.StatusBadRequest,
// 			Message:    "Fill The Required Fields",
// 			Error:      slice,
// 		})
// 	}

// 	return nil
// }
