package apps

import (
	KMHdl "sygap_new_knowledge_management/backend/handler/km"
	KMRepo "sygap_new_knowledge_management/backend/repository/km"
	FormRepo "sygap_new_knowledge_management/backend/repository/km/form"
	KMSvc "sygap_new_knowledge_management/backend/services/km"
	FormSvc "sygap_new_knowledge_management/backend/services/km/form"

	HistoryHandler "sygap_new_knowledge_management/backend/handler/km/history"
	HistoryRepo "sygap_new_knowledge_management/backend/repository/km/history"
	HistorySvc "sygap_new_knowledge_management/backend/services/km/history"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func setupList(db *gorm.DB, log *logrus.Logger) *KMHdl.KMListHandler {
	KMListRepos := KMRepo.NewKMListRepos(db, log)
	KMListSvc := KMSvc.NewKMListService(KMListRepos, log)
	return KMHdl.NewKMListHandler(KMListSvc, log)
}

func setupCRUD(db *gorm.DB, log *logrus.Logger) *KMHdl.FormHandler {
	// Submit
	SubmitRepos := FormRepo.NewSubmitRepos(db, log)
	SubmitSvc := FormSvc.NewSubmitService(SubmitRepos, log)

	// Detail
	DetailRepos := FormRepo.NewDetailRepos(db, log)
	DetailSvc := FormSvc.NewDetailService(DetailRepos, log)

	// Update
	UpdateRepos := FormRepo.NewUpdateRepos(db, log)
	UpdateSvc := FormSvc.NewUpdateService(UpdateRepos, log)

	return KMHdl.NewFormhandler(SubmitSvc, DetailSvc, UpdateSvc, log)
}

func setupHistory(db *gorm.DB, log *logrus.Logger) *HistoryHandler.SubmitHistoryHandler {
	HistoryRepos := HistoryRepo.NewSubmitRepos(db, log)
	HistoryService := HistorySvc.NewSubmitService(HistoryRepos, log)
	return HistoryHandler.NewHistoryhandler(HistoryService, log)
}
