package relation_hdl

import (
	"io"
	"strconv"
	"strings"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/model"
	services "sygap_new_knowledge_management/backend/services/relation"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
)

type RelationHdlr struct {
	service *services.RelationSvc
	logger  *logrus.Logger
}

func NewRelationHdlr(service *services.RelationSvc, logger *logrus.Logger) *RelationHdlr {
	return &RelationHdlr{service, logger}
}

// SearchRelationList handles POST requests to search Problem Relation List.
func (h *RelationHdlr) SearchRelationList(c *fiber.Ctx) error {
	// Validate Auth Token
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	// Parse payload data
	var payload entities.SearchRelationParams

	if err := c.BodyParser(&payload); err != nil {
		h.logger.Println("Failed to parse payload data:", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
			StatusCode: fiber.StatusBadRequest,
			Message:     "Error parsing payload data",
			Error:       err.Error(),
		})
	}

	// Retrieve related relation list from the service
	problems, priority, err := h.service.GetRelationList(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:     "Failed to retrieve data",
			Error:       err.Error(),
		})
	}

	if payload.Page == "" {
		payload.Page = "1"
	}

	if payload.Limit == "" {
		payload.Limit = "10"
	}

	// Pagination
	page, err := strconv.Atoi(payload.Page)
	if err != nil || page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:     "Invalid page number",
		})
	}

	limit, err := strconv.Atoi(payload.Limit)
	if err != nil || limit < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:     "Invalid limit value",
		})
	}

	data, metaPagination := utils.GetPaginated(c, page, limit, problems)
	dataMap := map[string]interface{}{
		"relation": data,
		"priority": priority,
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:     "Successfully retrieved data",
		Data:        dataMap,
		Meta:        utils.Meta{Pagination: metaPagination},
	})
}

// SearchReqTypeRelationList handles POST requests to search Problem Relation List.
func (h *RelationHdlr) SearchReqTypeRelationList(c *fiber.Ctx) error {
	// Validate Auth Token
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	// Parse payload data
	var payload entities.SearchReqTypeRelationParams
	if err := c.BodyParser(&payload); err != nil {
		h.logger.Println("Failed to parse payload data:", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
			StatusCode: fiber.StatusBadRequest,
			Message:     "Error parsing payload data",
			Error:       err.Error(),
		})
	}

	var codeList []string
	fileHeaders, err := c.FormFile("upload_file")
	if err == nil {

		// Open the file
		file, err := fileHeaders.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		// Read the file content into a byte slice
		fileContent, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		xlFile, err := xlsx.OpenBinary(fileContent)
		if err != nil {
			h.logger.Fatal(err)
		}

		// Iterate over sheets
		for _, sheet := range xlFile.Sheets {
			// Iterate over rows
			for _, row := range sheet.Rows { 
				// Iterate over cells in the row
				for _, cell := range row.Cells {
					// Check if the cell string contains "INC" or "PBI" or "CRQ" or "REQ"
					cellValue := cell.String()
					if strings.Contains(cellValue, "KMS") || strings.Contains(cellValue, "TAS") || strings.Contains(cellValue, "PKE") || strings.Contains(cellValue, "INC") || strings.Contains(cellValue, "PBI") || strings.Contains(cellValue, "CRQ") || strings.Contains(cellValue, "REQ") {
						// fmt.Printf("Found keyword in cell: %s\n", cellValue)
						codeList = append(codeList, cellValue)
					}
				}
			}
		}
	}

	// Retrieve related relation list from the service
	problems, err := h.service.GetReqTypeRelationList(payload, permission, codeList)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:     "Failed to retrieve data",
			Error:       err.Error(),
		})
	}

	// Pagination
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:     "Invalid page number",
		})
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:     "Invalid limit value",
		})
	}

	data, metaPagination := utils.GetPaginated(c, page, limit, problems)

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:     "Successfully retrieved data",
		Data:        data,
		Meta:        utils.Meta{Pagination: metaPagination},
	})
}

func (h *RelationHdlr) InsertRelation(c *fiber.Ctx) error {
	// Validate Auth Token
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	//parse payload data
	var payload model.InsertRelationParams
	if err := c.BodyParser(&payload); err != nil {
		h.logger.Println("Failed to parse payload data:", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
			StatusCode: fiber.StatusBadRequest,
			Message:     "Error parsing payload data",
			Error:       err.Error(),
		})
	}

	// Insert relation data to the service
	err := h.service.InsertRelation(payload, permission)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:     "Failed to insert relation",
			Error:       err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:     "Successfully insert relation",
	})
}

func (h *RelationHdlr) DeleteRelations(c *fiber.Ctx) error {
	// Validate Auth Token
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	// Parse payload data
	var requestBody struct {
		ID string `json:"id"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		h.logger.Println("Failed to parse payload data:", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
			StatusCode: fiber.StatusBadRequest,
			Message:     "Error parsing payload data",
			Error:       err.Error(),
		})
	}

	// Delete relations data from the service
	err := h.service.DeleteRelations(requestBody.ID, permission)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:     "Failed to delete relation",
			Error:       err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:     "Successfully delete relation",
	})
}

// SearchRelationList handles POST requests to search Problem Relation List.
func (h *RelationHdlr) ExportRelationList(c *fiber.Ctx) error {

	// Validate Auth Token
	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:     "The token you entered is invalid",
			Error:       "Authorization failed",
		})
	}

	//parse payload data
	var payload model.ExportRelationParams
	if err := c.BodyParser(&payload); err != nil {
		h.logger.Println("Failed parsed payload data")
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseData{
			StatusCode: fiber.StatusBadRequest,
			Message:     "Error parsed payload data",
			Error:       err.Error(),
		})
	}

	parsePayload := entities.SearchRelationParams{ // parse due to get list using query struct model and for export using json
		IDEntityType: payload.IDEntityType,
	}
	// retrieves related incident from the service
	problems, err := h.service.ExportRelationList(parsePayload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:     "Failed to retrieve data",
			Error:       err.Error(),
		})

	}

	// get header excel
	getHeader, _ := h.service.GetHeaderTittleExcelSvc(payload.IDEntityType)
	file_excell := utils.ExportExcellCustom(problems, getHeader)
	c.Set(fiber.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set(fiber.HeaderContentDisposition, "attachment; filename=km-relation.xls")
	return c.Send(file_excell)
}
