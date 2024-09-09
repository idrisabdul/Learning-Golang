package routes

import (
	"sygap_new_knowledge_management/backend/handler/options"

	"github.com/gofiber/fiber/v2"
)

type Options struct {
	App                 *fiber.App
	OpCatHandler        *options.OperationalCategorysHandler
	Product             *options.ProductHandler
	ProductType   	  *options.ProductTypeHandler
	Organization     *options.OrganizationHandler
	Company             *options.CompanyHandler
	Expertee            *options.ExperteeHandler
	ContentType         *options.ContentTypeHandler
	UpdateRequestStatus *options.UpdateRequestHandler
	Status              *options.StatusHandler
	Symptoms            *options.SymptomsHandler
	WorkDetailType      *options.WorkDetailTypeHandler
	RelationType      *options.RelationTypeHandler
}

func (app *Options) SetupOptions() {
	Options := app.App.Group("/api/v1/km/options")

	Options.Get("/get-operational-category/:idParent?", app.OpCatHandler.GetOpCat)
	Options.Get("/get-product", app.Product.GetListProduct)
	Options.Get("/get-product-type-relation", app.ProductType.GetOptionProductTypeRelationHdlr)
	Options.Get("/get-companies", app.Company.GetCompanies)
	Options.Get("/get-expertee-group", app.Expertee.GetExperteeGroup)
	Options.Get("/get-expertees", app.Expertee.GetExpertees)
	//Status
	Options.Get("/get-status", app.Status.GetOptionStatus)
	//Organization
	Options.Get("/get-organization", app.Organization.GetOrganizationHandler)
	//Sympthom
	Options.Get("/get-list-symptoms-relation", app.Symptoms.GetSymptomsRelationHandler)

	Options.Get("/get-content-type", app.ContentType.GetContentType)

	Options.Get("/get-update-request-status", app.UpdateRequestStatus.GetListUpdateRequestStatus)

	Options.Get("/get-workdetail-type", app.WorkDetailType.GetWorkDetailType)
	Options.Get("/get-relation-type", app.RelationType.GetRelationType)

}
