package routes

import (
	testcrudhandler "sygap_new_knowledge_management/backend/handler/test_crud_handler"

	"github.com/gofiber/fiber/v2"
)

type TestCrudRoute struct {
	App             *fiber.App
	TestCrudHandler *testcrudhandler.TestCrudHandler
}

func (c *TestCrudRoute) SetupTestCrudRoutes() {
	TestCrudHandler := c.App.Group("api/v1/km/test")

	TestCrudHandler.Get("/list", c.TestCrudHandler.GetListCrudHandler)
	TestCrudHandler.Get("/detail/:id", c.TestCrudHandler.GetDetailHandler)
	TestCrudHandler.Post("/add", c.TestCrudHandler.CreateCrudTestHandler)
	TestCrudHandler.Put("/update", c.TestCrudHandler.UpdateCrudTestHandler)
}
