package routes

import (
	"github.com/gofiber/fiber/v2"
	historyhdl "sygap_new_knowledge_management/backend/handler/history"
)

type HistoryRoute struct {
	App         *fiber.App
	HistoryHdlr *historyhdl.HistoryHandler
}

func (c *HistoryRoute) HistoryRoute() {
	HistoryRoute := c.App.Group("/api/v1/km/history")

	HistoryRoute.Get("/list/:id", c.HistoryHdlr.GetHistoryListApproved)
	HistoryRoute.Get("/list/search/:id", c.HistoryHdlr.GetHistoryListApprovedReject)
	HistoryRoute.Get("/list/notif/:id", c.HistoryHdlr.GetHistoryListRequested)

}
