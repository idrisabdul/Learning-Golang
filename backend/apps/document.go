package apps

import (
	"sygap_new_knowledge_management/backend/handler/km"
	DocRepo "sygap_new_knowledge_management/backend/repository/km/document"
	DocSvc "sygap_new_knowledge_management/backend/services/km/document"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func setupDocumentCRUD(db *gorm.DB, log *logrus.Logger) *km.DocumentHandler {
	SubmitRepo := DocRepo.NewSubmitRepos(db, log)
	SubmitSvc := DocSvc.NewSubmitService(SubmitRepo, log)

	DetailRepo := DocRepo.NewDetailRepos(db, log)
	DetailSvc := DocSvc.NewDetailService(DetailRepo, log)

	DeleteRepo := DocRepo.NewDeleteDocument(db, log)
	DeleteSvc := DocSvc.NewDeleteService(DeleteRepo, log)

	return km.NewDocumentHandler(SubmitSvc, DetailSvc, DeleteSvc, log)
}
