package apps

import (
	SearchHdl "sygap_new_knowledge_management/backend/handler/search"
	SearchRepo "sygap_new_knowledge_management/backend/repository/search"
	SearchSvc "sygap_new_knowledge_management/backend/services/search"

	HistorySubmitHdl "sygap_new_knowledge_management/backend/handler/search/history"
	HistorySubmitRepo "sygap_new_knowledge_management/backend/repository/search/history"
	HistorySubmitSvc "sygap_new_knowledge_management/backend/services/search/history"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetupSearch(db *gorm.DB, log *logrus.Logger) *SearchHdl.SearchHandler {
	SearchRepos := SearchRepo.NewSearchRepos(db, log)
	SearchSvc := SearchSvc.NewSearchService(SearchRepos, log)
	return SearchHdl.NewSearchHandler(SearchSvc, log)
}

func SetupHistory(db *gorm.DB, log *logrus.Logger) *HistorySubmitHdl.SubmitHistoryHandler {
	HistoryRepo := HistorySubmitRepo.NewSubmitRepos(db, log)
	HistorySvc := HistorySubmitSvc.NewSubmitService(HistoryRepo, log)
	return HistorySubmitHdl.NewHistoryhandler(HistorySvc, log)
}

//func setupHistoryListSearch(mysql *gorm.DB, logger *logrus.Logger) *historyhdl.HistoryHandler {
//	historyRepo := historyrepo.NewHistoryRepo(mysql, logger)
//	historyService := historysvc.NewHistoryService(historyRepo, logger)
//	return historyhdl.NewHistoryListHandler(historyService, logger)
//}
