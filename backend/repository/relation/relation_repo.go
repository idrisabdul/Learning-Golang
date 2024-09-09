package relation_repo

import (
	"errors"
	"strconv"
	"strings"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/model"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RelationRepo struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewRelationRepo(db *gorm.DB, logger *logrus.Logger) *RelationRepo {
	return &RelationRepo{db, logger}
}

// GetRelationList retrieves related problem from the database.
func (r *RelationRepo) GetRelationList(params entities.SearchRelationParams) ([]model.RelationFilteredResponseModel, error) {

	var relations []model.ProblemRelationResponseModel
	query := r.db

	query = query.Table("relation_knowledge_management").
		Select(`
        relation_knowledge_management.id AS id_relation, 
        relation_knowledge_management.id_request_type, 
        relation_knowledge_management.entity_type, 
        relation_knowledge_management.id_entity_type, 
        relation_knowledge_management.request_type, 
        relation_knowledge_management.relation_type, 
        problem.id AS id_problem, 
        problem.assigned_group, 
        problem.id_symptom, 
        problem.id_service, 
        problem.problem_code AS problem_code, 
        problem.status AS problem_status, 
        problem.summary AS problem_summary, 
        problem.investigated_at, 
        problem.actual_closed_date, 
        problem.created_at AS problem_created_at, 
		problem.created_at AS created_date, 
        problem.note AS problem_note, 
        problem.resolution AS problem_resolution, 
        problem.actual_closed_date AS problem_resolution_date, 
        problem.priority AS problem_priority, 
        incident.id AS id_incident, 
        incident.incident_code AS incident_code, 
        incident.id_service, 
        incident.id_symptom, 
        incident.solver_group, 
        incident.incident_status, 
        incident.actual_response_date, 
        incident.actual_closed_date AS incident_end, 
        incident.created_at AS incident_created_at, 
		incident.created_at AS created_date, 
        incident.summary AS incident_summary, 
        incident.note AS incident_note, 
        incident.resolution AS incident_resolution, 
        incident.actual_resolution_date AS incident_resolution_date, 
        incident.priority AS incident_priority, 
        change.id AS id_change, 
        change.manager_group, 
        change.id_service, 
		st.service_type_name as change_service_type,
        change.change_code AS change_code, 
        change.status AS change_status, 
        change.summary AS change_summary, 
        change.start_at, 
        change.end_date, 
        change.created_at AS change_created, 
		change.created_at AS change_created_at,
        change.note AS change_note, 
        change.task_approved_note AS change_resolution, 
        change.actual_end_date AS change_resolution_date, 
        change.priority AS change_priority, 
		ke.id as id_known_error,
		ke.known_error_code as known_error_code,
		ke.status as known_error_status,
		ke.summary as known_error_summary,
		ke.product_item_id,
		ke.note as known_error_note,
		ke.priority as known_error_priority,
		ke.created_at as known_error_created_at,
		ke.resolution as known_error_resolution,
		ke.corrected_at as known_error_resolution_date,
		kc.id as id_knowledge_content,
		kc.knowledge_content_list_id as knowledge_content_list_id,
		kc.title as knowledge_content_title,
		kc.status as knowledge_content_status,
		kc.expert_group as knowledge_content_expert_group_id,
		kc.expertee as knowledge_content_expertee_id,
		kc.version as knowledge_content_version,
		kc.retire_date as knowledge_content_retire_date,
		o1.organization_name as organization_name,
		o2.organization_name as organization_name,
		o3.organization_name as organization_name,
		o4.organization_name as organization_name,
		o5.organization_name as organization_name,
        symptoms.id AS id_symptom, 
        symptoms.symptom_name, 
        request_fulfillment.id AS id_request_fulfillment, 
        request_fulfillment.request_fulfillment_code AS request_code, 
        request_fulfillment.status AS request_status, 
        request_fulfillment.id_request_item, 
        request_fulfillment.actual_start_date AS request_start, 
        request_fulfillment.actual_end_date AS request_end, 
        request_fulfillment.created_at AS request_created_at, 
        request_fulfillment.summary AS request_summary, 
        request_fulfillment.note AS request_note, 
        request_fulfillment.task_approved_note AS request_resolution, 
        request_fulfillment.actual_end_date AS request_resolution_date, 
        request_fulfillment.priority AS request_priority
    `).
		Joins("LEFT JOIN problem ON problem.id = relation_knowledge_management.id_request_type AND relation_knowledge_management.request_type = 'problem investigation'").
		Joins("LEFT JOIN incident ON incident.id = relation_knowledge_management.id_request_type AND relation_knowledge_management.request_type = 'incident'").
		Joins("LEFT JOIN `change` ON `change`.id = relation_knowledge_management.id_request_type AND relation_knowledge_management.request_type = 'infrastructure change'").
		Joins("LEFT JOIN known_error ke ON ke.id = relation_knowledge_management.id_request_type AND relation_knowledge_management.request_type = 'known error'").
		Joins("LEFT JOIN knowledge_content kc ON kc.id = relation_knowledge_management.id_request_type AND relation_knowledge_management.request_type = 'knowledge management'").
		Joins("LEFT JOIN request_fulfillment ON request_fulfillment.id = relation_knowledge_management.id_request_type AND relation_knowledge_management.request_type = 'request fulfillment'").
		Joins("LEFT JOIN request_item ON request_item.id = request_fulfillment.id_request_item").
		Joins("LEFT JOIN symptoms ON symptoms.id = incident.id_symptom").
		Joins("LEFT JOIN organization o1 ON o1.id = incident.id_organization").
		Joins("LEFT JOIN organization o2 ON o2.id = problem.id_organization").
		Joins("LEFT JOIN organization o3 ON o3.id = change.id_organization"). // use if in change add field organization
		Joins("LEFT JOIN organization o4 ON o4.id = request_fulfillment.id_organization").
		Joins("LEFT JOIN organization o5 ON o5.id = ke.organization_id").
		Joins(`LEFT JOIN product_categories pc ON pc.id = change.id_service`). // for get id service type change
		Joins(`LEFT JOIN service_type st ON st.id_service = pc.id_service`).   // for get id service type change
		Where("relation_knowledge_management.id_entity_type = ? AND relation_knowledge_management.entity_type = ?", params.IDEntityType, "knowledge management").
		Where("relation_knowledge_management.deleted = ?", 0)

	// Apply filter conditions
	if params.RequestType != "" {
		query = query.Where("relation_knowledge_management.request_type = ?", strings.ToLower(params.RequestType))
	}

	if params.RequestTypeSearch != "" {
		query = query.Where("incident.incident_code LIKE ? OR problem.problem_code LIKE ? OR change.change_code LIKE ? OR request_fulfillment.request_fulfillment_code LIKE ? OR ke.known_error_code LIKE ? OR kc.knowledge_id LIKE ?", "%"+params.RequestTypeSearch+"%", "%"+params.RequestTypeSearch+"%", "%"+params.RequestTypeSearch+"%", "%"+params.RequestTypeSearch+"%", "%"+params.RequestTypeSearch+"%", "%"+params.RequestTypeSearch+"%")
	}

	if params.IDService != "" {
		query = query.Where("incident.id_service_type = ? OR problem.id_service_type = ? OR st.id = ? OR request_fulfillment.id_service_type = ? OR ke.product_item_id = ? OR kc.service_name_id = ?", params.IDService, params.IDService, params.IDService, params.IDService, params.IDService, params.IDService)
	}

	// if params.AssignedGroup != "" {
	// 	query = query.Where("(incident.solver_group = ? OR problem.assigned_group = ? OR change.manager_group = ? OR request_fulfillment.coordinator_group = ?)", params.AssignedGroup, params.AssignedGroup, params.AssignedGroup, params.AssignedGroup)
	// }

	if params.OrganizationName != "" {
		query = query.Where("o1.organization_name LIKE ? OR o2.organization_name LIKE ? OR o3.organization_name LIKE ? OR o4.organization_name LIKE ? OR o5.organization_name LIKE ?", "%"+params.OrganizationName+"%", "%"+params.OrganizationName+"%", "%"+params.OrganizationName+"%", "%"+params.OrganizationName+"%", "%"+params.OrganizationName+"%")
	}

	if params.Status != "" && params.Status != "ALL STATUS" {
		query = query.Where("incident.incident_status = ? OR problem.status = ? OR change.status = ? OR request_fulfillment.status = ? OR ke.status = ? OR kc.status = ?", strings.ToUpper(params.Status), strings.ToUpper(params.Status), strings.ToUpper(params.Status), strings.ToUpper(params.Status), strings.ToUpper(params.Status), strings.ToUpper(params.Status))
	}

	if params.StartDate != "" {
		startDate := utils.ConvertStringToTime(params.StartDate)
		switch params.DateType {
		case "2":
			query = query.Where("incident.actual_resolution_date >= ? OR problem.actual_closed_date >= ? OR change.actual_end_date >= ? OR request_fulfillment.actual_end_date >= ? OR ke.created_at >= ? OR kc.created_at >= ?", startDate, startDate, startDate, startDate, startDate, startDate)
		default:
			query = query.Where("incident.created_at >= ? OR problem.created_at >= ? OR change.created_at >= ? OR request_fulfillment.created_at >= ? OR ke.created_at >= ? OR kc.created_at >= ?", startDate, startDate, startDate, startDate, startDate, startDate)
		}
	}

	if params.EndDate != "" {
		endDate := utils.ConvertStringToTime(params.EndDate)
		switch params.DateType {
		case "2":
			query = query.Where("incident.actual_resolution_date <= ? OR problem.actual_closed_date <= ? OR change.actual_end_date <= ? OR request_fulfillment.actual_end_date <= ? OR ke.corrected_at <= ? OR kc.retire_date <= ?", endDate, endDate, endDate, endDate, endDate, endDate)
		default:
			query = query.Where("incident.created_at <= ? OR problem.created_at <= ? OR change.created_at <= ? OR request_fulfillment.created_at <= ? OR ke.created_at <= ? OR kc.created_at <= ?", endDate, endDate, endDate, endDate, endDate, endDate)
		}
	}

	if params.IDSymptom != "" {
		query = query.Where("incident.id_symptom = ? OR problem.id_symptom = ?", params.IDSymptom, params.IDSymptom)
	}

	if err := query.Group("relation_knowledge_management.id").Order("relation_knowledge_management.id DESC").Find(&relations).Error; err != nil {
		r.logger.Error("Error retrieving services: ", err)
		return nil, err
	}

	var filteredRelations []model.RelationFilteredResponseModel

	for _, param := range relations {
		if param.IDIncident != 0 {
			data := model.RelationFilteredResponseModel{
				ID:             param.IncidentCode,
				IDRelation:     param.IDRelation,
				Summary:        param.IncidentSummary,
				Note:           param.IncidentNote,
				Resolution:     param.IncidentResolution,
				Priority:       param.IncidentPriority,
				Status:         param.IncidentStatus,
				RelationType:   param.RelationType,
				CreatedDate:    param.IncidentCreatedAt,
				ResolutionDate: param.IncidentResolutionDate,
				SymptompName:   param.SymptomName,
			}

			filteredRelations = append(filteredRelations, data)
		}

		if param.IDProblem != 0 {
			data := model.RelationFilteredResponseModel{
				ID:             param.ProblemCode,
				IDRelation:     param.IDRelation,
				Summary:        param.ProblemSummary,
				Note:           param.ProblemNote,
				Resolution:     param.ProblemResolution,
				Priority:       param.ProblemPriority,
				Status:         param.ProblemStatus,
				RelationType:   param.RelationType,
				CreatedDate:    param.ProblemCreatedAt,
				ResolutionDate: param.ProblemResolutionDate,
				SymptompName:   param.SymptomName,
			}

			filteredRelations = append(filteredRelations, data)
		}

		if param.IDChange != 0 {
			data := model.RelationFilteredResponseModel{
				ID:             param.ChangeCode,
				IDRelation:     param.IDRelation,
				Summary:        param.ChangeSummary,
				Note:           param.ChangeNote,
				Resolution:     param.ChangeResolution,
				Priority:       param.ChangePriority,
				Status:         param.ChangeStatus,
				RelationType:   param.RelationType,
				CreatedDate:    param.ChangeCreatedAt,
				ResolutionDate: param.ChangeResolutionDate,
				SymptompName:   param.SymptomName,
			}

			filteredRelations = append(filteredRelations, data)
		}

		if param.IDRequestFulfillment != 0 {
			data := model.RelationFilteredResponseModel{
				ID:             param.RequestCode,
				IDRelation:     param.IDRelation,
				Summary:        param.RequestSummary,
				Note:           param.RequestNote,
				Resolution:     param.RequestResolution,
				Priority:       param.RequestPriority,
				Status:         param.RequestStatus,
				RelationType:   param.RelationType,
				CreatedDate:    param.RequestCreatedAt,
				ResolutionDate: param.RequestResolutionDate,
				SymptompName:   param.SymptomName,
			}

			filteredRelations = append(filteredRelations, data)
		}

		if param.IDKnownError != 0 {
			data := model.RelationFilteredResponseModel{
				ID:             param.KnownErrorCode,
				IDRelation:     param.IDRelation,
				Summary:        param.KnownErrorSummary,
				Note:           param.KnownErrorNote,
				Resolution:     param.KnownErrorResolution,
				Priority:       param.KnownErrorPriority,
				Status:         param.KnownErrorStatus,
				RelationType:   param.RelationType,
				CreatedDate:    param.KnownErrorCreatedAt,
				ResolutionDate: param.KnownErrorResolutionDate,
				SymptompName:   param.SymptomName,
			}

			filteredRelations = append(filteredRelations, data)
		}

		if param.IDKnowledgeContent != 0 {
			data := model.RelationFilteredResponseModel{
				ID:             param.KnowledgeContentCode,
				IDRelation:     param.IDRelation,
				Summary:        param.KnowledgeContentSummary,
				Note:           param.KnowledgeContentNote,
				Resolution:     param.KnowledgeContentResolution,
				Priority:       param.KnowledgeContentPriority,
				Status:         param.KnowledgeContentStatus,
				RelationType:   param.RelationType,
				CreatedDate:    param.KnownErrorCreatedAt,
				ResolutionDate: param.KnownErrorResolutionDate,
				SymptompName:   param.SymptomName,
			}

			filteredRelations = append(filteredRelations, data)
		}
	}

	return filteredRelations, nil
}

