package entities

import "time"

type SearchListResponse struct {
	ID          int         `gorm:"primaryKey" json:"id"`
	Type        string      `json:"type"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Keyword     string      `json:"keyword"`
	Keywords    interface{} `json:"keywords"`
}

func (SearchListResponse) TableName() string {
	return "knowledge_content"
}

// Detail Knowledge Moanagement
type SearchListDescriptionResponse struct {
	Workaround  string `json:"workaround"`
	FixSolution string `json:"fix_solution"`
	Reference   string `json:"reference"`
}
type SearchListQuestionResponse struct {
	Question string `json:"question"`
	Option   string `json:"option"`
}

// Get Detail
type KnowledgeContentSearch struct {
	ID                     int
	Version                int
	KnowledgeId            string
	KnowledgeContentListID int
	Title                  string
	CompanyID              int
	OperationCategory1ID   int `gorm:"column:operation_category_1_id"`
	OperationCategory2ID   int `gorm:"column:operation_category_2_id"`
	ServiceTypeID          int
	ServiceCategory1ID     int `gorm:"column:service_category_1_id"`
	ServiceCategory2ID     int `gorm:"column:service_category_2_id"`
	ExpertGroup            int
	Expertee               *int
	Status                 string
	Author                 string     `json:"author"`
	RetireDate             *time.Time `gorm:"default:null"`
	Keyword                string     `json:"keyword"`
	Type                   string     `json:"type"`
	CreatedAt              time.Time  `json:"created_at"`
	PublishedDate          time.Time  `json:"published_date"`
	LastVisitor            int        `gorm:"column:last_visitor"`
}

func (KnowledgeContentSearch) TableName() string {
	return "knowledge_content"
}

// How To , Reference
type SearchDetailResponse struct {
	ID          int         `gorm:"primaryKey" json:"id"`
	Type        string      `json:"type"`
	Title       string      `json:"title"`
	LastVisitor int         `json:"last_visitor"`
	Keywords    interface{} `json:"keywords"`
	Content     interface{} `json:"content"`
	Sidebar     interface{} `json:"sidebar"`
	Attachment  interface{} `json:"attachment"`
}

type SearchPreviewDetailResponse struct {
	ID            int         `gorm:"primaryKey" json:"id"`
	Type          string      `json:"type"`
	Title         string      `json:"title"`
	LastVisitor   int         `json:"last_visitor"`
	Keywords      interface{} `json:"keywords"`
	Content       interface{} `json:"content"`
	Sidebar       interface{} `json:"sidebar"`
	Attachment    interface{} `json:"attachment"`
	HistoryStatus string      `json:"history_status"`
}

func (SearchDetailResponse) TableName() string {
	return "knowledge_content"
}

type SearchDetailChildResponse struct {
	Question      string `json:"question"`
	Workaround    string `json:"workaround"`
	FixSolution   string `json:"fix_solution"`
	TechnicalNote string `json:"technical_note"`
	Reference     string `json:"reference"`
}

type SearchDetailSidebarResponse struct {
	KnowledgeId   string `json:"knowledge_id"`
	Author        string `json:"author"`
	PublishedDay  string `json:"published_day"`
	Version       string `json:"version"`
	ReportArticle bool   `json:"report_article"`
	Bookmark      bool   `json:"bookmark"`
}

func (SearchDetailChildResponse) TableName() string {
	return "knowledge_content_detail"
}

type SearchDetailFeedbakcResponse struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	KnowledgeId int    `json:"knowledge_id"`
	SubmitterId int    `json:"submitter_id"`
	Submitter   string `json:"submitter"`
	Usefull     string `json:"usefull"`
	Rating      int    `json:"rating"`
	Comment     string `json:"comment"`
	DateSubmit  string `json:"date_submit"`
}

func (SearchDetailFeedbakcResponse) TableName() string {
	return "knowledge_content_feedback"
}

type SearchDetailFeedbakcParentResponse struct {
	Comments    interface{} `json:"comments"`
	RatingTotal string      `json:"rating_total"`
}

type SearchDetailQuestionResponse struct {
	ID       int                  `json:"id"`
	Question string               `json:"question"`
	Options  []SearchDetailOption `gorm:"foreignKey:KnowledgeContentQuestionID;references:ID" json:"options"`
}

func (SearchDetailQuestionResponse) TableName() string {
	return "knowledge_content_question"
}

type SearchDetailOption struct {
	Id                         int    `json:"id"`
	Option                     string `json:"option"`
	Solution                   string `json:"solution"`
	KnowledgeContentQuestionID int    `json:"knowledge_content_question_id"`
}

func (SearchDetailOption) TableName() string {
	return "knowledge_content_option"
}
