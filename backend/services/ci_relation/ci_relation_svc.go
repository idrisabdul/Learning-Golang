package ci_relation_svc

import (
	"strconv"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/model"
	repository "sygap_new_knowledge_management/backend/repository/ci_relation"
	"sygap_new_knowledge_management/backend/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CiRelationService struct {
	repo *repository.CiRelationRepository
	log  *logrus.Logger
}

func NewCiRelationService(repo *repository.CiRelationRepository, logger *logrus.Logger) *CiRelationService {
	return &CiRelationService{
		repo: repo,
		log:  logger,
	}
}

func (p *CiRelationService) GetCiType() ([]model.CiRelationTypeFromCiClass, error) {
	return p.repo.GetCiType()
}

func (p *CiRelationService) GetCiRelationType() ([]entities.CiRelationTypeOption, error) {
	return p.repo.GetCiRelationType()
}

func (p *CiRelationService) GetCiName(c *fiber.Ctx) ([]model.CiNameOption, error) {
	ci_type := c.Query("ci_type")

	dettailCiNameOption, _ := p.repo.GetDataCiNameOption(ci_type)

	return dettailCiNameOption, nil
}

func (p *CiRelationService) SubmitCiRelation(payload model.CiRelationSubmit, user_id string) (interface{}, error) {
	idToInt, _ := strconv.Atoi(user_id)
	data := model.CreateCiRelation{
		CiTypeID:      payload.CiTypeID,
		CiNameID:      payload.CiNameID,
		RelationType:  payload.RelationType,
		HistoryID:     payload.HistoryID,
		CreatedAt:     utils.ConvertTimeToString(time.Now(), "default"),
		IDRequestType: payload.IDRequestType,
		RequestType:   payload.RequestType,
		CreatedBy:     idToInt,
	}
	return p.repo.SubmitCiRelation(data)
}

func (p *CiRelationService) GetCiRelationList(c *fiber.Ctx) ([]model.CiRelationList, error) {
	return p.repo.GetCiRelationList(c)
}

func (p *CiRelationService) UpdateCiRelation(c *fiber.Ctx, payload model.CiRelationUpdate, user_id string) (interface{}, error) {
	var relation_id, _ = strconv.Atoi(c.Params("relation_id"))
	idToInt, _ := strconv.Atoi(user_id)
	data := model.CreateCiRelation{
		CiTypeID:      payload.CiTypeID,
		CiNameID:      payload.CiNameID,
		RelationType:  payload.RelationType,
		HistoryID:     payload.HistoryID,
		IDRequestType: payload.IDRequestType,
		RequestType:   payload.RequestType,
		CreatedAt:     utils.ConvertTimeToString(time.Now(), "default"),
		CreatedBy:     idToInt,
	}
	return p.repo.UpdateCiRelation(data, relation_id)
}

func (p *CiRelationService) DeleteCiRelation(c *fiber.Ctx, user_id string) (interface{}, error) {
	var relation_id, _ = strconv.Atoi(c.Params("relation_id"))
	idToInt, _ := strconv.Atoi(user_id)
	timeNow := utils.ConvertTimeToString(time.Now(), "default")
	data := model.CreateCiRelation{
		DeletedAt: &timeNow,
		Deleted:   &idToInt,
	}
	return p.repo.DeleteCiRelation(relation_id, data)
}

func (p *CiRelationService) GetDataCiHistorySvc(c *fiber.Ctx) ([]model.CiHistoryResponse, error) {
	p.log.Println("Execute function GetDataCiHistorySvc")

	var resultCiHistory []model.CiHistoryResponse
	getDataCiHistory, _ := p.repo.GetDataCiHistory(c.Query("ci_name_id", "1"))

	for _, data := range getDataCiHistory {

		parseValue, err := utils.ParseJSONString(data.Value)
		if err != nil {
			p.log.Error("Failed to parse json string in GetDataCiHistorySvc")
			return nil, err
		}

		dataCiHistory := model.CiHistoryResponse{
			ID:          data.ID,
			HistoryName: data.HistoryName,
			Log:         parseValue,
		}
		resultCiHistory = append(resultCiHistory, dataCiHistory)
	}

	return resultCiHistory, nil
}

func (p *CiRelationService) GetAttributeNameOptionSvc(c *fiber.Ctx) ([]model.CiAttributeNameOption, error) {
	return p.repo.GetAttributeNameOption(c.QueryInt("ci_type_id", 0))
}
