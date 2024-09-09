package entities

import (
	"time"
)

type Problem struct {
	ID                    uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	ProblemCode           string     `gorm:"size:20" json:"problem_code"`
	IDTenant              *int       `json:"tenant_id"`
	IDCompany             *int       `json:"company_id"`
	IDOrganization        *int       `json:"organization_id"`
	IDDepartment          *int       `json:"department_id"`
	IDStatusPBI           *int       `json:"id_status_pbi"`
	IDInvestigationDriver *int       `json:"id_investigation_driver"`
	CoordinatorGroup      *int       `json:"coordinator_group"`
	ProblemCoordinator    *int       `json:"problem_coordinator"`
	KnownErrorLoc         *string    `gorm:"size:255" json:"known_error_location"`
	IDService             *int       `json:"service_id"`
	IDServiceType         *int       `json:"service_type_id"`
	IDSymptom             *int       `json:"symptom_id"`
	Summary               *string    `gorm:"size:50" json:"summary"`
	Note                  *string    `gorm:"type:longtext" json:"note"`
	TargetDate            *time.Time `json:"target_date"`
	Impact                *string    `gorm:"size:20" json:"impact"`
	Urgency               *string    `gorm:"size:20" json:"urgency"`
	Priority              *string    `gorm:"size:20" json:"priority"`
	Channel               *string    `gorm:"size:20" json:"channel"`
	AssignedGroup         *int       `json:"assigned_group"`
	Assignee              *int       `json:"assignee"`
	Status                string     `gorm:"size:50" json:"status"`
	Vendor                *string    `gorm:"size:255" json:"vendor"`
	VendorTicket          *string    `gorm:"size:255" json:"vendor_ticket"`
	Workaround            *string    `gorm:"type:text" json:"workaround"`
	Resolution            *string    `gorm:"type:text" json:"resolution"`
	CreatedBy             *uint      `json:"created_by"`
	CreatedAt             time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedBy             *int       `json:"updated_by"`
	UpdatedAt             *time.Time `json:"updated_at"`
	ApprovalNote          *string    `gorm:"type:text" json:"approval_note"`
	ApprovedBy            *int       `json:"approved_by"`
	ApprovedAt            *time.Time `json:"approved_at"`
	RejectedBy            *int       `json:"rejected_by"`
	RejectedAt            *time.Time `json:"rejected_at"`
	CancelledBy           *int       `json:"cancelled_by"`
	CancelledAt           *time.Time `json:"cancelled_at"`
	CompletedAt           *time.Time `json:"completed_at"`
	InvestigatedAt        *time.Time `json:"investigated_at"`
	ActualClosedDate      *time.Time `json:"actual_closed_date"`
	IDWorkdetail          *int       `json:"workdetail_id"`
	CreatedByWorkdetail   *int       `json:"created_by_workdetail"`
	CreatedAtWorkdetail   *time.Time `json:"created_at_workdetail"`
	NoteWorkdetail        *string    `gorm:"type:longtext" json:"note_workdetail"`
	EndDate               *time.Time `json:"end_date"`
	IsPending             *int       `json:"is_pending"`
	ClosedNote            *string    `json:"closed_note"`
}