package entities

type SearchContentReport struct {
	ID                 int    `gorm:"primaryKey" json:"id"`
	KnowledgeContentId int    `json:"knowledge_content_id"`
	EmployeeId         int    `json:"employee_id"`
	Notes              string `json:"notes"`
}

func (SearchContentReport) TableName() string {
	return "knowledge_content_report"
}

type SearchContentAttachemntReport struct {
	ID                              int    `gorm:"primaryKey" json:"id"`
	KnowledgeContentId              int    `json:"knowledge_content_id"`
	Attachment                      string `json:"attachment"`
	Filename                        string `json:"filename"`
	Size                            int    `json:"size"`
	KnowledgeContentUpdateRequestID int    `json:"-"`
}

func (SearchContentAttachemntReport) TableName() string {
	return "knowledge_content_report_attachment"
}

type SearchContentUpdateRequest struct {
	ID                int    `json:"-"`
	KnowledgeId       int    `json:"knowledge_id"`
	UpdateRequestType string `json:"update_request_type"`
	ArticleVersion    string `json:"article_version"`
	SubmitterId       int    `json:"submitter_id"`
	RequestSummary    string `json:"request_summary"`
}

func (SearchContentUpdateRequest) TableName() string {
	return "knowledge_content_update_request"
}
