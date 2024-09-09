package km

import (
	"strconv"
	"sygap_new_knowledge_management/backend/services/km/document"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type DocumentHandler struct {
	submit *document.SubmitService
	detail *document.DetailService
	delete *document.DeleteService
	log    *logrus.Logger
}

func NewDocumentHandler(submit *document.SubmitService, detail *document.DetailService, delete *document.DeleteService, log *logrus.Logger) *DocumentHandler {
	return &DocumentHandler{submit, detail, delete, log}
}

func (h *DocumentHandler) SubmitDocumentKM(c *fiber.Ctx) error {
	h.log.Print("Execute SubmitDocumentKM function on Handler")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	idKM := c.Params("id")
	new := c.QueryBool("new")

	var usedID int

	if new {

		DecodedIDKM, errDecodedIDKM := utils.GenerateNumberDecode(idKM)
		if errDecodedIDKM != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseAuth{
				StatusCode: fiber.StatusBadRequest,
				Message:    "The ID you entered is invalid",
				Error:      "Invalid ID",
			})
		}

		usedID = *DecodedIDKM
	} else {
		usedID, _ = strconv.Atoi(idKM)
	}

	form, errMultipartForm := c.MultipartForm()
	if errMultipartForm != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed parsing form-data",
			Error:      errMultipartForm.Error(),
		})
	}

	file := form.File["attachment"]
	IDFiles, errSubmitDocument := h.submit.SubmitDocument(file, usedID)
	if errSubmitDocument != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed submit document",
			Error:      errSubmitDocument.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully Submit data",
		Data:       IDFiles,
	})
}

func (h *DocumentHandler) ListDetailDocumentKM(c *fiber.Ctx) error {
	h.log.Print("Execute ListDetailDocumentKM function on Handler")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	idKM := c.Params("id")

	listDocument, errListDocument := h.detail.DetailDocumentKM(idKM)
	if errListDocument != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed Retrieve data",
			Error:      errListDocument.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully Retrieve data",
		Data:       listDocument,
	})
}

func (h *DocumentHandler) URLDocumentKM(c *fiber.Ctx) error {
	h.log.Print("Execute URLDocumentKM function on Handler")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	fileName := c.Params("filename")

	fileURL, errFileURL := h.detail.GetFileLink(fileName)
	if errFileURL != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseData{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed Retrieve data",
			Error:      errFileURL.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseData{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully Retrieve data",
		Data:       fileURL,
	})
}

func (h *DocumentHandler) DeleteDocumentKM(c *fiber.Ctx) error {
	h.log.Print("Execute DeleteDocumentKM function on Handler")

	permission := utils.ValidateToken(c)
	if permission["status"] != "200" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "The token you entered is invalid",
			Error:      "Authorization failed",
		})
	}

	idFile := c.Params("id")
	author := permission["user"].(map[string]any)["id"].(string)

	if errDeleteDocumentKM := h.delete.DeleteDocumentKM(idFile, author); errDeleteDocumentKM != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseAuth{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed Delete Document",
			Error:      errDeleteDocumentKM.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseAuth{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully Delete Document",
	})
}
