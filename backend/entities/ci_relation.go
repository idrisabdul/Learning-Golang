package entities

import "time"

type CiType struct {
	ID        int       `json:"id"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

type CiTypeOption struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

func (CiTypeOption) TableName() string {
	return "ci_type"
}

type CiRelationType struct {
	ID        int       `json:"id"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

type CiRelationTypeOption struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

func (CiRelationTypeOption) TableName() string {
	return "ci_relation_type"
}
