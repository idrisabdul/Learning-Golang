package entities

import "time"

type WorkdetailKnowledgeManagement struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	IDParent    int
	Type        string
	Note        string
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy   int
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedBy   int
}

type WorkDetailHasDocument struct {
	ID           int `gorm:"primaryKey;autoIncrement"`
	IDWorkDetail int
	Type         string
	FileHash     string
	FileOri      string
	FileType     string
	FileSize     string
	FileName     string
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" `
	CreatedBy    int
	DeletedAt    time.Time `gorm:"default:null"`
	DeletedBy    int       `gorm:"default:null"`
}

type WorkDetailChangeResponse struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Type         string    `json:"type"`
	Note         string    `json:"note"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	FileHash     string    `json:"file_hash"`
	FileOri      string    `json:"file_ori"`
	FileSize     string    `json:"file_size"`
	RequestType  string    `json:"request_type"`
	EmployeeName string    `json:"employee_name"`
}
