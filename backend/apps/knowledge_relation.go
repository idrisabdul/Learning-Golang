package apps

import (
	knowledgehdl "sygap_new_knowledge_management/backend/handler/knowledge_relation"
	knowledgerepo "sygap_new_knowledge_management/backend/repository/knowledge_relation"
	knowledgesvc "sygap_new_knowledge_management/backend/services/knowledge_relation"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func setupKnowledgeRelation(mysql *gorm.DB, logger *logrus.Logger) *knowledgehdl.KnowledgeRelationHandler {
	knowledgeRelationRepo := knowledgerepo.NewKnowledgeRelation(mysql, logger)
	knowledgeRelationService := knowledgesvc.NewCiRelationService(knowledgeRelationRepo, logger)
	return knowledgehdl.NewKnowledgeRelationHandler(knowledgeRelationService, logger)
}
