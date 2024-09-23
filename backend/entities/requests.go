package entities

import (
	"mime/multipart"
	"time"
)

type (
	SearchListKM struct {
		Search        string `json:"keyword"`
		Status        string `json:"status"`
		Company       string `json:"company"`
		ContentType   string `json:"source_name"`
		CreateFrom    string `json:"start_date"`
		CreatedTo     string `json:"end_date"`
		ExperteeGroup string `json:"expertee_group"`
		Expertee      string `json:"expertee"`
		PublishedDate string `json:"published_date"`
		ProductName   string `json:"product_name"`
		CreatedDate   string `json:"created_date"`
		AssignTo      string `json:"assign_to"`
		SubmittedByMe bool   `json:"submitted_by_me"`
	}

	SubmitKMNonDecisionTree struct {
		ID                     int    `json:"id,omitempty"`
		Version                int    `json:"version" validate:"required"`
		KnowledgeContentListID int    `json:"type_content"`
		Title                  string `json:"title" validate:"required"`
		// Content field start
		Question      *string `json:"question,omitempty"`
		Error         *string `json:"error,omitempty"`
		RootCause     *string `json:"root_cause,omitempty"`
		Workaround    *string `json:"workaround,omitempty"`
		FixSolution   *string `json:"fix_solution,omitempty"`
		TechnicalNote *string `json:"technical_note,omitempty"`
		Reference     *string `json:"reference,omitempty"`
		// Content field end
		Keyword              []string `json:"keyword" validate:"required"`
		Company              int      `json:"company" validate:"required"`
		OperationCategory1ID int      `json:"operational_category_1" validate:"required"`
		OperationCategory2ID int      `json:"operational_category_2" validate:"required"`
		ServiceType          int      `json:"product_name" validate:"required"`
		ServiceCategory1ID   int      `json:"product_category" validate:"required"`
		ServiceCategory2ID   int      `json:"product_parent_category" validate:"required"`
		ExpertGroup          int      `json:"expert_group" validate:"required"`
		Expertee             *int     `json:"expertee,omitempty"`
		Status               string   `json:"status,omitempty"`
		Note                 string   `default:"-" json:"note"`
		KeyContent           int      `json:"key_content,omitempty"`
		CreatedBy            int
	}

	SubmitKMDecisionTree struct {
		ID                     int
		Version                int                   `json:"version" validate:"required"`
		KnowledgeContentListID int                   `json:"type_content"`
		Title                  string                `json:"title" validate:"required"`
		Question               *string               `json:"question,omitempty"`
		Content                []DecisionTreeContent `json:"content"`
		Keyword                []string              `json:"keyword" validate:"required"`
		Company                int                   `json:"company" validate:"required"`
		OperationCategory1ID   int                   `json:"operational_category_1" validate:"required"`
		OperationCategory2ID   int                   `json:"operational_category_2" validate:"required"`
		ServiceType            int                   `json:"product_name" validate:"required"`
		ServiceCategory1ID     int                   `json:"product_category" validate:"required"`
		ServiceCategory2ID     int                   `json:"product_parent_category" validate:"required"`
		ExpertGroup            int                   `json:"expert_group" validate:"required"`
		Expertee               *int                  `json:"expertee,omitempty"`
		Status                 string                `json:"status,omitempty"`
		Note                   string                `default:"-" json:"note"`
		KeyContent             int                   `json:"key_content,omitempty"`
		CreatedBy              int
	}

	DecisionTreeContent struct {
		ID       int                   `json:"id"`
		Question string                `json:"question"`
		Options  []DecisionTreeOptions `json:"options"`
		Action   string                `json:"action,omitempty"`
	}

	DecisionTreeOptions struct {
		ID                         int    `json:"id"`
		KnowledgeContentQuestionID int    `json:"id_question"`
		Option                     string `json:"option"`
		Answer                     string `json:"answer"`
	}

	SetClosedVersion struct {
		ID   int    `json:"id"`
		Note string `json:"note"`
	}

	SubmitUpdateRequest struct {
		ID                int                     `form:"id,omitempty"`
		KnowledgeID       int                     `form:"knowledge_id,omitempty"`
		UpdateRequestType string                  `form:"update_request_type" validate:"required"`
		ArticleVersion    string                  `form:"version"`
		Status            string                  `form:"status" validate:"required"`
		RequestSummary    string                  `form:"request_summary" validate:"required"`
		RequestDetail     string                  `form:"request_detail" validate:"required"`
		SubmitDate        time.Time               `form:"submit_date"`
		Submitter         int                     `form:"-"`
		Attachment        []*multipart.FileHeader `form:"-"`
	}

	SubmitWorkDetail struct {
		ID         int
		IDParent   int                     `form:"id_knowledge"`
		Type       string                  `form:"type"`
		Note       string                  `form:"notes"`
		SubmitDate time.Time               `form:"submit_date"`
		Submitter  int                     `form:"-"`
		CreatedBy  int                     `form:"-"`
		Attachment []*multipart.FileHeader `form:"-"`
	}

	RequestHistoryKnowledge struct {
		KnowledgeContentID string `json:"knowledge_content_id" validate:"required"`
		Note               string `json:"note" validate:"required"`
		Type               string `json:"type" validate:"required"`
		Value              string `json:"value" validate:"required"`
	}

	ApprovalKM struct {
		ApprovedStatus string `json:"approved_status" validate:"required"`
	}
)
