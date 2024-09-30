package entities

import "time"

type (
	ListKM struct {
		ID            int    `json:"id"`
		KnowledgeID   string `json:"knowledge_id"`
		Title         string `json:"title"`
		SourceName    string `json:"source_name"`
		Status        string `json:"status,omitempty"`
		Author        string `json:"author"`
		Assignee      string `json:"assignee"`
		ModifiedDate  string `json:"modified_date"`
		RequestUpdate string `json:"request_update,omitempty"`
	}

	ListUpdateRequest struct {
		ID                int    `json:"id"`
		UpdateRequestType string `json:"update_request_type"`
		ArticleVersion    string `json:"article_version"`
		StatusID          int    `json:"status_id"`
		Status            string `json:"status"`
		RequestSummary    string `json:"request_summary"`
		RequestDetail     string `json:"request_detail"`
		SubmitterUsername string `json:"submitter_username"`
		SubmitterName     string `json:"submitter_name"`
	}
	DetailKM struct {
		ID                     int    `json:"id"`
		KnowledgeID            string `json:"knowledge_id"`
		KnowledgeContentListID int    `json:"type_content"`
		Version                int    `json:"version"`
		Title                  string `json:"title"`
		// content field start
		Question      *string `json:"question,omitempty"`
		Error         *string `json:"error,omitempty"`
		RootCause     *string `json:"root_cause,omitempty"`
		Workaround    *string `json:"workaround,omitempty"`
		FixSolution   *string `json:"fix_solution,omitempty"`
		TechnicalNote *string `json:"technical_note,omitempty"`
		Reference     *string `json:"reference,omitempty"`
		// content field end
		Keywords             []string         `json:"keyword" gorm:"-"`
		CompanyID            int              `json:"company"`
		OperationCategory1ID int              `json:"operational_category_1" gorm:"column:operation_category_1_id"`
		OperationCategory2ID int              `json:"operational_category_2" gorm:"column:operation_category_2_id"`
		ServiceNameID        int              `json:"product_name"`
		ServiceCategory1ID   int              `json:"product_category" gorm:"column:service_category_1_id"`
		ServiceCategory2ID   int              `json:"product_parent_category" gorm:"column:service_category_2_id"`
		ExpertGroup          int              `json:"expert_group"`
		Expertee             *int             `json:"expertee,omitempty"`
		Status               string           `json:"status"`
		Author               string           `json:"author"`
		RetireDate           *time.Time       `json:"retire_date"`
		PublishedDate        *time.Time       `json:"published_date"`
		Keyword              string           `json:"-"`
		KeyContent           string           `json:"key_content,omitempty"`
		CreatedBy            string           `json:"created_by"`
		Permission           ButtonPermission `json:"button_permission" gorm:"-"`
	}

	DetailKMDecisionTree struct {
		ID                     int                       `json:"id"`
		KnowledgeID            string                    `json:"knowledge_id"`
		KnowledgeContentListID int                       `json:"type_content"`
		Version                int                       `json:"version"`
		Title                  string                    `json:"title"`
		Question               string                    `json:"question,omitempty"`
		Content                []*KnowledgeContentOption `json:"content" gorm:"-"`
		Keywords               []string                  `json:"keyword" gorm:"-"`
		CompanyID              int                       `json:"company"`
		OperationCategory1ID   int                       `json:"operational_category_1" gorm:"column:operation_category_1_id"`
		OperationCategory2ID   int                       `json:"operational_category_2" gorm:"column:operation_category_2_id"`
		ServiceNameID          int                       `json:"product_name"`
		ServiceCategory1ID     int                       `json:"product_category" gorm:"column:service_category_1_id"`
		ServiceCategory2ID     int                       `json:"product_parent_category" gorm:"column:service_category_2_id"`
		ExpertGroup            int                       `json:"expert_group"`
		Expertee               *int                      `json:"expertee,omitempty"`
		Status                 string                    `json:"status"`
		Author                 string                    `json:"author"`
		RetireDate             *time.Time                `json:"retire_date"`
		PublishedDate          *time.Time                `json:"published_date"`
		Keyword                string                    `json:"-"`
		KeyContent             string                    `json:"key_content,omitempty"`
		Permission             ButtonPermission          `json:"button_permission" gorm:"-"`
	}

	DetailDocument struct {
		ID         int    `json:"id"`
		Filename   string `json:"file_name"`
		FileHash   string `json:"file_hash,omitempty"`
		Attachment string `json:"attachment,omitempty"`
		Size       int    `json:"size"`
	}

	DetailUpdateRequest struct {
		ID                  int              `json:"id"`
		UpdateRequestTypeID int              `json:"update_request_type_id"`
		UpdateRequestType   string           `json:"update_request_type"`
		StatusID            int              `json:"status_id"`
		Status              string           `json:"status"`
		RequestDetail       string           `json:"request_detail"`
		RequestSummary      string           `json:"request_summary"`
		Submitter           string           `json:"submitter"`
		SubmitDate          string           `json:"submit_date"`
		ActualStartDate     *time.Time       `json:"actual_start_date"`
		ActualEndDate       *time.Time       `json:"actual_end_date"`
		File                []DetailDocument `json:"file" gorm:"-"`
	}

	ButtonPermission struct {
		NextStep   bool `json:"is_next_step"`
		Cancel     bool `json:"is_cancel"`
		Reject     bool `json:"is_reject"`
		Approve    bool `json:"is_approve"`
		Submit     bool `json:"is_submit"`
		Save       bool `json:"is_save"`
		Retire     bool `json:"is_retire"`
		Return     bool `json:"is_return"`
		NewVersion bool `json:"is_new_version"`
	}
)
