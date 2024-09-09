package knowledge_relation_hdl

import (
	"strconv"
	"sygap_new_knowledge_management/backend/model"
	services "sygap_new_knowledge_management/backend/services/knowledge_relation"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type KnowledgeRelationHandler struct {
	service *services.KnowledgeRelationService
	logger  *logrus.Logger
}

func NewKnowledgeRelationHandler(service *services.KnowledgeRelationService, logger *logrus.Logger) *KnowledgeRelationHandler {
	return &KnowledgeRelationHandler{
		service: service,
		logger:  logger,
	}
}

func (p *KnowledgeRelationHandler) SearchKnowledge(c *fiber.Ctx) error {
	// validation
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	var search_param model.KnowledgeRelationSearch
	if err := c.BodyParser(&search_param); err != nil {
		p.logger.WithError(err).Error("failed to parse request body")
		p.logger.Println("Error parsing JSON:", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:     "Failed to parse request body",
			Error:       err.Error(),
		})
	}

	knowledges, err := p.service.SearchKnowledge(c, search_param)
	if err != nil {
		p.logger.Printf("Failed to retrieve data: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:     "Failed to retrieve ci type",
			Error:       err.Error(),
		})
	}

	// create page and limit for pagination
	Page, _ := strconv.Atoi(c.Query("page", "1"))
	Limit, _ := strconv.Atoi(c.Query("limit", "10"))
	Data, MetaPagination := utils.GetPaginated(c, Page, Limit, knowledges)
	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:     "Success",
		Data:        Data,
		Meta:        utils.Meta{Pagination: MetaPagination},
	})
}

func (p *KnowledgeRelationHandler) SubmitKnowledge(c *fiber.Ctx) error {
	// validation
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	// validator format
	var payload_submit model.KnowledgeRelationToTicketSubmit
	validate := validator.New()
	if err := c.BodyParser(&payload_submit); err != nil {
		p.logger.WithError(err).Error("failed to parse request body")
		p.logger.Println("Error parsing JSON:", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:     "Failed to parse request body",
			Error:       err.Error(),
		})
	}
	if err := validate.Struct(&payload_submit); err != nil {
		var slice []string
		for _, err := range err.(validator.ValidationErrors) {
			p.logger.WithError(err).Error("Validation error:")
			p.logger.Println("Validation error:", err.Field(), err.Tag(), err.Param())
			slice = append(slice, err.Field())
		}
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseValidator{
			StatusCode: fiber.StatusBadRequest,
			Message:     "Fill The Required Fields",
			Error:       slice,
		})
	}

	users := permission["user"].(map[string]interface{})
	id := users["id"].(string)
	add_work_detail, err := p.service.SubmitKnowledge(payload_submit, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:     "Failed to parse request body",
			Error:       err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:     "Knowledge relation success to created",
		Data:        add_work_detail,
	})
}

func (p *KnowledgeRelationHandler) ListKnowledge(c *fiber.Ctx) error {
	// validation
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	knowledges, err := p.service.ListKnowledge(c)
	if err != nil {
		p.logger.Printf("Failed to retrieve data: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:     "Failed to retrieve knowledge relation",
			Error:       err.Error(),
		})
	}

	// create page and limit for pagination
	Page, _ := strconv.Atoi(c.Query("page", "1"))
	Limit, _ := strconv.Atoi(c.Query("limit", "10"))
	Data, MetaPagination := utils.GetPaginated(c, Page, Limit, knowledges)
	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:     "Success",
		Data:        Data,
		Meta:        utils.Meta{Pagination: MetaPagination},
	})
}

func (p *KnowledgeRelationHandler) ExportKnowledge(c *fiber.Ctx) error {
	// validation
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	knowledges, err := p.service.ExportKnowledge(c)
	if err != nil {
		p.logger.Printf("Failed to retrieve data: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:     "Failed to retrieve ci type",
			Error:       err.Error(),
		})
	}

	// get header
	getHeader, _ := p.service.GetHeaderTittleExcelSvc(c)
	file_excell := utils.ExportExcellCustom(knowledges, getHeader)
	c.Set(fiber.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set(fiber.HeaderContentDisposition, "attachment; filename=km-knowledge_relation.xls")
	return c.Send(file_excell)
}

func (p *KnowledgeRelationHandler) DeleteKnowledge(c *fiber.Ctx) error {
	// validation
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	// validator format
	var payload_delete model.KnowledgeRelationToTicketDelete
	if err := c.BodyParser(&payload_delete); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:     "Failed to parse request body",
			Error:       err.Error(),
		})
	}

	users := permission["user"].(map[string]interface{})
	id := users["id"].(string)
	add_work_detail, err := p.service.DeleteKnowledge(payload_delete, id, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:     "Failed to parse request body",
			Error:       err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:     "Knowledge relation success to deleted",
		Data:        add_work_detail,
	})
}
