package routes

import (
	"sygap_new_knowledge_management/backend/handler/km"

	"github.com/gofiber/fiber/v2"
)

type DocumentRoutes struct {
	App      *fiber.App
	Document *km.DocumentHandler
}

func (app *DocumentRoutes) SetupDocumentRoutes() {
	Document := app.App.Group("/api/v1/km/document")

	// Submit to KM
	Document.Post("/post/:id", app.Document.SubmitDocumentKM)

	// Get Document
	Document.Get("/list/:id", app.Document.ListDetailDocumentKM)
	Document.Get("/link/:filename", app.Document.URLDocumentKM)

	// Delete Document
	Document.Put("/delete/:id", app.Document.DeleteDocumentKM)
}
