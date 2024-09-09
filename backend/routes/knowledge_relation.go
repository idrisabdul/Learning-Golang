package routes

import (
	knowledgehdl "sygap_new_knowledge_management/backend/handler/knowledge_relation"

	"github.com/gofiber/fiber/v2"
)

type KnowledgeRelationRoute struct {
	App                   *fiber.App
	KnowledgeRelationHdlr *knowledgehdl.KnowledgeRelationHandler
}

// Knowledge Relation API Group
func (c *KnowledgeRelationRoute) SetupKnowledgeRelation() {
	KnowledgeRelationRoute := c.App.Group("/api/v1/km/knowledge-relation")
	// Knowledge option
	KnowledgeRelationRoute.Post("/search", c.KnowledgeRelationHdlr.SearchKnowledge)
	KnowledgeRelationRoute.Post("/submit", c.KnowledgeRelationHdlr.SubmitKnowledge)
	KnowledgeRelationRoute.Get("/:km_id/list", c.KnowledgeRelationHdlr.ListKnowledge)
	KnowledgeRelationRoute.Get("/:km_id/export", c.KnowledgeRelationHdlr.ExportKnowledge)
	KnowledgeRelationRoute.Post("/delete", c.KnowledgeRelationHdlr.DeleteKnowledge)
}
