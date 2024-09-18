package apps

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	historyhdl "sygap_new_knowledge_management/backend/handler/history"
	historyrepo "sygap_new_knowledge_management/backend/repository/history"
	historysvc "sygap_new_knowledge_management/backend/services/history"
)

func setupHistoryListKM(mysql *gorm.DB, logger *logrus.Logger) *historyhdl.HistoryHandler {
	historyRepo := historyrepo.NewHistoryRepo(mysql, logger)
	historyService := historysvc.NewHistoryService(historyRepo, logger)
	return historyhdl.NewHistoryListHandler(historyService, logger)
}