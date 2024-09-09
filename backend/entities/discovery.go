package entities

import "time"

type DiscoveryMachineLogOption struct {
	CiId      int    `json:"ci_id"`
	MachineId string `json:"machine_id"`
	CiName    string `json:"ci_name"`
	CiKeyword string `json:"ci_keyword"`
	CiCreate  string `json:"ci_create"`
}

func (DiscoveryMachineLogOption) TableName() string {
	return "discovery_machine_log"
}

type DiscoveryStorageLogOption struct {
	CiId      int    `json:"ci_id"`
	MachineId string `json:"machine_id"`
	CiName    string `json:"ci_name"`
	CiKeyword string `json:"ci_keyword"`
	CiCreate  string `json:"ci_create"`
}

func (DiscoveryStorageLogOption) TableName() string {
	return "discovery_storage_log"
}

type DiscoveryPackageLogOption struct {
	CiId      int    `json:"ci_id"`
	MachineId string `json:"machine_id"`
	CiName    string `json:"ci_name"`
	CiKeyword string `json:"ci_keyword"`
	CiCreate  string `json:"ci_create"`
}

func (DiscoveryPackageLogOption) TableName() string {
	return "discovery_package_log"
}

type DiscoveryServiceLogOption struct {
	CiId      int    `json:"ci_id"`
	MachineId string `json:"machine_id"`
	CiName    string `json:"ci_name"`
	CiKeyword string `json:"ci_keyword"`
	CiCreate  string `json:"ci_create"`
}

func (DiscoveryServiceLogOption) TableName() string {
	return "discovery_service_log"
}

type DiscoveryTicketRelation struct {
	ID              int    `gorm:"primaryKey;autoIncrement"`
	CiId            int    `json:"ci_id"`
	MachineId       string `json:"machine_id"`
	TicketTable     string `json:"ticket_table"`
	TicketId        int    `json:"ticket_id"`
	CreatedAt       string `gorm:"default:CURRENT_TIMESTAMP"  json:"created_at"`
	CreatedBy       string `json:"created_by"`
	RelationType    string `json:"relation_type"`
	CiTable         string `json:"ci_table"`
	CiColumn        string `json:"ci_column"`
	CiType          string `json:"ci_type"`
	CiStringSummary string `json:"ci_string_summary"`
	CiKeyword       string `json:"ci_keyword"`
	CiName          string `json:"ci_name"`
}

type DiscoveryTicketRelationList struct {
	ID              int       `gorm:"primaryKey;autoIncrement"`
	CiId            int       `json:"ci_id"`
	MachineId       string    `json:"machine_id"`
	TicketTable     string    `json:"ticket_table"`
	TicketId        int       `json:"ticket_id"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedBy       string    `json:"created_by"`
	RelationType    string    `json:"relation_type"`
	CiTable         string    `json:"ci_table"`
	CiColumn        string    `json:"ci_column"`
	CiType          string    `json:"ci_type"`
	CiStringSummary string    `json:"ci_string_summary"`
	CiKeyword       string    `json:"ci_keyword"`
	CiName          string    `json:"ci_name"`
	Parent          string    `json:"parent"`
}

func (DiscoveryTicketRelationList) TableName() string {
	return "discovery_ticket_relation"
}

type UpdateDiscoveryTicketRelation struct {
	ID              int       `gorm:"primaryKey;autoIncrement"`
	CiId            int       `json:"ci_id"`
	MachineId       string    `json:"machine_id"`
	UpdatedAt       time.Time `json:"updated_at"`
	UpdatedBy       string    `json:"updated_by"`
	RelationType    string    `json:"relation_type"`
	CiTable         string    `json:"ci_table"`
	CiColumn        string    `json:"ci_column"`
	CiType          string    `json:"ci_type"`
	CiStringSummary string    `json:"ci_string_summary"`
	CiKeyword       string    `json:"ci_keyword"`
	CiName          string    `json:"ci_name"`
}

type DeletedDiscoveryTicketRelation struct {
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy string    `json:"deleted_by"`
}

func (UpdateDiscoveryTicketRelation) TableName() string {
	return "discovery_ticket_relation"
}
func (DeletedDiscoveryTicketRelation) TableName() string {
	return "discovery_ticket_relation"
}
