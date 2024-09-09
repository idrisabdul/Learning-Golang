package apps

import (
	SearchHdl "sygap_new_knowledge_management/backend/handler/search"
	SearchRepo "sygap_new_knowledge_management/backend/repository/search"
	SearchSvc "sygap_new_knowledge_management/backend/services/search"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetupSearch(db *gorm.DB, log *logrus.Logger) *SearchHdl.SearchHandler {
	SearchRepos := SearchRepo.NewSearchRepos(db, log)
	SearchSvc := SearchSvc.NewSearchService(SearchRepos, log)
	return SearchHdl.NewSearchHandler(SearchSvc, log)
}
