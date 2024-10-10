package apps

import (
	testcrudhandler "sygap_new_knowledge_management/backend/handler/test_crud_handler"
	testcrud "sygap_new_knowledge_management/backend/repository/test_crud"
	testcrudservices "sygap_new_knowledge_management/backend/services/test_crud_services"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func setupCrudTest(db *gorm.DB, log *logrus.Logger) *testcrudhandler.TestCrudHandler {
	CrudTestRepos := testcrud.NewTestCrudRepository(db, log)
	CrudTestSvc := testcrudservices.NewTestCrudSvc(CrudTestRepos, log)
	return testcrudhandler.NewTestCrudHandler(CrudTestSvc, log)
}
