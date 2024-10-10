package model

type ListCrudModel struct {
	KnowledgeId string `json:"knowledge_id"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Status      string `json:"status"`
	CreatedBy   string `json:"created_by"`
	CompanyId   string `json:"company_id"`
	Company     string `json:"company_name"`
}

type GetDetailCrudModel struct {
	KnowledgeId string `json:"knowledge_id"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Status      string `json:"status"`
	CreatedBy   string `json:"created_by"`
	CompanyId   string `json:"company_id"`
	Company     string `json:"company_name"`
}

type AddKnowledgeContent struct {
	Title     string `json:"title"`
	Keyword   string `json:"keyword"`
	Status    string `json:"status"`
	CreatedBy int    `json:"created_by"`
	CompanyId int    `json:"company_id"`
}

type UpdateKnowledgeContent struct {
	ID        int    `json:"knowledge_content_id"`
	Title     string `json:"title"`
	Keyword   string `json:"keyword"`
	Status    string `json:"status"`
	CreatedBy int    `json:"created_by"`
	CompanyId int    `json:"company_id"`
}
