package entities

import "time"

type Feedback struct {
	ID          		int       `gorm:"primaryKey;autoIncrement" json:"id"`
	KnowledgeID    		int       `json:"knowledge_id"`
	SubmitterID         int    `json:"submitter_id"`
	Usefull        		string    `json:"usefull"`
	Rating      		int    `json:"rating"`
	Comment     		string    `json:"comment"`
	DateSubmit  		time.Time    `json:"date_submit"`
}

type FeedbackList struct {
	ID          		int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Code    			string       `json:"code"`
	Submitter       	string    `json:"submitter"`
	Usefull        		string    `json:"usefull"`
	Rating      		int    `json:"rating"`
	Comment     		string    `json:"comment"`
	DateSubmit  		*time.Time    `json:"date_submit"`
}

func (FeedbackList) TableName() string {
	return "knowledge_content_feedback"
}
