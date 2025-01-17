package routes

import (
	"sygap_new_knowledge_management/backend/handler/km"
	"sygap_new_knowledge_management/backend/handler/km/history"

	"github.com/gofiber/fiber/v2"
)

type KMRoutes struct {
	App     *fiber.App
	List    *km.KMListHandler
	CRUD    *km.FormHandler
	History *history.SubmitHistoryHandler
}

func (app *KMRoutes) SetupKMRoutes() {

	KM := app.App.Group("/api/v1/km")

	// List
	KM.Post("/list", app.List.GetListKM)

	// CRUD
	KM.Post("/form/submit", app.CRUD.SubmitKM)
	KM.Post("/form/update/:step?", app.CRUD.UpdateKM)
	KM.Get("/form/:id", app.CRUD.DetailKM)
	KM.Post("/form/set-close", app.CRUD.SetClosedVersion)
	KM.Post("/form/approvalKM/:id", app.History.ApprovalKM)
}
