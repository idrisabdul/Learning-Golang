package routes

import (
	updaterequest "sygap_new_knowledge_management/backend/handler/update-request"

	"github.com/gofiber/fiber/v2"
)

type UpdateRequestRoutes struct {
	App  *fiber.App
	CRUD *updaterequest.UpdateRequestHandler
}

func (app *UpdateRequestRoutes) SetupUpdateRequestRoutes() {

	updateRequest := app.App.Group("/api/v1/km/update-request")

	updateRequest.Get("/list/:id", app.CRUD.GetListUpdateRequest)
	updateRequest.Get("/list/detail/:id", app.CRUD.GetDetailUpdateRequest)

	updateRequest.Post("/submit", app.CRUD.SubmitUpdateRequest)

	updateRequest.Put("/update/:id", app.CRUD.UpdateUpdateRequest)

	updateRequest.Delete("/delete/:id", app.CRUD.DeleteUpdateRequest)
	updateRequest.Delete("/document/delete/:id", app.CRUD.DeleteUpdateRequestAttachment)
}
