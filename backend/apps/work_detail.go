package apps

import (
	workshdl "sygap_new_knowledge_management/backend/handler/work-detail"
	worksrepo "sygap_new_knowledge_management/backend/repository/work-detail"
	workssvc "sygap_new_knowledge_management/backend/services/work-detail"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func setupWorkDetail(mysql *gorm.DB, logger *logrus.Logger) *workshdl.WorkDetailHandler {
	workDetailRepo := worksrepo.NewWorkDetail(mysql, logger)
	workDetailService := workssvc.NewWorkDetailService(workDetailRepo, logger)
	return workshdl.NewWorkDetailHandler(workDetailService, logger)
}
