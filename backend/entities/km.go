package entities

import (
	"time"
)

// Everything here is for table consumption purposes. DON'T MIX IT WITH ANOTHER PURPOSED STRUCT

type (
	KnowledgeContent struct {
		ID                     int
		KnowledgeID            string
		Version                int
		KnowledgeContentListID int
		Title                  string
		CompanyID              int
		OperationCategory1ID   int `gorm:"column:operation_category_1_id"`
		OperationCategory2ID   int `gorm:"column:operation_category_2_id"`
		ServiceNameID          int
		ServiceCategory1ID     int `gorm:"column:service_category_1_id"`
		ServiceCategory2ID     int `gorm:"column:service_category_2_id"`
		ExpertGroup            int
		Expertee               *int
		Status                 string
		CreatedBy              int
		RetireDate             time.Time `gorm:"default:null"`
		PublishedDate          time.Time `gorm:"default:null"`
		IsRetired              int       `gorm:"default:1"`
		Keyword                string
		DeletedBy              int    `gorm:"default:null"`
		DeletedAt              string `gorm:"default:null"`
		UpdatedBy              int    `gorm:"default:null"`
		UpdatedAt              string `gorm:"default:null"`
	}

	KnowledgeContentDetail struct {
		ID                 int     `json:"-"`
		KnowledgeContentID int     `json:"knowledge_content_id"`
		Question           *string `json:"question"`
		Error              *string `json:"error"`
		RootCause          *string `json:"root_cause"`
		Workaround         *string `json:"workaround"`
		FixSolution        *string `json:"fix_solution"`
		TechnicalNote      *string `json:"technical_note"`
		Reference          *string `json:"reference"`
	}

	KnowledgeContentOption struct {
		ID                 int
		Label              string
		Solution           string
		Question           string
		DeletedAt          time.Time `gorm:"default:null"`
		DeletedBy          int       `gorm:"default:null"`
		OptionParentId     *int      `gorm:"default:null"`
		KnowledgeContentID int
		Options            []*KnowledgeContentOption `gorm:"foreignKey:OptionParentId" json:"options"`
	}

	KnowledgeContentQuestion struct {
		ID                 int
		KnowledgeContentID int
		Question           string
		DeletedAt          time.Time `gorm:"default:null"`
		DeletedBy          int       `gorm:"default:null"`
	}

	KnowledgeContentLog struct {
		ID                 int
		KnowledgeContentID int
		Action             string
		Note               string
		Status             string
		CreatedBy          int
	}

	KnowledgeContentBookmark struct {
		ID                 int
		KnowledgeContentID int
		EmployeeId         string
	}

	KnowledgeContentReport struct {
		ID                 int
		KnowledgeContentID int
		EmployeeId         string
	}

	KnowledgeContentAttachment struct {
		ID                 int
		KnowledgeContentID int
		Attachment         string
		Filename           string
		Size               int
		DeletedAt          time.Time `gorm:"default:null"`
		DeletedBy          string    `gorm:"default:null"`
	}

	KnowledgeContentFeedback struct {
		ID          int
		KnowledgeId int
		SubmitterId int
		Usefull     string
		Rating      int
		Comment     string
	}

	KnowledgeContentUpdateRequest struct {
		ID                int
		KnowledgeID       int
		UpdateRequestType string
		Status            string
		ArticleVersion    string
		RequestSummary    string
		RequestDetail     string
		SubmitterID       int
		ActualStartDate   time.Time `gorm:"default:null"`
		ActualEndDate     time.Time `gorm:"default:null"`
		SubmitDate        time.Time `gorm:"autoCreateTime"`
		DeletedAt         time.Time `gorm:"default:null"`
		DeletedBy         int       `gorm:"default:null"`
	}

	KnowledgeContentReportAttachment struct {
		ID                              int
		KnowledgeContentID              int
		Attachment                      string
		Filename                        string
		Size                            string
		DeletedAt                       time.Time `gorm:"default:null"`
		DeletedBy                       string    `gorm:"default:null"`
		KnowledgeContentUpdateRequestID int
	}

	KnowledgeContentLogVersion struct {
		ID                 int
		IDKnowledgeContent int
		KeyContent         string
		IsFirst            string    `gorm:"default:'false'"`
		CreatedAt          time.Time `gorm:"autoCreateTime"`
		CreatedBy          string
	}

	HistoryKnowledge struct {
		ID                 int
		KnowledgeContentId int
		Note               string
		Type               string
		Value              string
		Status             string
		Date               time.Time `gorm:"default:null"` // Date when approved or rejected
		Requestor          string
		CreatedAt          time.Time `gorm:"autoCreateTime"`
		CreatedBy          int       `gorm:"default:null"`
		UpdatedAt          time.Time `gorm:"default:null"`
		UpdatedBy          int       `gorm:"default:null"`
		DeletedAt          time.Time `gorm:"default:null"`
		Deleted            int       `gorm:"default:null"`
	}
)
