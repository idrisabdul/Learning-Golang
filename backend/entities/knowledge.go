package entities

import "time"

type KnowledgeCategory struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `json:"name"`
	Status    int       `json:"status"`
	CreatedBy int       `json:"created_by"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type Knowledge struct {
	ID              uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Code            string     `json:"code"`
	Title           string     `json:"title"`
	Summary         string     `json:"summary"`
	CategoryID      uint64     `json:"category_id"`
	Status          string     `json:"status"`
	ExpertGroupID   uint64     `json:"expert_group_id"`
	ExperteeID      uint64     `json:"expertee_id"`
	Content         string     `json:"content"`
	Version         float32    `json:"version"`
	ParentID        *uint64    `json:"parent_id"`
	ApprovalGroupID *uint64    `json:"approval_group_id"`
	ApprovedBy      *uint64    `json:"approved_by"`
	ApprovedAt      *time.Time `json:"approved_at"`
	CreatedBy       int        `json:"created_by"`
	CreatedAt       time.Time  `json:"created_at"`
	Access          string     `json:"access"`
	PublishedDate   *time.Time `json:"published_date"`
}

// type KnowledgeContent struct {
// 	ID                        uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
// 	KnowledgeID               string     `json:"knowledge_id"`
// 	Title                     string     `json:"title"`
// 	Version                   int        `json:"version"`
// 	KnowledgeContentListID    *int       `json:"knowledge_content_list_id"`
// 	CompanyID                 *int       `json:"company_id"`
// 	OperationCategory1ID      *int       `json:"operation_category_1_id"`
// 	OperationCategory2ID      *int       `json:"operation_category_2_id"`
// 	ServiceTypeID             *int       `json:"service_type_id"`
// 	ServiceCategory1ID        *int       `json:"service_category_1_id"`
// 	ServiceCategory2ID        *int       `json:"service_category_2_id"`
// 	ExpertGroup               *int       `json:"expert_group"`
// 	Expertee                  *int       `json:"expertee"`
// 	Keyword                   string     `json:"keyword"`
// 	Status                    string     `json:"status"`
// 	IsRetired                 *int       `json:"is_retired"`
// 	PublishedDate             *time.Time `json:"published_date"`
// 	RetireDate                *time.Time `json:"retire_date"`
// 	CreatedAt                 time.Time  `json:"created_at"`
// 	CreatedBy                 string     `json:"created_by"`
// 	UpdatedAt                 *time.Time `json:"updated_at"`
// 	UpdatedBy                 string     `json:"updated_by"`
// 	DeletedAt                 *time.Time `json:"deleted_at"`
// 	DeletedBy                 string     `json:"deleted_by"`
// }

type KnowledgeTag struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	KnowledgeID int       `json:"knowledge_id"`
	Tag         string    `json:"tag"`
	CreatedBy   int       `json:"created_by"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type KnowledgeRelation struct {
	ID           uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	KnowledgeID  string     `json:"knowledge_id"`
	ModuleType   string     `json:"module_type"`
	ModuleID     uint64     `json:"module_id"`
	RelationType string     `json:"relation_type"`
	CreatedAt    time.Time  `json:"created_at"`
	CreatedBy    uint64     `json:"created_by"`
	DeletedAt    *time.Time `json:"deleted_at"`
	DeletedBy    *uint64    `json:"deleted_by"`
	UpdatedAt    *time.Time `json:"updated_at"`
	UpdatedBy    *uint64    `json:"updated_by"`
}

type KnowledgeRelationOption struct {
	ID            int    `gorm:"primaryKey;autoIncrement"`
	Code          string `json:"code"`
	Title         string `json:"title"`
	Summary       string `json:"summary"`
	ExpertGroupId int    `json:"expert_group_id"`
	ExpertGroup   string `json:"expert_group"`
	Status        string `json:"status"`
	Symptom       string `json:"symptom"`
}

func (KnowledgeRelationOption) TableName() string {
	return "knowledge"
}

type KnowledgeRelationPopup struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	KnowledgeId  string    `json:"knowledge_id"`
	ModuleType   string    `json:"module_type"`
	ModuleId     int       `json:"module_id"`
	RelationType string    `json:"relation_type"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    int       `json:"created_by"`
}

func (KnowledgeRelationPopup) TableName() string {
	return "knowledge_relation"
}

type KnowledgeRelationToTicketPopup struct {
	ID            uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	EntityType    string     `gorm:"size:50" json:"entity_type"`
	IDEntityType  int        `json:"id_entity_type"`
	RequestType   string     `gorm:"size:50" json:"request_type"`
	IDRequestType int        `json:"id_request_type"`
	RelationType  string     `gorm:"size:50" json:"relation_type"`
	CreatedAt     *time.Time `json:"created_at"`
	CreatedBy     int        `json:"created_by"`
	DeletedAt     *time.Time `json:"deleted_at"`
	DeletedBy     *int       `json:"deleted_by"`
	UpdatedAt     *time.Time `json:"updated_at"`
	UpdatedBy     *int       `json:"updated_by"`
}

func (KnowledgeRelationToTicketPopup) TableName() string {
	return "knowledge_relation_to_ticket"
}

type KnowledgeRelationList struct {
	ID            int    `gorm:"primaryKey;autoIncrement"`
	Type          string `json:"relation_type"`
	Code          string `json:"code"`
	Title         string `json:"title"`
	Summary       string `json:"summary"`
	ExpertGroupId int    `json:"expert_group_id"`
	ExpertGroup   string `json:"expert_group"`
	Status        string `json:"status"`
	Symptom       string `json:"symptom"`
}

func (KnowledgeRelationList) TableName() string {
	return "knowledge_relation"
}

type KnowledgeRelationToTicketList struct {
	ID            int    `gorm:"primaryKey;autoIncrement"`
	Type          string `json:"relation_type"`
	Code          string `json:"code"`
	Title         string `json:"title"`
	ExpertGroupId int    `json:"expert_group_id"`
	ExpertGroup   string `json:"expert_group"`
	Status        string `json:"status"`
	Symptom       string `json:"symptom"`
}

func (KnowledgeRelationToTicketList) TableName() string {
	return "knowledge_relation_to_ticket"
}

type KnowledgeRelationDelete struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	ModuleId  int       `json:"module_id"`
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy int       `json:"deleted_by"`
}

func (KnowledgeRelationDelete) TableName() string {
	return "knowledge_relation"
}

type KnowledgeRelationToTicketDelete struct {
	IDEntityType        int       `json:"id_entity_type"`
	IDRequestType  int       `json:"id_request_type"`
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy int       `json:"deleted_by"`
}

func (KnowledgeRelationToTicketDelete) TableName() string {
	return "knowledge_relation"
}
