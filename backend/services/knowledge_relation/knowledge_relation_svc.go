package knowledge_relation_svc

import (
	"fmt"
	"strconv"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/model"
	repository "sygap_new_knowledge_management/backend/repository/knowledge_relation"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type KnowledgeRelationService struct {
	repo *repository.KnowledgeRelationRepository
	log  *logrus.Logger
}

func NewCiRelationService(repo *repository.KnowledgeRelationRepository, logger *logrus.Logger) *KnowledgeRelationService {
	return &KnowledgeRelationService{
		repo: repo,
		log:  logger,
	}
}

func (p *KnowledgeRelationService) SearchKnowledge(c *fiber.Ctx, search_param model.KnowledgeRelationSearch) (interface{}, error) {

	organizations, _ := strconv.Atoi(c.Query("organization", "0"))
	products, _ := strconv.Atoi(c.Query("product", "0"))
	symptoms, _ := strconv.Atoi(c.Query("symptom", "0"))

	// parse data from params to models
	param := model.KnowledgeRelationSearch{
		DateFrom:     search_param.DateFrom,
		DateTo:       search_param.DateTo,
		DateType:     search_param.DateType,
		Organization: organizations,
		Status:       search_param.Status,
		Product:      products,
		Symptom:      symptoms,
		Keyword:      search_param.Keyword,
		KnowledgeId:  search_param.KnowledgeId,
	}

	return p.repo.SearchKnowledge(c, param)
}

func (p *KnowledgeRelationService) SubmitKnowledge(payload_param model.KnowledgeRelationToTicketSubmit, id string) (interface{}, error) {
	user_id, _ := strconv.Atoi(id)
	for _, id_request_type := range payload_param.IDRequestType {
		payload := entities.KnowledgeRelationToTicketPopup{
			IDEntityType:  payload_param.KnowledgeId,
			EntityType:    "knowledge_management",
			IDRequestType: id_request_type.Id,
			RequestType:   "knowledge_management",
			RelationType:  payload_param.RelationType,
			CreatedBy:     user_id,
		}
		p.repo.SubmitKnowledge(payload)
	}

	return "Data success created", nil
}

func (p *KnowledgeRelationService) ListKnowledge(c *fiber.Ctx) (interface{}, error) {
	km_id, _ := utils.GenerateDecoded(c.Params("km_id"))
	organizations, _ := strconv.Atoi(c.Query("organization", "0"))
	products, _ := strconv.Atoi(c.Query("product", "0"))
	symptoms, _ := strconv.Atoi(c.Query("symptom", "0"))
	search_param := model.KnowledgeRelationSearch{
		DateFrom:     c.Query("date_from"),
		DateTo:       c.Query("date_to"),
		DateType:     c.Query("date_type"),
		Organization: organizations,
		Status:       c.Query("status"),
		Product:      products,
		Symptom:      symptoms,
		Keyword:      c.Query("keyword"),
	}
	return p.repo.ListKnowledge(km_id, search_param)
}

func (p *KnowledgeRelationService) ExportKnowledge(c *fiber.Ctx) ([][]string, error) {
	km_id, _ := utils.GenerateDecoded(c.Params("km_id"))
	organizations, _ := strconv.Atoi(c.Query("organization", "0"))
	products, _ := strconv.Atoi(c.Query("product", "0"))
	symptoms, _ := strconv.Atoi(c.Query("symptom", "0"))
	search_param := model.KnowledgeRelationSearch{
		DateFrom:     c.Query("date_from"),
		DateTo:       c.Query("date_to"),
		DateType:     c.Query("date_type"),
		Organization: organizations,
		Status:       c.Query("status"),
		Product:      products,
		Symptom:      symptoms,
	}
	fetch, err := p.repo.ListKnowledge(km_id, search_param)
	if err != nil {
		return nil, err
	}

	var response [][]string
	newData := []string{"Code", "Title", "Symptom", "Expert Group", "Status", "Relation Type"}
	response = append(response, newData)
	for _, param := range fetch {
		newParam := []string{param.Code, param.Title, param.Symptom, param.ExpertGroup, param.Status, param.Type}
		response = append(response, newParam)
	}
	return response, nil
}

func (p *KnowledgeRelationService) DeleteKnowledge(payload_param model.KnowledgeRelationToTicketDelete, id string, c *fiber.Ctx) (interface{}, error) {
	user_id, _ := strconv.Atoi(id)
	for _, param := range *payload_param.IDRequestType {
		payload := entities.KnowledgeRelationToTicketDelete{
			IDEntityType:  *payload_param.KnowledgeId,
			IDRequestType: param.Id,
			DeletedBy:     user_id,
			DeletedAt:     utils.ConvertStringToTime(utils.GetTimeNow("datetime")),
		}
		p.repo.DeleteKnowledgeRelation(payload)
	}

	return payload_param, nil
}

func (p *KnowledgeRelationService) GetHeaderTittleExcelSvc(c *fiber.Ctx) (string, error) {
	p.log.Println("Execute function GetHeaderTittleExcelSvc")
	knowledge_managementIDToInt, _ := strconv.Atoi(c.Params("knowledge_id"))
	detailKnowledgeManagement, _ := p.repo.GetDetailKnowledge(knowledge_managementIDToInt)
	now := utils.ConvertStringToTime(utils.GetTimeNow("datetime"))

	nowCustom := utils.ConvertTimeToString(now, "fullname")

	title := fmt.Sprintf("Ticket code : %v - Exported at : %v ", detailKnowledgeManagement.Code, nowCustom)

	return title, nil
}
