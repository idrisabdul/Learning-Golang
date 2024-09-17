package routes

import (
	"sygap_new_knowledge_management/backend/handler/search"
	"sygap_new_knowledge_management/backend/handler/search/history"

	"github.com/gofiber/fiber/v2"
)

type Search struct {
	App            *fiber.App
	SearchHandler  *search.SearchHandler
	HistoryHandler *history.SubmitHistoryHandler
}

func (app *Search) SearchRoute() {
	Search := app.App.Group("/api/v1/km/")
	Search.Get("/search-list", app.SearchHandler.GetSearchList)
	// Detail
	Search.Get("/search-detail/:content_id", app.SearchHandler.GetContentDetail)
	Search.Get("/search-detail-comment/:content_id", app.SearchHandler.GetContentFeedback)
	// Action
	Search.Post("/search-detail-report/:content_id", app.SearchHandler.ReportContentDetail)
	Search.Post("/search-detail-bookmark/:content_id", app.SearchHandler.BookmarkContentDetail)
	Search.Post("/search-detail-feedback/:content_id", app.SearchHandler.FeedbackContentDetail)
	Search.Post("/search-detail/editKm", app.HistoryHandler.SubmitRequestToUpdateKM)
}
