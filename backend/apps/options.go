package apps

import (
	OptionsHdl "sygap_new_knowledge_management/backend/handler/options"
	OptionsRepo "sygap_new_knowledge_management/backend/repository/options"
	OptionsSvc "sygap_new_knowledge_management/backend/services/options"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetupOpCat(db *gorm.DB, log *logrus.Logger) *OptionsHdl.OperationalCategorysHandler {
	OpCatRepo := OptionsRepo.NewOperationalCategorysRepos(db, log)
	OpCatSvc := OptionsSvc.NewOperationalCategorysService(OpCatRepo, log)
	return OptionsHdl.NewOperationalCategorysHandler(OpCatSvc, log)
}

func SetupProduct(db *gorm.DB, log *logrus.Logger) *OptionsHdl.ProductHandler {
	ProductRepo := OptionsRepo.NewProductRepos(db, log)
	ProductSvc := OptionsSvc.NewProductService(ProductRepo, log)
	return OptionsHdl.NewProductHandler(ProductSvc, log)
}

func SetupProductType(db *gorm.DB, log *logrus.Logger) *OptionsHdl.ProductTypeHandler {
	ProductTypeRepo := OptionsRepo.NewProductTypeRepos(db, log)
	ProductTypeSvc := OptionsSvc.NewProductTypeService(ProductTypeRepo, log)
	return OptionsHdl.NewProductTypeHandler(ProductTypeSvc, log)
}

func SetupOrganization(db *gorm.DB, log *logrus.Logger) *OptionsHdl.OrganizationHandler {
	OrganizationRepo := OptionsRepo.NewOrganizationRepos(db, log)
	OrganizationSvc := OptionsSvc.NewOrganizationService(OrganizationRepo, log)
	return OptionsHdl.NewOrganizationHandler(OrganizationSvc, log)
}

func SetupCompany(db *gorm.DB, log *logrus.Logger) *OptionsHdl.CompanyHandler {
	CompanyRepo := OptionsRepo.NewCompanyRepos(db, log)
	CompanySvc := OptionsSvc.NewCompanyService(CompanyRepo, log)
	return OptionsHdl.NewCompanyHandler(CompanySvc, log)
}

func setupExpertees(db *gorm.DB, log *logrus.Logger) *OptionsHdl.ExperteeHandler {
	ExperteeRepo := OptionsRepo.NewExperteeRepos(db, log)
	ExperteeSvc := OptionsSvc.NewExperteeService(ExperteeRepo, log)
	return OptionsHdl.NewExperteeHandler(ExperteeSvc, log)
}

func setupContentType(db *gorm.DB, log *logrus.Logger) *OptionsHdl.ContentTypeHandler {
	ContentTypeRepo := OptionsRepo.NewContentTypeRepos(db, log)
	ContentTypeSvc := OptionsSvc.NewContentTypeService(ContentTypeRepo, log)
	return OptionsHdl.NewContentTypeHandler(ContentTypeSvc, log)
}

func setupStatus(db *gorm.DB, log *logrus.Logger) *OptionsHdl.StatusHandler {
	StatusRepo := OptionsRepo.NewStatusRepos(db, log)
	StatusSvc := OptionsSvc.NewStatusService(StatusRepo, log)
	return OptionsHdl.NewStatusHandler(StatusSvc, log)
}

func setupSymptoms(db *gorm.DB, log *logrus.Logger) *OptionsHdl.SymptomsHandler {
	SymptomsRepo := OptionsRepo.NewSymptomsRepos(db, log)
	SymptomsSvc := OptionsSvc.NewSymptomsService(SymptomsRepo, log)
	return OptionsHdl.NewSymptomsHandler(SymptomsSvc, log)
}

func setupRequestUpdateStatus(db *gorm.DB, log *logrus.Logger) *OptionsHdl.UpdateRequestHandler {
	RequestUpdateRepo := OptionsRepo.NewUpdateRequestRepos(db, log)
	RequestUpdateSvc := OptionsSvc.NewUpdateRequestService(RequestUpdateRepo, log)
	return OptionsHdl.NewUpdateRequest(RequestUpdateSvc, log)
}

func setupWorkDetailType(db *gorm.DB, log *logrus.Logger) *OptionsHdl.WorkDetailTypeHandler {
	WorkDetailRepo := OptionsRepo.NewWorkDetailTypeRepos(db, log)
	WorkDetailSvc := OptionsSvc.NewworkDetailTypeService(WorkDetailRepo, log)
	return OptionsHdl.NewWorkDetailTypeHandler(WorkDetailSvc, log)
}

func setupRelationType(db *gorm.DB, log *logrus.Logger) *OptionsHdl.RelationTypeHandler {
	RelationTypeRepo := OptionsRepo.NewRelationTypeRepos(db, log)
	RelationTypeSvc := OptionsSvc.NewRelationTypeService(RelationTypeRepo, log)
	return OptionsHdl.NewRelationTypeHandler(RelationTypeSvc, log)
}
