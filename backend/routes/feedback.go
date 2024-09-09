package routes

import (
	feedback_hdl "sygap_new_knowledge_management/backend/handler/feedback"

	"github.com/gofiber/fiber/v2"
)

type FeedbackRoute struct {
	App                   *fiber.App
	FeedbackHdlr *feedback_hdl.FeedbackHandler
}


// Feedback API Group
func (c *FeedbackRoute) SetupFeedback() {
	FeedbackRoute := c.App.Group("/api/v1/km/feedback")
	FeedbackRoute.Get("/:km_id", c.FeedbackHdlr.GetFeedbackList)
	FeedbackRoute.Get("/:km_id/export", c.FeedbackHdlr.ExportFeedbackList)
}
