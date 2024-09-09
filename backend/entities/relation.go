package entities

import (
	"mime/multipart"
	"time"
)

type RelationType struct {
	ID        int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Type      string     `json:"type"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy string     `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy string     `json:"deleted_by"`
}

type SearchRelationParams struct {
	Page              string `json:"page"`
	Limit             string `json:"limit"`
	IDEntityType      string `json:"id_entity_type"`
	RequestType       string `json:"request_type"`
	RequestTypeSearch string `json:"request_type_search"`
	IDService         string `json:"id_service"`
	DateType          string `json:"date_type"`
	IDSymptom         string `json:"id_symptom"`
	Status            string `json:"status"`
	AssignedGroup     string `json:"organization"`
	OrganizationName  string `json:"organization_name"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
}

type SearchReqTypeRelationParams struct {
	IDEntityType      string                `form:"id_entity_type"`
	RequestType       string                `form:"request_type"`
	RequestTypeSearch string                `form:"request_type_search"`
	IDService         string                `form:"id_service"`
	DateType          string                `form:"date_type"`
	AssignedGroup     string                `form:"id_organization"`
	IDSymptom         string                `form:"id_symptom"`
	Organization      string                `form:"organization"`
	Status            string                `form:"status"`
	StartDate         string                `form:"start_date"`
	EndDate           string                `form:"end_date"`
	UploadFile        *multipart.FileHeader `form:"upload_file"`
	IDKnowledge       string                `form:"knowledge_id"`
}

type ProblemRelation struct {
	ID            uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	EntityType    string     `gorm:"size:50" json:"entity_type"`
	IDEntityType  int        `json:"id_entity_type"`
	RequestType   string     `gorm:"size:50" json:"request_type"`
	IDRequestType int        `json:"id_request_type"`
	RelationType  string     `gorm:"size:50" json:"relation_type"`
	Deleted       int        `json:"deleted"`
	CreatedAt     *time.Time `json:"created_at"`
	CreatedBy     int        `json:"created_by"`
	DeletedAt     *time.Time `json:"deleted_at"`
	DeletedBy     *int       `json:"deleted_by"`
	UpdatedAt     *time.Time `json:"updated_at"`
	UpdatedBy     *int       `json:"updated_by"`
}

type Priority struct {
	Critical int `json:"critical"`
	Mayor    int `json:"mayor"`
	High     int `json:"high"`
	Medium   int `json:"medium"`
	Low      int `json:"low"`
}
