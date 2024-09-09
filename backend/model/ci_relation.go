package model

import "time"

type CiRelationTypeFromCiClass struct {
	ID           int    `json:"id"`
	IDSuperClass int    `json:"id_super_class"`
	ClassName    string `json:"class_name"`
}

type CiNameOption struct {
	ID            int     `gorm:"primaryKey;autoIncrement" json:"id"`
	CIClassID     *int    `json:"ci_class_id"`
	CINameCombine *string `json:"ci_name"`
	SerialNumber  *string `json:"serial_number"`
}

type CiHistoryOption struct {
	ID          int    `json:"id"`
	HistoryName string `json:"history_name"`
	Value       string `json:"-"`
}

type CiHistoryResponse struct {
	ID          int                    `json:"id"`
	HistoryName string                 `json:"history_name"`
	Log         map[string]interface{} `json:"log"`
}

type CreateCiRelation struct {
	CiTypeID      int     `json:"ci_type_id"`
	CiNameID      int     `json:"ci_name_id"`
	RelationType  string  `json:"relation_type"`
	HistoryID     *int    `json:"history_id"`
	CreatedAt     string  `json:"created_at"`
	CreatedBy     int     `json:"created_by"`
	UpdatedAt     *string `json:"updated_at"`
	UpdatedBy     *int    `json:"updated_by"`
	Deleted       *int    `json:"deleted"`
	DeletedAt     *string `json:"deleted_at"`
	IDRequestType int     `json:"id_request_type"`
	RequestType   string  `json:"request_type"`
}

type CiRelationSubmit struct {
	CiTypeID      int    `json:"ci_type_id" validate:"required"`
	CiNameID      int    `json:"ci_name_id" validate:"required"`
	RelationType  string `json:"relation_type" validate:"required"`
	HistoryID     *int   `json:"history_id"`
	RequestType   string `json:"request_type"`
	IDRequestType int    `json:"id_request_type"`
}

type CiRelationUpdate struct {
	CiTypeID      int    `json:"ci_type_id" validate:"required"`
	CiNameID      int    `json:"ci_name_id" validate:"required"`
	RelationType  string `json:"relation_type" validate:"required"`
	HistoryID     *int   `json:"history_id"`
	RequestType   string `json:"request_type"`
	IDRequestType int    `json:"id_request_type"`
}

type CiRelationList struct {
	ID            int        `json:"id"`
	CiTypeID      int        `json:"ci_type_id"`
	CiTypeName    string     `json:"ci_type_name"`
	CiNameID      int        `json:"ci_name_id"`
	CiName        string     `json:"ci_name"`
	RelationType  string     `json:"relation_type"`
	AttributeName string     `json:"attribute_name"`
	HistoryID     int        `json:"history_id"`
	HistoryDate   *time.Time `json:"history_date"`
	CreatedAt     time.Time  `json:"created_at"`
	Parent        string     `json:"parent"`
}

type CiAttributeNameOption struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
