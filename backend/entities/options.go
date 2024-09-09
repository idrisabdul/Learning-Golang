package entities

import (
	"sygap_new_knowledge_management/backend/utils"

	"gorm.io/gorm"
)

type (
	OperationCategory struct {
		ID                  int    `json:"id"`
		OperationalCategory string `json:"operational_category"`
	}

	ServiceClass struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	StatusChange struct {
		ID     int    `json:"id"`
		Status string `json:"status"`
	}

	Product struct {
		ID         string `json:"id"`
		ParentID   string `json:"parent_id"`
		Name       string `json:"name"`
		ParentName string `json:"parent_name"`
	}

	ListProduct struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Category string `json:"category"`
	}

	ListServiceType struct {
		ID              int     `gorm:"primaryKey;autoIncrement" json:"id"`
		ServiceTypeName *string `json:"service_type_name"`
		Status          *string `json:"status"`
		Type            *string `json:"type"`
		IDService       *int    `json:"id_service"`
	}

	Company struct {
		ID          string `json:"id"`
		CompanyName string `json:"company_name"`
	}

	Organization struct {
		ID               string `json:"id"`
		OrganizationName string `json:"organization_name"`
		IDCompany        string `json:"id_company"`
	}

	ListExpertees struct {
		ID           string `json:"id"`
		Username     string `json:"username"`
		EmployeeName string `json:"employee_name"`
		// IDCompany        string `json:"id_company"`
		IDRole string `json:"id_role,omitempty"`
		// OrganizationName string `json:"group"`
	}

	ContentType struct {
		ID          int    `json:"id"`
		ContentName string `json:"content_name"`
	}

	ListupdateRequestStatus struct {
		ID     int    `json:"id"`
		Status string `json:"status"`
	}

	ListWorkDetailType struct {
		ID      int              `json:"-"`
		Label   string           `json:"label"`
		Options []WorkDetailType `json:"options"`
	}

	WorkDetailType struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	}
)

func (p *Product) SetParentName(db *gorm.DB) error {
	if p.ParentID == "0" {
		p.ParentName = ""
		return nil
	}

	var parentCategory2, parentCategory3 Product

	if errParentCategory3 := db.Table(utils.TABLE_PRODUCT_CATEGORIES).
		Where("pc.id = ?", p.ParentID).
		Find(&parentCategory3).Error; errParentCategory3 != nil {
		return errParentCategory3
	}
	if errParentCategory2 := db.Table(utils.TABLE_PRODUCT_CATEGORIES).
		Where("pc.id = ?", parentCategory3.ParentID).
		Find(&parentCategory2).Error; errParentCategory2 != nil {
		return errParentCategory2
	}

	p.Name = parentCategory3.Name
	p.ParentName = parentCategory2.Name

	return nil
}
