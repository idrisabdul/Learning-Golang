package apps

import (
	updaterequestHdl "sygap_new_knowledge_management/backend/handler/update-request"
	updaterequestRepo "sygap_new_knowledge_management/backend/repository/update-request"
	updaterequestSvc "sygap_new_knowledge_management/backend/services/update-request"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func setupUpdateRequest(db *gorm.DB, log *logrus.Logger) *updaterequestHdl.UpdateRequestHandler {

	// Submit
	submitRepo := updaterequestRepo.NewSubmitRepos(db, log)
	submitSvc := updaterequestSvc.NewSubmitService(submitRepo, log)

	//Update
	updateRepo := updaterequestRepo.NewUpdateRepos(db, log)
	updateSvc := updaterequestSvc.NewUpdateService(updateRepo, log)

	// Detail
	detailRepo := updaterequestRepo.NewDetailRepos(db, log)
	detailSvc := updaterequestSvc.NewDetailService(detailRepo, log)

	return updaterequestHdl.NewUpdateRequestHandler(submitSvc, updateSvc, detailSvc, log)

}
