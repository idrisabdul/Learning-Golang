package apps

import (
	cishdl "sygap_new_knowledge_management/backend/handler/ci_relation"
	cirepo "sygap_new_knowledge_management/backend/repository/ci_relation"
	cisvc "sygap_new_knowledge_management/backend/services/ci_relation"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func setupCiRelation(mysql *gorm.DB, logger *logrus.Logger) *cishdl.CiRelationHandler {
	ciRelationRepo := cirepo.NewCiRelation(mysql, logger)
	ciRelationService := cisvc.NewCiRelationService(ciRelationRepo, logger)
	return cishdl.NewCiRelationHandler(ciRelationService, logger)
}
