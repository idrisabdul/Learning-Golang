package apps

import (
	relationshdl "sygap_new_knowledge_management/backend/handler/relation"
	relationrepo "sygap_new_knowledge_management/backend/repository/relation"
	relationssvc "sygap_new_knowledge_management/backend/services/relation"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func setupRelation(mysql *gorm.DB, logger *logrus.Logger) *relationshdl.RelationHdlr {
	relationRepo := relationrepo.NewRelationRepo(mysql, logger)
	relationService := relationssvc.NewRelationSvc(relationRepo, logger)
	return relationshdl.NewRelationHdlr(relationService, logger)
}
