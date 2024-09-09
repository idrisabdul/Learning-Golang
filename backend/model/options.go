package model

type ListOrganization struct {
	ID               int    `gorm:"primaryKey;autoIncrement" json:"id"`
	OrganizationName string `gorm:"size:255;default:null" json:"organization_name"`
	OrganizationCode string `gorm:"size:45;default:null" json:"organization_code"`
}

type ListStatus struct {
	Status string `json:"status"`
}

type ListSymptoms struct {
	ID            int     `gorm:"primaryKey;autoIncrement" json:"id"`
	Kolom2        int     `json:"kolom_2"`
	IDService     *string `json:"id_service"`
	SymptomName   *string `json:"symptom_name"`
	Description   *string `json:"description"`
	IDServiceType *int    `json:"id_service_type"`
}

type ListRelationType struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Type string `json:"type"`
}