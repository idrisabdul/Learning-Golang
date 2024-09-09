package routes

import (
	cihdl "sygap_new_knowledge_management/backend/handler/ci_relation"

	"github.com/gofiber/fiber/v2"
)

type CiRelationRoute struct {
	App            *fiber.App
	CiRelationHdlr *cihdl.CiRelationHandler
}

// Work Detail API Group
func (c *CiRelationRoute) SetupCiRelation() {
	CiRelationRoute := c.App.Group("/api/v1/km/ci-relation")
	// Ci option
	CiRelationRoute.Get("/ci-type", c.CiRelationHdlr.GetCiType)
	CiRelationRoute.Get("/ci-name", c.CiRelationHdlr.GetCiName)
	CiRelationRoute.Get("/relation-type", c.CiRelationHdlr.GetRelationType)
	CiRelationRoute.Get("/ci-history", c.CiRelationHdlr.GetCiHistoryHdl)
	CiRelationRoute.Post("/submit", c.CiRelationHdlr.CreateCiRelation)
	CiRelationRoute.Get("/:ticket_id/:request_type", c.CiRelationHdlr.GetCiRelationList)
	CiRelationRoute.Post("/update/:relation_id", c.CiRelationHdlr.UpdateCiRelation)
	CiRelationRoute.Delete("/delete/:relation_id", c.CiRelationHdlr.DeleteCiRelation)
	CiRelationRoute.Get("/ci-attribute-name", c.CiRelationHdlr.GetAttributeNameHdl)
}
