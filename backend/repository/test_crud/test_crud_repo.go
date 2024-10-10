package testcrud

import (
	"sygap_new_knowledge_management/backend/model"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TestCrudRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewTestCrudRepository(db *gorm.DB, log *logrus.Logger) *TestCrudRepository {
	return &TestCrudRepository{db, log}
}

func (r *TestCrudRepository) GetTestCrud() ([]model.ListCrudModel, error) {
	var ListCrud []model.ListCrudModel

	query := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT).
		Select(`knowledge_id,
	kc.title,
	kc.keyword,
	kc.status,
	e.employee_name  AS created_by ,
	kc.company_id,
	c.company_name AS company`).
		Joins(`LEFT JOIN employee e ON e.id = kc.created_by`).
		Joins(`LEFT JOIN company c ON c.id = kc.company_id`)

	if err := query.Find(&ListCrud).Error; err != nil {
		r.log.Error("Failed get db crud")
	}

	return ListCrud, nil
}

func (r *TestCrudRepository) GetDetailCrud(idKnowledgeContent int) (model.GetDetailCrudModel, error) {
	var GetDetail model.GetDetailCrudModel

	query := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT).
		Select(`knowledge_id,
	kc.title,
	kc.keyword,
	kc.status,
	e.employee_name AS created_by ,
	kc.company_id,
	c.company_name AS company`).
		Joins(`LEFT JOIN employee e ON e.id = kc.created_by`).
		Joins(`LEFT JOIN company c ON c.id = kc.company_id`).
		Where(`kc.id = ?`, idKnowledgeContent)

	if err := query.Find(&GetDetail).Error; err != nil {
		r.log.Errorln("Error on GetListUpdateRequest in Repo: ", err)
	}

	return GetDetail, nil

}

func (r *TestCrudRepository) InsertCrudTest(data model.AddKnowledgeContent) (model.AddKnowledgeContent, error) {
	query := r.db.Table("knowledge_content").Create(&data)

	if errInsert := query.Error; errInsert != nil {
		r.log.Errorln("Error on InsertCrudTest in Repo: ", errInsert)
	}

	return data, nil
}

func (r *TestCrudRepository) UpdateCrudTest(data model.UpdateKnowledgeContent) (model.UpdateKnowledgeContent, error) {

	query := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT).Where(`id = ?`, data.ID).Updates(&data)

	if errQuery := query.Error; errQuery != nil {
		r.log.Println("Error on UpdateCrudTest in Repo: ", errQuery)
	}

	return data, nil
}
