package routes

import (
	relationhdl "sygap_new_knowledge_management/backend/handler/relation"

	"github.com/gofiber/fiber/v2"
)

type RelationRoute struct {
	App          *fiber.App
	RelationHdlr *relationhdl.RelationHdlr
}

// Relation API Group
func (c *RelationRoute) SetupRelation() {
	RelationRoute := c.App.Group("/api/v1/km/relation")

	// Relation Detail option services
	RelationRoute.Post("/list", c.RelationHdlr.SearchRelationList)
	RelationRoute.Post("/req-type-list", c.RelationHdlr.SearchReqTypeRelationList)
	RelationRoute.Post("/insert-relation", c.RelationHdlr.InsertRelation)
	RelationRoute.Post("/delete-relation", c.RelationHdlr.DeleteRelations)
	RelationRoute.Post("/export", c.RelationHdlr.ExportRelationList)
}
