package options

import (
	"sygap_new_knowledge_management/backend/services/options"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	service *options.ProductService
	log     *logrus.Logger
}

func NewProductHandler(service *options.ProductService, log *logrus.Logger) *ProductHandler {
	return &ProductHandler{service, log}
}

func (h *ProductHandler) GetListProduct(c *fiber.Ctx) error {
	h.log.Print("Execute GetListProduct Function in Handler")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	idCompany := c.Query("id_company")
	search := c.Query("search")
	id_category := c.Query("id_category")

	var response any

	if id_category != "" {

		productParentName, errProductParentName := h.service.GetProductParentCategory(id_category)

		if errProductParentName != nil {
			h.log.Printf("Failed to retrieve data: %v", errProductParentName.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "Failed to retrieve product List",
				Error:      errProductParentName.Error(),
			})
		}

		response = productParentName
	} else {

		listProduct, errListProduct := h.service.GetListProduct(idCompany, search)

		if errListProduct != nil {
			h.log.Printf("Failed to retrieve data: %v", errListProduct.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "Failed to retrieve product parent name",
				Error:      errListProduct.Error(),
			})
		}

		response = listProduct
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Success Retrieve Product List",
		Data:       response,
	})
}
