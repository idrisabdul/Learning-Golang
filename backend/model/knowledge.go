package model

type KnowledgeRelationSearch struct {
	Keyword      string `json:"keyword"`
	DateFrom     string `json:"date_from"`
	DateTo       string `json:"date_to"`
	DateType     string `json:"date_type"`
	Organization int    `json:"organization"`
	Status       string `json:"status"`
	Product      int    `json:"product"`
	Symptom      int    `json:"symptom"`
	KnowledgeId  string `json:"id_knowledge"`
	Company      int    `json:"company"`
}

type KnowledgeRelationSubmit struct {
	KnowledgeId  []KnowledgeId `json:"knowledge_id"  validate:"required"`
	ModuleType   string        `json:"module_type"  validate:"required"`
	ModuleId     int           `json:"module_id"  validate:"required"`
	RelationType string        `json:"relation_type"  validate:"required"`
}

type KnowledgeRelationToTicketSubmit struct {
	KnowledgeId   int           `json:"knowledge_id"  validate:"required"`
	IDRequestType []KnowledgeId `json:"id_request_type"  validate:"required"`
	// RequestType   string        `json:"request_type"  validate:"required"`
	RelationType string `json:"relation_type"  validate:"required"`
}

type KnowledgeId struct {
	Id int `json:"id"  validate:"required"`
}

type KnowledgeRelationToTicketDelete struct {
	KnowledgeId   *int           `json:"knowledge_id"  validate:"required"`
	IDRequestType *[]KnowledgeId `json:"id_request_type"  validate:"required"`
}

type KnowledgeIdInt struct {
	Id int `json:"id"  validate:"required"`
}

type KnowledgeRelationOptionCustom struct {
	ID            int    `json:"ID"`
	Code          string `json:"code"`
	Title         string `json:"title"`
	ExpertGroupId int    `json:"expert_group_id"`
	ExpertGroup   string `json:"expert_group"`
	Status        string `json:"status"`
	Symptom       string `json:"symptom"`
	Product       string `json:"product"`
}

func (KnowledgeRelationOptionCustom) TableName() string {
	return "knowledge"
}