// GetReqTypeRelationList retrieves related problem from the database.
func (r *RelationRepo) GetReqTypeRelationList(params entities.SearchReqTypeRelationParams, idUser string, codeList []string) ([]model.ReqTypeRelationResponseModel, error) {

	reqType := strings.ToLower(params.RequestType)
	var relations []model.ReqTypeRelationResponseModel
	query := r.db

	switch reqType {
	case "incident":
		query = query.Table("incident").
			Select(`incident.id as id, incident.incident_code as code, incident.id_symptom, incident.actual_response_date as start_date, incident.actual_resolution_date as end_date, incident.incident_status as status, incident.created_at as created_date, incident.actual_resolution_date as resolution_date, incident.id_organization as organization,
		organization.organization_name, service_type.type, service_type.service_type_name, service_type.id_service, 
		symptoms.symptom_name as summary, 
		product_categories.name as service_name`).
			Joins(`LEFT JOIN service_type ON service_type.id = incident.id_service_type`).
			Joins(`LEFT JOIN product_categories ON product_categories.id = service_type.id_service`).
			Joins(`LEFT JOIN symptoms ON symptoms.id = incident.id_symptom`).
			Joins(`LEFT JOIN organization ON organization.id = incident.id_organization`).
			Joins(`LEFT JOIN relation_knowledge_management rkm on rkm.id_request_type = incident.id AND rkm.deleted = 0 AND rkm.request_type = ?`, "incident").
			Where(`incident.id NOT IN (Select rkm2.id_request_type FROM relation_knowledge_management rkm2 Where rkm2.deleted = 0 AND rkm2.request_type = ? AND rkm2.id_entity_type = ?)`, "incident", params.IDKnowledge).
			Group("incident.id").
			Order("incident.id DESC")

		if params.RequestTypeSearch != "" {
			query = query.Where("incident.incident_code LIKE ?", "%"+params.RequestTypeSearch+"%")
		}

		if params.IDService != "" {
			query = query.Where("incident.id_service_type = ?", params.IDService)
		}

		if params.StartDate != "" {
			startDate := utils.ConvertStringToTime(params.StartDate)
			switch params.DateType {
			case "resolution_date":
				query = query.Where("incident.actual_resolution_date >= ?", startDate)
			default:
				query = query.Where("incident.created_at >= ?", startDate)
			}
		}

		if params.EndDate != "" {
			endDate := utils.ConvertStringToTime(params.EndDate)
			switch params.DateType {
			case "resolution_date":
				query = query.Where("incident.actual_resolution_date <= ?", endDate)
			default:
				query = query.Where("incident.created_at <= ?", endDate)
			}
		}

		if params.AssignedGroup != "" {
			query = query.Where("incident.solver_group", params.AssignedGroup)
		}

		if params.Status != "" {
			query = query.Where("incident.incident_status = ?", strings.ToUpper(params.Status))
		}

		if params.IDSymptom != "" {
			query = query.Where("incident.id_symptom = ?", params.IDSymptom)
		}

		if params.Organization != "" {
			query = query.Where("incident.id_organization = ?", params.Organization)
		}
	case "infrastructure change":
		query = query.Table("change").
			Select(`change.id as id, change.change_code as code, change.id_service, change.summary, change.actual_start_date as start_date, change.actual_end_date as end_date, change.status as status, change.created_at as created_date, change.actual_end_date as resolution_date, 
			organization.organization_name, change.id_organization as organization, product_categories.id as service_id, product_categories.name as service_name`).
			Joins(`LEFT JOIN product_categories ON product_categories.id = change.id_service`).
			Joins(`LEFT JOIN service_type ON service_type.id_service = product_categories.id_service`).
			Joins(`LEFT JOIN organization ON organization.id = change.id_organization`).
			Joins(`LEFT JOIN relation_knowledge_management rkm on rkm.id_request_type = change.id AND rkm.deleted = 0 AND rkm.request_type = ?`, "infrastructure change").
			Where(`change.id NOT IN (Select rkm2.id_request_type FROM relation_knowledge_management rkm2 Where rkm2.deleted = 0 AND rkm2.request_type = ? AND rkm2.id_entity_type = ?)`, "infrastructure change", params.IDKnowledge).
			Group("change.id").
			Order("change.created_at DESC")

		if params.RequestTypeSearch != "" {
			query = query.Where("change.change_code LIKE ?", "%"+params.RequestTypeSearch+"%")
		}

		if params.IDService != "" {
			query = query.Where("service_type.id = ?", params.IDService)
		}

		if params.StartDate != "" {
			startDate := utils.ConvertStringToTime(params.StartDate)
			switch params.DateType {
			case "resolution_date":
				query = query.Where("change.actual_end_date >= ?", startDate)
			default:
				query = query.Where("change.created_at >= ?", startDate)
			}
		}

		if params.EndDate != "" {
			endDate := utils.ConvertStringToTime(params.EndDate)
			switch params.DateType {
			case "resolution_date":
				query = query.Where("change.actual_end_date <= ?", endDate)
			default:
				query = query.Where("change.created_at <= ?", endDate)
			}
		}

		if params.AssignedGroup != "" {
			query = query.Where("change.solver_group", params.AssignedGroup)
		}

		if params.Status != "" {
			query = query.Where("change.status = ?", strings.ToUpper(params.Status))
		}

		if params.Organization != "" {
			query = query.Where("change.id_organization = ?", params.Organization)
		}

	case "problem investigation":
		query = query.Table("problem").
			Select(`problem.id as id, problem.problem_code as code, problem.id_symptom, problem.investigated_at as start_date, problem.actual_closed_date as end_date, problem.status as status, problem.created_at as created_date, problem.actual_closed_date as resolution_date, 
			organization.organization_name, problem.id_organization as organization, service_type.type, service_type.service_type_name, service_type.id_service, 
		symptoms.symptom_name as summary, 
		product_categories.name as service_name`).
			Joins(`LEFT JOIN service_type ON service_type.id = problem.id_service_type`).
			Joins(`LEFT JOIN product_categories ON product_categories.id = service_type.id_service`).
			Joins(`LEFT JOIN symptoms ON symptoms.id = problem.id_symptom`).
			Joins(`LEFT JOIN organization ON organization.id = problem.id_organization`).
			Joins(`LEFT JOIN relation_knowledge_management rkm on rkm.id_request_type = problem.id AND rkm.deleted = 0 AND rkm.request_type = ?`, "problem investigation").
			Where(`problem.id NOT IN (Select rkm2.id_request_type FROM relation_knowledge_management rkm2 Where rkm2.deleted = 0 AND rkm2.request_type = ? AND rkm2.id_entity_type = ?)`, "problem investigation", params.IDKnowledge).Group("problem.id").
			Order("problem.created_at DESC")

		if params.RequestTypeSearch != "" {
			query = query.Where("problem.problem_code LIKE ?", "%"+params.RequestTypeSearch+"%")
		}

		if params.IDService != "" {
			query = query.Where("problem.id_service_type = ?", params.IDService)
		}

		if params.StartDate != "" {
			startDate := utils.ConvertStringToTime(params.StartDate)
			switch params.DateType {
			case "resolution_date":
				query = query.Where("problem.actual_closed_date >= ?", startDate)
			default:
				query = query.Where("problem.created_at >= ?", startDate)
			}
		}

		if params.EndDate != "" {
			endDate := utils.ConvertStringToTime(params.EndDate)
			switch params.DateType {
			case "resolution_date":
				query = query.Where("problem.actual_closed_date <= ?", endDate)
			default:
				query = query.Where("problem.created_at <= ?", endDate)
			}
		}

		if params.AssignedGroup != "" {
			query = query.Where("problem.solver_group", params.AssignedGroup)
		}

		if params.Status != "" {
			query = query.Where("problem.status = ?", strings.ToUpper(params.Status))
		}

		if params.IDSymptom != "" {
			query = query.Where("problem.id_symptom = ?", params.IDSymptom)
		}

		if params.Organization != "" {
			query = query.Where("problem.id_organization = ?", params.Organization)
		}

	case "request fulfillment":
		query = query.Table("request_fulfillment").
			Select(`request_fulfillment.id as id, request_fulfillment.request_fulfillment_code as code, request_fulfillment.id_request_item, request_fulfillment.actual_start_date as start_date, request_fulfillment.actual_end_date as end_date, request_fulfillment.status as status, request_fulfillment.created_at as created_date, request_fulfillment.actual_end_date as resolution_date,
			organization.organization_name, request_fulfillment.id_organization as organization, service_type.type, service_type.service_type_name as service_type_name,
			product_categories.name as service_name,
		request_item.request_item_name as summary`).
			Joins(`LEFT JOIN service_type ON service_type.id = request_fulfillment.id_service_type`).
			Joins(`LEFT JOIN product_categories ON product_categories.id = service_type.id_service`).
			Joins(`LEFT JOIN request_item ON request_item.id = request_fulfillment.id_request_item`).
			Joins(`LEFT JOIN organization ON organization.id = request_fulfillment.id_organization`).
			Joins(`LEFT JOIN relation_knowledge_management rkm on rkm.id_request_type = request_fulfillment.id AND rkm.deleted = 0 AND rkm.request_type = ?`, "request fulfillment").
			Where(`request_fulfillment.id NOT IN (Select rkm2.id_request_type FROM relation_knowledge_management rkm2 Where rkm2.deleted = 0 AND rkm2.request_type = ? AND rkm2.id_entity_type = ?)`, "request fulfillment", params.IDKnowledge).Group("request_fulfillment.id").
			Order("request_fulfillment.created_at DESC")

		if params.RequestTypeSearch != "" {
			query = query.Where("request_fulfillment.request_fulfillment_code LIKE ?", "%"+params.RequestTypeSearch+"%")
		}

		if params.IDService != "" {
			query = query.Where("request_fulfillment.id_service_type = ?", params.IDService)
		}

		if params.StartDate != "" {
			startDate := utils.ConvertStringToTime(params.StartDate)
			switch params.DateType {
			case "resolution_date":
				query = query.Where("request_fulfillment.actual_end_date >= ?", startDate)
			default:
				query = query.Where("request_fulfillment.created_at >= ?", startDate)
			}
		}

		if params.EndDate != "" {
			endDate := utils.ConvertStringToTime(params.EndDate)
			switch params.DateType {
			case "resolution_date":
				query = query.Where("request_fulfillment.actual_end_date <= ?", endDate)
			default:
				query = query.Where("request_fulfillment.created_at <= ?", endDate)
			}
		}

		if params.AssignedGroup != "" {
			query = query.Where("request_fulfillment.solver_group", params.AssignedGroup)
		}

		if params.Status != "" {
			query = query.Where("request_fulfillment.status = ?", strings.ToUpper(params.Status))
		}

		if params.IDSymptom != "" {
			query = query.Where("request_fulfillment.id_symptom = ?", params.IDSymptom)
		}

		if params.Organization != "" {
			query = query.Where("request_fulfillment.id_organization = ?", params.Organization)
		}
	case "known error":
		query = query.Table("known_error ke").
			Select(`ke.id as id, ke.known_error_code as code, ke.created_at as start_date, ke.corrected_at as end_date, ke.organization_id as organization,
			organization.organization_name, ke.status as status, ke.created_at as created_date, ke.corrected_at as closed_date, 
		service_type.type, service_type.service_type_name, service_type.id_service, product_categories.name as service_name`).
			Joins(`LEFT JOIN service_type ON service_type.id = ke.product_id`).
			Joins(`LEFT JOIN product_categories ON product_categories.id = service_type.id_service`).
			Joins(`LEFT JOIN organization ON organization.id = ke.organization_id`).
			Joins(`LEFT JOIN relation_knowledge_management rkm on rkm.id_request_type = ke.id AND rkm.deleted = 0 AND rkm.request_type = ?`, "known error").
			Where(`ke.id NOT IN (Select rkm2.id_request_type FROM relation_knowledge_management rkm2 Where rkm2.deleted = 0 AND rkm2.request_type = ? AND rkm2.id_entity_type = ?)`, "known error", params.IDKnowledge).
			Group("ke.id").
			Order("ke.id DESC")

		if params.RequestTypeSearch != "" {
			query = query.Where("ke.request_type LIKE ?", "%"+params.RequestTypeSearch+"%")
		}

		if params.IDService != "" {
			query = query.Where("ke.product_item_id = ?", params.IDService)
		}

		if params.StartDate != "" {
			startDate := utils.ConvertStringToTime(params.StartDate)
			switch params.DateType {
			case "resolution_date":
				query = query.Where("ke.created_at >= ?", startDate)
			default:
				query = query.Where("ke.created_at >= ?", startDate)
			}
		}

		if params.EndDate != "" {
			endDate := utils.ConvertStringToTime(params.EndDate)
			switch params.DateType {
			case "resolution_date":
				query = query.Where("ke.corrected_at <= ?", endDate)
			default:
				query = query.Where("ke.corrected_at <= ?", endDate)
			}
		}

		if params.AssignedGroup != "" {
			query = query.Where("ke.assignee_group", params.AssignedGroup)
		}

		if params.Status != "" {
			query = query.Where("ke.status = ?", strings.ToUpper(params.Status))
		}

		if params.IDSymptom != "" {
			query = query.Where("ke.known_error_code = ?", 0) // due to known error doesnt have symptom
		}

		if params.Organization != "" {
			query = query.Where("ke.organization_id = ?", params.Organization)
		}
	}

	if codeList != nil {
		query, _ = r.GetListRelationFromFile(codeList, params.IDKnowledge)
	}

	if err := query.Find(&relations).Error; err != nil {
		r.logger.Error("Error retrieving services: ", err)
		return nil, err
	}

	return relations, nil
}

