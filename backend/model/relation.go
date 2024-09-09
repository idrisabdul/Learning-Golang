package model

import (
	"time"
)

type ProblemRelationResponseModel struct {
	IDRelation               uint64     `json:"id_relation"`
	IDProblem                uint64     `json:"id_problem"`
	IDIncident               uint64     `json:"id_incident"`
	IDChange                 uint64     `json:"id_change"`
	IDRequestFulfillment     uint64     `json:"id_request_fulfillment"`
	IDKnownError             uint64     `json:"id_known_error"`
	IDKnowledgeContent       uint64     `json:"id_knowledge_content"`
	IDSymptom                string     `json:"id_symptom"`
	ProblemCode              string     `json:"problem_code"`
	IncidentCode             string     `json:"incident_code"`
	ChangeCode               string     `json:"change_code"`
	RequestCode              string     `json:"request_code"`
	KnownErrorCode           string     `json:"known_error_code"`
	KnowledgeContentCode     string     `json:"knowledge_id"`
	ProblemSummary           string     `json:"problem_summary"`
	IncidentSummary          string     `json:"incident_summary"`
	ChangeSummary            string     `json:"change_summary"`
	RequestSummary           string     `json:"request_summary"`
	KnowledgeContentSummary  string     `json:"knowledge_content_summary"`
	KnownErrorSummary        string     `json:"known_error_summary"`
	KnownErrorNote           string     `json:"known_error_note"`
	KnownErrorPriority       string     `json:"known_error_priority"`
	KnowledgeContentPriority string     `json:"knowledge_content_priority"`
	KnownErrorStatus         string     `json:"known_error_status"`
	KnowledgeContentStatus   string     `json:"knowledge_content_status"`
	KnownErrorResolution     string     `json:"known_error_resolution"`
	KnowledgeContentResolution string     `json:"knowledge_content_resolution"`
	SymptomName              string     `json:"symptom_name"`
	ProblemNote              string     `json:"problem_note"`
	IncidentNote             string     `json:"incident_note"`
	ChangeNote               string     `json:"change_note"`
	RequestNote              string     `json:"request_note"`
	KnowledgeContentNote     string     `json:"knowledge_content_note"`
	ProblemResolution        string     `json:"problem_resolution"`
	IncidentResolution       string     `json:"incident_resolution"`
	ChangeResolution         string     `json:"change_resolution"`
	RequestResolution        string     `json:"request_resolution"`
	ProblemStatus            string     `json:"problem_status"`
	IncidentStatus           string     `json:"incident_status"`
	ChangeStatus             string     `json:"change_status"`
	RequestStatus            string     `json:"request_status"`
	RelationType             string     `json:"relation_type"`
	ProblemCreatedAt         *time.Time `json:"problem_created_at"`
	IncidentCreatedAt        *time.Time `json:"incident_created_at"`
	ChangeCreatedAt          *time.Time `json:"change_created_at"`
	RequestCreatedAt         *time.Time `json:"request_created_at"`
	KnownErrorCreatedAt      *time.Time `json:"known_error_created_at"`
	ProblemResolutionDate    *time.Time `json:"problem_resolution_date"`
	IncidentResolutionDate   *time.Time `json:"incident_resolution_date"`
	ChangeResolutionDate     *time.Time `json:"change_resolution_date"`
	RequestResolutionDate    *time.Time `json:"request_resolution_date"`
	KnownErrorResolutionDate *time.Time `json:"known_error_resolution_date"`
	KnowledgeContentResolutionDate *time.Time `json:"knowledge_content_resolution_date"`
	ProblemPriority          string     `json:"problem_priority"`
	IncidentPriority         string     `json:"incident_priority"`
	ChangePriority           string     `json:"change_priority"`
	RequestPriority          string     `json:"request_priority"`
}

type ReqTypeRelationResponseModel struct {
	ID               string    `json:"id"`
	Code             string    `json:"code"`
	ServiceName      string    `json:"service_name"`
	ServiceTypeName  string    `json:"service_type_name"`
	Summary          string    `json:"summary"`
	Status           string    `json:"status"`
	StartDate        time.Time `json:"start_date"`
	EndDate          string    `json:"end_date"`
	CreatedDate      time.Time `json:"created_date"`
	OrganizationName string    `json:"organization_name"`
	ResolutionDate   string    `json:"resolution_date"`
}

type RelationFilteredResponseModel struct {
	ID             string     `json:"id"`
	IDRelation     uint64     `json:"id_relation"`
	Summary        string     `json:"summary"`
	Note           string     `json:"note"`
	Resolution     string     `json:"resolution"`
	Priority       string     `json:"priority"`
	Status         string     `json:"status"`
	RelationType   string     `json:"relation_type"`
	SymptompName   string     `json:"symptomp_name"`
	CreatedDate    *time.Time `json:"created_date"`
	ResolutionDate *time.Time `json:"resolution_date"`
}

type InsertRelationParams struct {
	IDParentRelation string `json:"id_parent_relation"`
	IDRequestType    string `json:"id_request_type"`
	RequestType      string `json:"request_type"`
	RelationType     string `json:"relation_type"`
}

type ExportRelationParams struct {
	IDEntityType string `json:"id_entity_type"`
}


