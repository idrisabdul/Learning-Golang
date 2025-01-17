package entities

import (
	"time"
)

type HistoryKnowledgeList struct {
	ID                 int       `json:"id"`
	KnowledgeContentID int       `json:"knowledge_content_id"`
	Note               string    `json:"note"`
	Type               string    `json:"type"`
	Value              string    `json:"value"`
	Status             string    `json:"status"`
	Date               time.Time `json:"date"`
	Requestor          string    `json:"requestor"`
	UpdatedAt          time.Time `json:"updated_at"`
	UpdatedBy          string    `json:"updated_by"`
	CreatedAt          time.Time `json:"created_at"`
	CreatedBy          int       `json:"created_by"`
	DeletedAt          time.Time `json:"deleted_at"`
	Deleted            int       `json:"deleted"`
}

type HistoryNotif struct {
	ID    int    `json:"id"`
	Notif string `json:"notif"`
}

func (HistoryNotif) TableName() string {
	return "history_knowledge"
}

func (HistoryKnowledgeList) TableName() string {
	return "history_knowledge"
}

type HistoryKnowledgePreview struct {
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