func (r *RelationRepo) InsertRelation(params model.InsertRelationParams, userId int) error {
	r.logger.Println("Execute function InsertRelation")

	// Split the input string by comma
	splitString := strings.Split(params.IDRequestType, ",")

	// Convert the slice of strings to a list of integers
	var ListIdRequestType []int
	for _, str := range splitString {
		ID, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		ListIdRequestType = append(ListIdRequestType, ID)
	}

	IDParentRelation, err := strconv.Atoi(params.IDParentRelation)
	if err != nil {
		return err
	}
	relationType := strings.ToLower(params.RelationType)

	// Prepare the data for insertion
	dataList := make([]entities.ProblemRelation, len(ListIdRequestType))
	for i, v := range ListIdRequestType {
		dataList[i] = entities.ProblemRelation{
			IDEntityType:  IDParentRelation,
			IDRequestType: v,
			EntityType:    "knowledge management",
			RequestType:   params.RequestType,
			RelationType:  relationType,
			CreatedBy:     userId,
			Deleted:       0,
			DeletedAt:     nil,
			UpdatedAt:     nil,
			UpdatedBy:     nil,
		}
	}

	// Check for duplicates in bulk
	var count int64
	if err := r.db.Table(`relation_knowledge_management`).
		Where(`id_request_type IN (?)`, ListIdRequestType).
		Where(`id_entity_type = ?`, IDParentRelation).
		Where("deleted_at IS NULL").
		Count(&count).Error; err != nil {
		return err // Return error if query fails
	}

	// If count is zero, no duplicates found, proceed to create new records
	if count == 0 {
		now := utils.GetTimeNow("datetime")

		// Prepare data for bulk update
		updateData := map[string]interface{}{
			"created_at": now,
			"updated_at": now,
		}
		// Bulk insert data
		if err := r.db.Table(`relation_knowledge_management`).Create(&dataList).Updates(&updateData).Error; err != nil {
			return err // Return error if creation fails
		}

	} else {
		return errors.New("duplicate data")
	}

	return nil
}

