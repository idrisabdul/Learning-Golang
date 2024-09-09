package ci_relation_repo

import (
	"net/url"
	"strconv"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/model"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CiRelationRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewCiRelation(db *gorm.DB, logger *logrus.Logger) *CiRelationRepository {
	return &CiRelationRepository{db, logger}
}

func (p *CiRelationRepository) GetCiType() ([]model.CiRelationTypeFromCiClass, error) {
	var ci_type []model.CiRelationTypeFromCiClass
	if err := p.db.Table("ci_class").
		Where("deleted_at IS NULL").Find(&ci_type).Error; err != nil {
		p.logger.Error("Error retrieving ci type: ", err)
		return nil, err
	}

	return ci_type, nil
}

func (p *CiRelationRepository) GetCiRelationType() ([]entities.CiRelationTypeOption, error) {
	var ci_relation_type []entities.CiRelationTypeOption
	if err := p.db.Table(utils.TABLE_CI_RELATION_TYPE).Where("crt.deleted_at IS NULL").Find(&ci_relation_type).Error; err != nil {
		p.logger.Error("Error retrieving ci relation type: ", err)
		return nil, err
	}

	return ci_relation_type, nil
}

func (p *CiRelationRepository) SubmitCiRelation(payload model.CreateCiRelation) (interface{}, error) {
	ci_relation := p.db.Table("ci_relation").
		Create(&payload)
	if err := ci_relation.Error; err != nil {
		p.logger.Println("Failed Execute function Save Ci Relation")
		return 0, err
	}
	return payload, nil
}

func (p *CiRelationRepository) GetCiRelationList(c *fiber.Ctx) ([]model.CiRelationList, error) {
	ticketIdToInt, _ := strconv.Atoi(c.Params("ticket_id"))
	requestType, _ := url.QueryUnescape(c.Params("request_type")) // remove % from known%20error

	var ci_relation_list []model.CiRelationList

	query := p.db.Table("ci_relation cr").
		Select(`cr.*, cc.class_name as ci_type_name, ci.ci_name as ci_name, cc2.class_name as parent, 
		GROUP_CONCAT(ccha.name SEPARATOR ', ') as attribute_name, csh.created_at as history_date
	`).
		Joins(`LEFT JOIN ci_class cc ON cc.id = cr.ci_type_id`).
		Joins(`LEFT JOIN ci_input ci ON ci.id = cr.ci_name_id`).
		Joins(`LEFT JOIN ci_class_has_attribute ccha ON ccha.id_ci_class = cc.id`).
		Joins(`LEFT JOIN ci_specification_history csh ON csh.ci_input_id = ci.id`).
		Joins(`LEFT JOIN ci_class cc2 ON cc2.id_superclass = cc.id`).
		Where("cr.deleted_at is null AND cr.request_type = ? AND cr.id_request_type = ?", requestType, ticketIdToInt)

	if c.Query("keyword") != "" {
		keyword := "%" + c.Query("keyword") + "%"
		query = query.Where("cc.class_name LIKE ? OR ci.ci_name LIKE ?", keyword, keyword)
	}

	if c.Query("created_at") != "" && c.Query("created_to") != "" {
		query.Where("cr.created_at >= ? AND cr.created_at <= ?", c.Query("created_at"), c.Query("created_to"))
	} else if c.Query("created_at") != "" {
		query.Where("cr.created_at >= ?", c.Query("created_at"))
	} else if c.Query("created_to") != "" {
		query.Where("cr.created_at <= ?", c.Query("created_to"))
	}

	if c.Query("date_history_from") != "" && c.Query("date_history_to") != "" {
		query.Where("csh.created_at >= ? AND csh.created_at <= ?", c.Query("date_history_from"), c.Query("date_history_to"))
	} else if c.Query("date_history_from") != "" {
		query.Where("csh.created_at >= ?", c.Query("date_history_from"))
	} else if c.Query("date_history_to") != "" {
		query.Where("csh.created_at <= ?", c.Query("date_history_to"))
	}

	if c.Query("relation_type") != "" {
		query.Where("cr.relation_type = ?", c.Query("relation_type", "0"))
	}
	if c.Query("ci_type") != "" {
		query.Where("cr.ci_type_id = ?", c.Query("ci_type", "0"))
	}
	if c.Query("attribute_name") != "" {
		query.Where("ccha.id = ?", c.Query("attribute_name", "0"))
	}

	if err := query.Group("cr.id").Order("cr.id desc").Find(&ci_relation_list).Error; err != nil {
		p.logger.Error("Error retrieving ci relation type: ", err)
		return nil, err
	}

	return ci_relation_list, nil
}

func (p *CiRelationRepository) UpdateCiRelation(payload model.CreateCiRelation, relation_id int) (interface{}, error) {
	updateNilValue := map[string]any{
		"history_id": payload.HistoryID,
	}

	if err := p.db.Table("ci_relation").
		Where("id = ?", relation_id).
		Updates(&payload).
		Updates(&updateNilValue).
		Error; err != nil {
		p.logger.Println("Failed to execute Query Update in  UpdateForm:", err)
		return nil, err
	}

	return payload, nil
}

func (p *CiRelationRepository) DeleteCiRelation(relation_id int, payload model.CreateCiRelation) (interface{}, error) {

	if err := p.db.Table("ci_relation").
		Where("id = ?", relation_id).
		Updates(&payload).
		Error; err != nil {
		p.logger.Println("Failed to execute Query Deleted :", err)
		return nil, err
	}

	return payload, nil
}

func (p *CiRelationRepository) GetDataCiNameOption(CiClassId string) ([]model.CiNameOption, error) {
	p.logger.Println("Execute function GetDataCiNameOption")

	var detailCiInput []model.CiNameOption
	err := p.db.Select(`
	ci_input.*, CONCAT(ci_input.ci_name, ' - ' ,ci_input.serial_number) as ci_name_combine`).
		Table("ci_input").Where("ci_class_id = ? AND deleted_at is NULL", CiClassId).Find(&detailCiInput).Error
	if err != nil {
		p.logger.Error("Failed to get detail ci name option", err)
		return nil, err
	}

	return detailCiInput, nil
}

func (p *CiRelationRepository) GetDataCiHistory(CiInputId string) ([]model.CiHistoryOption, error) {
	p.logger.Println("Execute function GetDataCiHistory")

	var detailCiHistory []model.CiHistoryOption
	err := p.db.Table("ci_specification_history csh").
		Select(`csh.*,
		CONCAT('v',ROW_NUMBER() OVER (ORDER BY csh.id), ' - ', DATE_FORMAT(csh.created_at, '%d-%m-%Y %H:%i:%s')) as history_name`).
		Where("csh.ci_input_id = ?", CiInputId).Find(&detailCiHistory).Error
	if err != nil {
		p.logger.Error("Failed to get detail ci history", err)
		return nil, err
	}

	return detailCiHistory, nil
}

func (p *CiRelationRepository) GetAttributeNameOption(CiClassId int) ([]model.CiAttributeNameOption, error) {
	p.logger.Println("Execute function GetAttributeNameOption")

	var detailCiAttributeName []model.CiAttributeNameOption
	err := p.db.Table("ci_class_has_attribute").Where("id_ci_class = ?", CiClassId).
		Find(&detailCiAttributeName).Error
	if err != nil {
		p.logger.Error("Failed to get detail ci attribute", err)
		return nil, err
	}

	return detailCiAttributeName, nil
}
