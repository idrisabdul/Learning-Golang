package routes

import (
	workhdl "sygap_new_knowledge_management/backend/handler/work-detail"

	"github.com/gofiber/fiber/v2"
)

type WorkDetailRoute struct {
	App            *fiber.App
	WorkDetailHdlr *workhdl.WorkDetailHandler
}

// Work Detail API Group
func (c *WorkDetailRoute) SetupWorkDetail() {
	WorkRoute := c.App.Group("/api/v1/km/work-detail")
	// Work Detail option services
	WorkRoute.Get("/list/:id", c.WorkDetailHdlr.GetWorkDetail)
	WorkRoute.Post("/submit/:id", c.WorkDetailHdlr.AddWorkDetail)
}
