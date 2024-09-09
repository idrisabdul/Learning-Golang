package options

import (
	options "sygap_new_knowledge_management/backend/services/options"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProductTypeHandler struct {
	service *options.ProductTypeService
	log     *logrus.Logger
}

func NewProductTypeHandler(service *options.ProductTypeService, log *logrus.Logger) *ProductTypeHandler {
	return &ProductTypeHandler{service, log}
}

func (h *ProductTypeHandler) GetProductTypeHandler(c *fiber.Ctx) error {
	h.log.Println("Executtye function GetProductTypeHandler")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	IDProductName := c.Query("id_product_name")

	listProductType, err := h.service.GetListProductType(IDProductName, c.Query("is_all")) // is all for product in task
	if err != nil {
		h.log.Error("Failed get list product type in GetProductTypeHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:     "Failed get list product type",
			Error:       err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:     "Success",
		Data:        listProductType,
	})

}

func (h *ProductTypeHandler) GetOptionProductTypeRelationHdlr(c *fiber.Ctx) error {
	h.log.Println("Executtye function GetOptionProductTypeRelationHdlr")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	reqType := c.Query("request_type")
	search := c.Query("search")
	isAll := c.Query("is_all")

	listProductType, err := h.service.GetListProductTypeRelation(reqType, search, isAll)
	if err != nil {
		h.log.Error("Failed get list product type in GetListProductTypeRelation")

		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:     "Failed get list product type",
			Error:       err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:     "Success",
		Data:        listProductType,
	})

}
