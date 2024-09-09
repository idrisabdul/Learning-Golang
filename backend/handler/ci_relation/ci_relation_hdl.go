package ci_relation_hdl

import (
	"strconv"
	"sygap_new_knowledge_management/backend/model"
	services "sygap_new_knowledge_management/backend/services/ci_relation"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CiRelationHandler struct {
	service *services.CiRelationService
	logger  *logrus.Logger
}

func NewCiRelationHandler(service *services.CiRelationService, logger *logrus.Logger) *CiRelationHandler {
	return &CiRelationHandler{
		service: service,
		logger:  logger,
	}
}

func (p *CiRelationHandler) GetCiType(c *fiber.Ctx) error {
	// validation
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	ci_type, err := p.service.GetCiType()
	if err != nil {
		p.logger.Printf("Failed to retrieve data: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to retrieve ci type",
			Error:      err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieved ci type",
		Data:       ci_type,
	})
}

func (p *CiRelationHandler) GetRelationType(c *fiber.Ctx) error {
	// validation
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	ci_relation_type, err := p.service.GetCiRelationType()
	if err != nil {
		p.logger.Printf("Failed to retrieve data: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to retrieve  ci relation type",
			Error:      err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieved ci relation type",
		Data:       ci_relation_type,
	})
}

func (p *CiRelationHandler) GetCiName(c *fiber.Ctx) error {
	// validation
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	ci_relation_type, err := p.service.GetCiName(c)
	if err != nil {
		p.logger.Printf("Failed to retrieve data: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to retrieve  ci relation type",
			Error:      err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieved ci relation type",
		Data:       ci_relation_type,
	})
}

func (h *CiRelationHandler) CreateCiRelation(c *fiber.Ctx) error {
	// validation
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	// validator format
	var payload_ci_relation model.CiRelationSubmit
	validate := validator.New()
	if err := c.BodyParser(&payload_ci_relation); err != nil {
		h.logger.WithError(err).Error("failed to parse request body")
		h.logger.Println("Error parsing JSON:", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to parse request body",
			Error:      err.Error(),
		})
	}
	if err := validate.Struct(&payload_ci_relation); err != nil {
		var slice []string
		for _, err := range err.(validator.ValidationErrors) {
			h.logger.WithError(err).Error("Validation error:")
			h.logger.Println("Validation error:", err.Field(), err.Tag(), err.Param())
			slice = append(slice, err.Field())
		}
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseValidator{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Fill The Required Fields",
			Error:      slice,
		})
	}

	users := permission["user"].(map[string]interface{})
	id := users["id"].(string)
	add_work_detail, err := h.service.SubmitCiRelation(payload_ci_relation, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to parse request body",
			Error:      err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Ci relation success to created",
		Data:       add_work_detail,
	})
}

func (p *CiRelationHandler) GetCiRelationList(c *fiber.Ctx) error {
	// validation
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

	ci_relation_list, err := p.service.GetCiRelationList(c)
	if err != nil {
		p.logger.Printf("Failed to retrieve data: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to retrieve  ci relation list",
			Error:      err.Error(),
		})
	}

	Data, MetaPagination := utils.GetPaginated(c, Page, Limit, ci_relation_list)

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Success",
		Data:       Data,
		Meta:       utils.Meta{Pagination: MetaPagination},
	})
}

func (h *CiRelationHandler) UpdateCiRelation(c *fiber.Ctx) error {
	// validation
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	// validator format
	var payload_ci_relation model.CiRelationUpdate
	validate := validator.New()
	if err := c.BodyParser(&payload_ci_relation); err != nil {
		h.logger.WithError(err).Error("failed to parse request body")
		h.logger.Println("Error parsing JSON:", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to parse request body",
			Error:      err.Error(),
		})
	}
	if err := validate.Struct(&payload_ci_relation); err != nil {
		var slice []string
		for _, err := range err.(validator.ValidationErrors) {
			h.logger.WithError(err).Error("Validation error:")
			h.logger.Println("Validation error:", err.Field(), err.Tag(), err.Param())
			slice = append(slice, err.Field())
		}
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseValidator{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Fill The Required Fields",
			Error:      slice,
		})
	}

	users := permission["user"].(map[string]interface{})
	id := users["id"].(string)
	add_work_detail, err := h.service.UpdateCiRelation(c, payload_ci_relation, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to parse request body",
			Error:      err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Ci relation success to updated",
		Data:       add_work_detail,
	})
}

func (p *CiRelationHandler) DeleteCiRelation(c *fiber.Ctx) error {
	// validation
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	users := permission["user"].(map[string]interface{})
	id := users["id"].(string)
	ci_relation_type, err := p.service.DeleteCiRelation(c, id)
	if err != nil {
		p.logger.Printf("Failed to retrieve data: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to delete ci relation",
			Error:      err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully delete ci relation",
		Data:       ci_relation_type,
	})
}

func (p *CiRelationHandler) GetCiHistoryHdl(c *fiber.Ctx) error {
	// validation
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(401).JSON(utils.ResponseAuth{
			StatusCode: 401,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	ci_history, err := p.service.GetDataCiHistorySvc(c)
	if err != nil {
		p.logger.Printf("Failed to retrieve data: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to retrieve ci history",
			Error:      err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieved ci history",
		Data:       ci_history,
	})
}

func (p *CiRelationHandler) GetAttributeNameHdl(c *fiber.Ctx) error {
	// validation
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(401).JSON(utils.ResponseAuth{
			StatusCode: 401,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	ci_attribute, err := p.service.GetAttributeNameOptionSvc(c)
	if err != nil {
		p.logger.Printf("Failed to retrieve data: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to retrieve ci attribute",
			Error:      err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully retrieved ci attribute",
		Data:       ci_attribute,
	})
}