// DeleteRelation will soft delete relations from the database.
func (r *RelationRepo) DeleteRelations(idRelations []int, userId int) error {
	// Get current time
	now := utils.GetTimeNow("datetime")

	// Prepare data for bulk update
	updateData := map[string]interface{}{
		"deleted":    1,
		"deleted_at": now,
		"deleted_by": userId,
		"updated_at": now,
		"updated_by": userId,
	}

	// Execute bulk update query
	err := r.db.Table("relation_knowledge_management").Where("id IN (?)", idRelations).Updates(updateData).Error
	if err != nil {
		r.logger.Error("Failed to delete data from table relation_knowledge_management in DeleteRelations:", err)
		return err
	}

	return nil
}

// get data problem
func (p *RelationRepo) GetDetailKM(ID int) (entities.KnowledgeContent, error) {
	p.logger.Println("Execute function GetDetailKM")

	var detailKnowledge entities.KnowledgeContent
	err := p.db.Table(utils.TABLE_KNOWLEDGE_CONTENT).Where("kc.id = ?", ID).Find(&detailKnowledge).Error
	if err != nil {
		p.logger.Error("failed to get detail problem", err)
		return entities.KnowledgeContent{}, err
	}

	return detailKnowledge, nil
}

func (p *RelationRepo) GetListRelationFromFile(codeList []string, IDReport string) (*gorm.DB, error) {

	query := p.db.Raw(`SELECT * FROM
    (
        -- Incident Section
        SELECT
            incident.id AS id,
            incident.incident_code AS code,
            incident.id_symptom,
            incident.actual_response_date AS start_date,
            incident.actual_resolution_date AS end_date,
            incident.incident_status AS status,
            incident.created_at AS created_date,
            incident.actual_resolution_date AS resolution_date,
            incident.id_organization AS organization,
            'incident' AS request_type_from_file,
            organization.organization_name,
            service_type.type,
            service_type.service_type_name,
            service_type.id_service,
            symptoms.symptom_name AS summary,
            product_categories.name AS service_name,
            NULL AS id_problem_relation,  -- Kolom tambahan
            NULL AS id_request_type       -- Kolom tambahan
        FROM
            incident
            LEFT JOIN service_type ON service_type.id = incident.id_service_type
            LEFT JOIN product_categories ON product_categories.id_service = service_type.id_service
            LEFT JOIN symptoms ON symptoms.id = incident.id_symptom
            LEFT JOIN organization ON organization.id = incident.id_organization
            LEFT JOIN problem_relation pr ON pr.id_request_type = incident.id
            AND pr.deleted = 0
            AND pr.request_type = 'incident'
        WHERE
            incident.id NOT IN (
                SELECT
                    pr2.id_request_type
                FROM
                    problem_relation pr2
                WHERE
                    pr2.deleted = 0
                    AND pr2.request_type = 'incident'
                    AND pr2.id_entity_type = ?
            )
        GROUP BY
            incident.id
   
        UNION ALL
   
        -- Change Section
        SELECT
            ch.id AS id,
            ch.change_code AS code,
            NULL AS id_symptom,            -- Kolom tambahan
            ch.actual_start_date AS start_date,
            ch.actual_end_date AS end_date,
            ch.status AS status,
            ch.created_at AS created_date,
            ch.actual_end_date AS resolution_date,
            ch.id_organization AS organization,
            'infrastructure change' AS request_type_from_file,
            organization.organization_name,
            ch.id_service AS type,
            NULL AS service_type_name,     -- Kolom tambahan
            NULL AS id_service,            -- Kolom tambahan
            ch.summary AS summary,
            product_categories.name AS service_name,
            pr.id AS id_problem_relation,
            pr.id_request_type
        FROM
            `+"`change` ch"+`
            LEFT JOIN product_categories ON product_categories.id = ch.id_service
            LEFT JOIN service_type ON service_type.id_service = product_categories.id_service
            LEFT JOIN organization ON organization.id = ch.id_organization
            LEFT JOIN problem_relation pr ON pr.id_request_type = ch.id
            AND pr.deleted = 0
            AND pr.request_type = 'infrastructure change'
        WHERE
            ch.id NOT IN (
                SELECT
                    pr2.id_request_type
                FROM
                    problem_relation pr2
                WHERE
                    pr2.deleted = 0
                    AND pr2.request_type = 'infrastructure change'
                    AND pr2.id_entity_type = ?
            )
        GROUP BY
            ch.id
   
        UNION ALL
   
        -- Problem Section
        SELECT
            problem.id AS id,
            problem.problem_code AS code,
            problem.id_symptom,
            problem.investigated_at AS start_date,
            problem.actual_closed_date AS end_date,
            problem.status AS status,
            problem.created_at AS created_date,
            problem.actual_closed_date AS resolution_date,
            problem.id_organization AS organization,
            'problem investigation' AS request_type_from_file,
            organization.organization_name,
            service_type.type,
            service_type.service_type_name,
            service_type.id_service,
            symptoms.symptom_name AS summary,
            product_categories.name AS service_name,
            pr.id AS id_problem_relation,
            pr.id_request_type
        FROM
            problem
            LEFT JOIN service_type ON service_type.id = problem.id_service_type
            LEFT JOIN product_categories ON product_categories.id_service = service_type.id_service
            LEFT JOIN symptoms ON symptoms.id = problem.id_symptom
            LEFT JOIN organization ON organization.id = problem.id_organization
            LEFT JOIN problem_relation pr ON pr.id_request_type = problem.id
            AND pr.deleted = 0
            AND pr.request_type = 'problem investigation'
        WHERE
            problem.id NOT IN (
                SELECT
                    pr2.id_request_type
                FROM
                    problem_relation pr2
                WHERE
                    pr2.deleted = 0
                    AND pr2.request_type = 'problem investigation'
                    AND pr2.id_entity_type = ?
            )
        GROUP BY
            problem.id
   
        UNION ALL
   
        -- Request Fulfillment Section
        SELECT
            request_fulfillment.id AS id,
            request_fulfillment.request_fulfillment_code AS code,
            NULL AS id_symptom,            -- Kolom tambahan
            request_fulfillment.actual_start_date AS start_date,
            request_fulfillment.actual_end_date AS end_date,
            request_fulfillment.status AS status,
            request_fulfillment.created_at AS created_date,
            request_fulfillment.actual_end_date AS resolution_date,
            request_fulfillment.id_organization AS organization,
            'request fulfillment' AS request_type_from_file,
            organization.organization_name,
            service_type.type,
            service_type.service_type_name AS service_type_name,
            service_type.id_service,
            request_item.request_item_name AS summary,
            product_categories.name AS service_name,
            pr.id AS id_problem_relation,
            pr.id_request_type
        FROM
            request_fulfillment
            LEFT JOIN service_type ON service_type.id = request_fulfillment.id_service_type
            LEFT JOIN product_categories ON product_categories.id_service = service_type.id_service
            LEFT JOIN request_item ON request_item.id = request_fulfillment.id_request_item
            LEFT JOIN organization ON organization.id = request_fulfillment.id_organization
            LEFT JOIN problem_relation pr ON pr.id_request_type = request_fulfillment.id
            AND pr.deleted = 0
            AND pr.request_type = 'request fulfillment'
        WHERE
            request_fulfillment.id NOT IN (
                SELECT
                    pr2.id_request_type
                FROM
                    problem_relation pr2
                WHERE
                    pr2.deleted = 0
                    AND pr2.request_type = 'request fulfillment'
                    AND pr2.id_entity_type = ?
            )
        GROUP BY
            request_fulfillment.id
   
        UNION ALL
   
        -- Known Error Section
        SELECT
            ke.id AS id,
            ke.known_error_code AS code,
            NULL AS id_symptom,            -- Kolom tambahan
            ke.created_at AS start_date,
            ke.corrected_at AS end_date,
            ke.status AS status,
            ke.created_at AS created_date,
            ke.corrected_at AS resolution_date,
            ke.organization_id AS organization,
            'known error' AS request_type_from_file,
            organization.organization_name,
            service_type.type,
            service_type.service_type_name,
            service_type.id_service,
            NULL AS summary,               -- Kolom tambahan
            product_categories.name AS service_name,
            pr.id AS id_problem_relation,
            pr.id_request_type
        FROM
            known_error ke
            LEFT JOIN service_type ON service_type.id = ke.product_item_id
            LEFT JOIN product_categories ON product_categories.id_service = ke.product_id
            LEFT JOIN organization ON organization.id = ke.organization_id
            LEFT JOIN problem_relation pr ON pr.id_request_type = ke.id
            AND pr.deleted = 0
            AND pr.request_type = 'known error'
        WHERE
            ke.id NOT IN (
                SELECT
                    pr2.id_request_type
                FROM
                    problem_relation pr2
                WHERE
                    pr2.deleted = 0
                    AND pr2.request_type = 'known error'
                    AND pr2.id_entity_type = ?
            )
        GROUP BY
            ke.id
    ) AS combined_results
    WHERE
        code IN (?)
    ORDER BY
        code DESC;
    `, IDReport, IDReport, IDReport, IDReport, IDReport, codeList)

	return query, nil
}
