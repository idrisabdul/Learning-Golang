package options

import (
	"sygap_new_knowledge_management/backend/entities"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductTypeRepo struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewProductTypeRepos(db *gorm.DB, log *logrus.Logger) *ProductTypeRepo {
	return &ProductTypeRepo{db, log}
}

func (r *ProductTypeRepo) GetDetailProductType(IDProductName int) ([]entities.ListServiceType, error) {
	r.log.Println("Execute function GetDetailProductType")
	var listProductType []entities.ListServiceType

	err := r.db.Table("product_categories pn").Select(`
			st.id, st.service_type_name as service_type_name, st.status, st.type, st.id_service
	`).
		Joins("LEFT JOIN service_type as st ON pn.id_service = st.id_service").
		Where("pn.id_service = ?", IDProductName).
		Group("st.id").
		Order("st.service_type_name asc").
		Find(&listProductType).Error
	if err != nil {
		r.log.Error("Failed to get list product type in GetDetailProductType", err)
		return nil, err
	}

	return listProductType, nil

}

func (r *ProductTypeRepo) GetAllDetailProductType(IDProductName int) ([]entities.ListServiceType, error) {
	r.log.Println("Execute function GetAllDetailProductType")
	var listProductType []entities.ListServiceType

	err := r.db.Table("product_categories pn").Select(`
			st.id, st.service_type_name as service_type_name, st.status, st.type, st.id_service
	`).
		Joins("LEFT JOIN service_type as st ON pn.id_service = st.id_service").
		Where("st.id IS NOT NULL").
		Group("st.id").
		Order("st.service_type_name asc").
		Find(&listProductType).Error
	if err != nil {
		r.log.Error("Failed to get list product type in GetAllDetailProductType", err)
		return nil, err
	}

	return listProductType, nil

}

func (r *ProductTypeRepo) GetDetailProductTypeRelation(IDModule int, search string) ([]entities.ListServiceType, error) {
	r.log.Println("Execute function GetDetailProductType")
	var listProductType []entities.ListServiceType

	if search != "" {
		err := r.db.Table("services_has_module shm").Select(`
		st.id, shm.id_service_type, CONCAT(st.service_type_name, ' - '  ,pc.name) as service_type_name, st.status, st.id_service
`).
			Joins("LEFT JOIN service_type as st ON st.id = shm.id_service_type").
			Joins("LEFT JOIN product_categories pc ON st.id_service = pc.id_service").
			Where("shm.id_module = ? AND st.service_type_name LIKE ? COLLATE utf8_general_ci", IDModule, "%"+search+"%").
			Group("st.id").
			Find(&listProductType).Error

		if err != nil {
			r.log.Error("Failed to get list product type in GetDetailProductType", err)
			return nil, err
		}
	} else {
		err := r.db.Table("services_has_module shm").Select(`
		st.id, shm.id_service_type, CONCAT(st.service_type_name, ' - '  ,pc.name) as service_type_name, st.status, st.id_service
`).
			Joins("LEFT JOIN service_type as st ON st.id = shm.id_service_type").
			Joins("LEFT JOIN product_categories pc ON st.id_service = pc.id_service").
			Where("shm.id_module = ?", IDModule).
			Group("st.id").
			Find(&listProductType).Error

		if err != nil {
			r.log.Error("Failed to get list product type in GetDetailProductType", err)
			return nil, err
		}

	}

	return listProductType, nil

}

func (r *ProductTypeRepo) GetAllDetailProductTypeRelation(search string) ([]entities.ListServiceType, error) {
	r.log.Println("Execute function GetDetailProductType")
	var listProductType []entities.ListServiceType

	if search != "" {
		err := r.db.Table("services_has_module shm").Select(`
		st.id, shm.id_service_type, CONCAT(st.service_type_name, ' - '  ,pc.name) as service_type_name, st.status, st.id_service
`).
			Joins("LEFT JOIN service_type as st ON st.id = shm.id_service_type").
			Joins("LEFT JOIN product_categories pc ON st.id_service = pc.id_service").
			Where("st.service_type_name LIKE ? COLLATE utf8_general_ci", "%"+search+"%").
			Group("st.id").
			Find(&listProductType).Error

		if err != nil {
			r.log.Error("Failed to get list product type in GetDetailProductType", err)
			return nil, err
		}
	} else {
		err := r.db.Table("services_has_module shm").Select(`
		st.id, shm.id_service_type, CONCAT(st.service_type_name, ' - ' ,pc.name) as service_type_name, st.status, st.id_service
`).
			Joins("LEFT JOIN service_type as st ON st.id = shm.id_service_type").
			Joins("LEFT JOIN product_categories pc ON st.id_service = pc.id_service").
			Group("st.id").
			Find(&listProductType).Error

		if err != nil {
			r.log.Error("Failed to get list product type in GetDetailProductType", err)
			return nil, err
		}

	}

	return listProductType, nil

}
