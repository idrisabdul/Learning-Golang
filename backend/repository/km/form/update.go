package form

import (
	"sygap_new_knowledge_management/backend/entities"
	workdetail "sygap_new_knowledge_management/backend/repository/work-detail"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UpdateRepos struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewUpdateRepos(db *gorm.DB, log *logrus.Logger) *UpdateRepos {
	return &UpdateRepos{db, log}
}

func (r *UpdateRepos) UpdateToKnowledgeContent(data entities.KnowledgeContent) error {
	r.log.Print("Execute UpdateToKnowledgeContent Function")

	if errUpdateToKnowledgeContent := r.db.Where("id = ?", data.ID).Updates(&data).Error; errUpdateToKnowledgeContent != nil {
		r.log.Println("Failed Execute UpdateToKnowledgeContent Function: ", errUpdateToKnowledgeContent.Error())
		return errUpdateToKnowledgeContent
	}

	return nil
}

func (r *UpdateRepos) UpdateToKnowledgeContentDetail(data entities.KnowledgeContentDetail) error {
	r.log.Print("Execute UpdateToKnowledgeContentDetail Function")

	// if errUpdateToKnowledgeContentDetail := r.db.Where("knowledge_content_id = ?", data.KnowledgeContentID).Save(&data).Error; errUpdateToKnowledgeContentDetail != nil {
	// 	r.log.Errorln("Failed to UpdateToKnowledgeContentDetail: ", errUpdateToKnowledgeContentDetail.Error())
	// 	return errUpdateToKnowledgeContentDetail
	// }

	convertedData, _ := utils.ConvertStructToMap(data)

	if errUpdateToKnowledgeContentDetail := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "knowledge_content_id"}},
		DoUpdates: clause.Assignments(convertedData),
	}).Save(&data).Error; errUpdateToKnowledgeContentDetail != nil {
		r.log.Errorln("Failed to UpdateToKnowledgeContentDetail: ", errUpdateToKnowledgeContentDetail.Error())
		return errUpdateToKnowledgeContentDetail
	}

	return nil
}

func (r *UpdateRepos) UpdateToKnowledgeContentOption(data entities.KnowledgeContentOption) error {
	r.log.Print("Execute UpdateToKnowledgeContentKeyword Function")

	if errUpdateToKnowledgeContentOption := r.db.Where("id = ?", data.ID).Updates(&data).Error; errUpdateToKnowledgeContentOption != nil {
		r.log.Errorln("Failed to UpdateToKnowledgeContentOption: ", errUpdateToKnowledgeContentOption.Error())
		return errUpdateToKnowledgeContentOption
	}

	return nil
}

func (r *UpdateRepos) UpdateToKnowledgeContentQuestion(data entities.KnowledgeContentQuestion) error {
	r.log.Print("Execute UpdateToKnowledgeContentKeyword Function")

	if errUpdateToKnowledgeContentOption := r.db.Where("id = ?", data.ID).Updates(&data).Error; errUpdateToKnowledgeContentOption != nil {
		r.log.Errorln("Failed to UpdateToKnowledgeContentQuestion: ", errUpdateToKnowledgeContentOption.Error())
		return errUpdateToKnowledgeContentOption
	}

	return nil
}

func (r *UpdateRepos) InstanceSubmitToKnowledgeContentQuestion(data []entities.KnowledgeContentQuestion) ([]int, error) {
	return NewSubmitRepos(r.db, r.log).SubmitToKnowledgeContentQuestion(data)
}

func (r *UpdateRepos) InstanceSubmitToKnowledgeContentOption(data []entities.KnowledgeContentOption) error {
	return NewSubmitRepos(r.db, r.log).SubmitToKnowledgeContentOption(data)
}

func (r *UpdateRepos) InstanceSubmitToKnowledgeContentLog(data entities.KnowledgeContentLog) error {
	return NewSubmitRepos(r.db, r.log).SubmitToKnowledgeContentLog(data)
}

func (r *UpdateRepos) InstanceSubmitToWorkdetailKnowledgeManagement(data entities.WorkdetailKnowledgeManagement) (int, error) {
	return workdetail.NewWorkDetail(r.db, r.log).SubmitToWorkdetailKnowledgeManagement(data)
}
