package testcrudhandler

import (
	"strconv"
	"sygap_new_knowledge_management/backend/model"
	testcrudservices "sygap_new_knowledge_management/backend/services/test_crud_services"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TestCrudHandler struct {
	service *testcrudservices.TestCrudSvc
	log     *logrus.Logger
}

func NewTestCrudHandler(service *testcrudservices.TestCrudSvc, log *logrus.Logger) *TestCrudHandler {
	return &TestCrudHandler{service, log}
}

func (h *TestCrudHandler) GetListCrudHandler(c *fiber.Ctx) error {
	h.log.Println("Execute GetListCrudHandler Function")

	getlistcrud, errGetListCrud := h.service.GetListTest()
	if errGetListCrud != nil {
		h.log.Println("Execute GetListCrudHandler Function")
		return c.Status(500).JSON(utils.ResponseData{
			StatusCode: 500,
			Message:    "Failed to retrieve data",
			Error:      errGetListCrud.Error(),
		})
	}

	return c.Status(200).JSON(utils.ResponseData{
		StatusCode: 200,
		Message:    "Successfully retrieve data",
		Data:       getlistcrud,
	})
}

func (h *TestCrudHandler) GetDetailHandler(c *fiber.Ctx) error {
	h.log.Print("Execute GetDetailHandler Function")

	idKnowledgeContent := c.Params("id")
	idKC, _ := strconv.Atoi(idKnowledgeContent)

	detailCrudTest, errDetailCrudTest := h.service.GetDetailTestSvc(idKC)
	if errDetailCrudTest != nil {
		h.log.Printf("Failed Execute DetailUpdateRequest Function: %v", errDetailCrudTest.Error())
		return c.Status(500).JSON(utils.ResponseData{
			StatusCode: 500,
			Message:    "Failed to retrieve data",
			Error:      errDetailCrudTest.Error(),
		})
	}

	return c.Status(200).JSON(utils.ResponseData{
		StatusCode: 200,
		Message:    "Successfully retrieve data",
		Data:       detailCrudTest,
	})
}

func (h *TestCrudHandler) CreateCrudTestHandler(c *fiber.Ctx) error {
	h.log.Print("Execute CreateCrudTestHandler Function")
	data, errParseAndValidate := utils.ParseAndValidate[model.AddKnowledgeContent](c, c.Request().Body())
	if errParseAndValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errParseAndValidate)
	}

	response, errInsertCrudTest := h.service.InsertCrudTestSvc(data)
	if errInsertCrudTest != nil {
		return c.Status(500).JSON(utils.ResponseData{
			StatusCode: 500,
			Message:    "Failed to insert crud test",
		})
	}

	return c.Status(200).JSON(utils.Response{
		StatusCode: 200,
		Message:    "Successfully Add Data Test",
		Data:       response,
	})

}

func (h *TestCrudHandler) UpdateCrudTestHandler(c *fiber.Ctx) error {
	h.log.Print("Execute UpdateCrudTestHandler Function")
	data, errParseAndValidate := utils.ParseAndValidate[model.UpdateKnowledgeContent](c, c.Request().Body())
	if errParseAndValidate != nil {
		return c.Status(400).JSON(errParseAndValidate)
	}

	response, errResponse := h.service.UpdateCrudTestSvc(data)
	if errResponse != nil {
		return c.Status(500).JSON(utils.Response{
			StatusCode: 500,
			Message:    "Failed to update Crud Test",
			Error:      errResponse.Error(),
		})
	}

	return c.Status(200).JSON(utils.Response{
		StatusCode: 200,
		Message:    "Successfully Update Data",
		Data:       response,
	})

}
