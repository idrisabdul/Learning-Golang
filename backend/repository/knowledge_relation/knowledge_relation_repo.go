package knowledge_relation_repo

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/model"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type KnowledgeRelationRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewKnowledgeRelation(db *gorm.DB, logger *logrus.Logger) *KnowledgeRelationRepository {
	return &KnowledgeRelationRepository{db, logger}
}

func (p *KnowledgeRelationRepository) SearchKnowledge(c *fiber.Ctx, search_param model.KnowledgeRelationSearch) (interface{}, error) {
	var ci_relation_list []model.KnowledgeRelationOptionCustom
	query := p.db.Table(utils.TABLE_KNOWLEDGE_CONTENT).Select(`
		kc.id, kc.knowledge_id as code, kc.title, kc.status, kc.expert_group as expert_group_id, kc.created_at,
		o.organization_name as expert_group, s.symptom_name as symptom, pc.name as product
	`).Joins(`
		LEFT JOIN organization o ON o.id = kc.expert_group
	`).Joins(`
		LEFT JOIN knowledge_symptoms ks ON ks.knowledge_id = kc.id
	`).Joins(`
		LEFT JOIN symptoms s ON s.id = ks.symptom_id
	`).Joins(`
		LEFT JOIN product_categories pc ON pc.id_service_type = s.id_service_type
	`).Joins("LEFT JOIN "+utils.TABLE_KNOWLEDGE_RELATION_TO_TICKET+" ON krt.id_entity_type = ?", search_param.KnowledgeId).
		Where("kc.id NOT IN (?)",
			p.db.Table(utils.TABLE_KNOWLEDGE_RELATION_TO_TICKET+"2").
				Select(`krt2.id_request_type`).
				Where("krt2.id_entity_type = ?", search_param.KnowledgeId))

	if search_param.DateType == "created_at" {
		query.Where("kc.created_at >= ? AND kc.created_at <= ?", search_param.DateFrom, search_param.DateTo)
	}
	if search_param.Organization != 0 {
		query.Where("kc.expert_group", search_param.Organization)
	}
	if search_param.Company != 0 {
		query.Where("kc.company_id", search_param.Company)
	}
	if search_param.Status != "" {
		query.Where("kc.status", search_param.Status)
	}
	if search_param.Product != 0 {
		query.Where("pc.id", search_param.Product)
	}
	if search_param.Symptom != 0 {
		query.Where("ks.symptom_id", search_param.Symptom)
	}

	// using keyword search
	keyword := c.Query("keyword")
	if keyword != "" {
		query.Where("kc.status LIKE ? OR kc.knowledge_id LIKE ? OR kc.title LIKE ? OR symptom LIKE ? OR expert_group LIKE ?", keyword, keyword, keyword, keyword, keyword)
	}

	if err := query.Group("kc.id").Find(&ci_relation_list).Error; err != nil {
		p.logger.Error("Error retrieving knowledge relation: ", err)
		return nil, err
	}

	return ci_relation_list, nil
}

func (p *KnowledgeRelationRepository) SubmitKnowledge(payload entities.KnowledgeRelationToTicketPopup) (interface{}, error) {
	knowledge_relation := p.db.Create(&payload)
	if err := knowledge_relation.Error; err != nil {
		p.logger.Println("Failed Execute function Save Knowledge Relation")
		return 0, err
	}
	return payload.ID, nil
}

func (p *KnowledgeRelationRepository) ListKnowledge(knowledge_id string, search_param model.KnowledgeRelationSearch) ([]entities.KnowledgeRelationToTicketList, error) {
	var knowledge_relation []entities.KnowledgeRelationToTicketList
	query := p.db.Table(utils.TABLE_KNOWLEDGE_RELATION_TO_TICKET).Select(`
		krt.id, 
		krt.id_entity_type, 
		krt.relation_type, 
		kc.knowledge_id as code, 
		kc.title, 
		kc.status, 
		kc.expert_group as expert_group_id, 
		rt.type,
		o.organization_name as expert_group, 
		CONCAT(services.name,' - ',symptoms.symptom_name) as symptom
	`).Joins(`
		LEFT JOIN knowledge_content kc ON kc.id = krt.id_request_type
	`).Joins(`
		LEFT JOIN organization o ON o.id = kc.expert_group
	`).Joins(`
		LEFT JOIN knowledge_symptoms ks ON ks.knowledge_id = kc.id
	`).Joins(` 
		LEFT JOIN symptoms ON symptoms.id = ks.symptom_id
	`).Joins(`
		LEFT JOIN services ON services.id = symptoms.id_service
	`).Joins(`
		LEFT JOIN relation_type rt ON rt.type = krt.relation_type
	`).Where("krt.entity_type = 'knowledge_management' AND krt.id_entity_type = ?", knowledge_id).Where("krt.deleted_at IS NULL")

	if search_param.DateType == "created_at" {
		query.Where("kc.created_at >= ? AND kc.created_at <= ?", search_param.DateFrom, search_param.DateTo)
	}
	if search_param.Organization != 0 {
		query.Where("kc.expert_group", search_param.Organization)
	}
	if search_param.Status != "" {
		query.Where("kc.status", search_param.Status)
	}
	if search_param.Product != 0 {
		query.Where("symptoms.id_service", search_param.Product)
	}
	if search_param.Symptom != 0 {
		query.Where("ks.symptom_id", search_param.Symptom)
	}

	// using keyword search
	keyword := search_param.Keyword
	if keyword != "" {
		likeKeyword := "%" + keyword + "%"
		query.Where("kc.status LIKE ? OR kc.knowledge_id LIKE ? OR kc.title LIKE ? OR CONCAT(services.name,' - ',symptoms.symptom_name) LIKE ? OR o.organization_name LIKE ?",
			likeKeyword, likeKeyword, likeKeyword, likeKeyword, likeKeyword, likeKeyword)
	}

	if err := query.Group("krt.id").Order("krt.id desc").
		Find(&knowledge_relation).Error; err != nil {
		p.logger.Error("Error retrieving knowledge relation: ", err)
		return nil, err
	}

	return knowledge_relation, nil
}

func (p *KnowledgeRelationRepository) DeleteKnowledgeRelation(delete entities.KnowledgeRelationToTicketDelete) (string, error) {
	if err := p.db.Table(utils.TABLE_KNOWLEDGE_RELATION_TO_TICKET).
		Where("krt.id_entity_type = ? AND krt.id = ?", delete.IDEntityType, delete.IDRequestType).
		Updates(&delete).
		Error; err != nil {
		return "failed", err
	}

	return "success", nil
}

func (p *KnowledgeRelationRepository) GetDetailKnowledge(ID int) (entities.Knowledge, error) {
	p.logger.Println("Execute function GetDetailKnowledge")

	var detailKnowledge entities.Knowledge
	err := p.db.Table(utils.TABLE_KNOWLEDGE_CONTENT).Where("kc.id = ?", ID).Find(&detailKnowledge).Error
	if err != nil {
		p.logger.Error("failed to get detail knowledge", err)
		return entities.Knowledge{}, err
	}

	return detailKnowledge, nil
}
