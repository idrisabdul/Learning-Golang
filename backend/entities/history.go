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
