package model

type ResponseWorkDetailType struct {
	Label  string      `json:"label"`
	Option interface{} `json:"option"`
}

type AddWorkDetail struct {
	Type         string `json:"type" validate:"required"`
	Notes        string `json:"notes"`
	SubmitDate   string `json:"submit_date"  validate:"required"`
	Receiver     string `json:"receiver"`
	Notification string `json:"notification" validate:"required"`
	RequestType  string `json:"request_type"`
}
