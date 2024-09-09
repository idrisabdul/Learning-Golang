package entities

type KmVisitors struct {
	KnowledgeId int    `json:"knowledge_id"`
	UserId      int    `json:"user_id"`
	VisitDate   string `json:"visit_date"`
	Code        string `json:"code"`
}
