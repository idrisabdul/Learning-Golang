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

func (r *UpdateRepos) InstanceSubmitToKnowledgeContentLog(data entities.KnowledgeContentLog) error {
	return NewSubmitRepos(r.db, r.log).SubmitToKnowledgeContentLog(data)
}

func (r *UpdateRepos) InstanceSubmitToWorkdetailKnowledgeManagement(data entities.WorkdetailKnowledgeManagement) (int, error) {
	return workdetail.NewWorkDetail(r.db, r.log).SubmitToWorkdetailKnowledgeManagement(data)
}

func (r *UpdateRepos) DeleteQuestionAndOptions(idKM int) error {
	r.log.Print("Execute delete before update decision tree KM")
	var content entities.KnowledgeContentOption

	if errDeleteDocument := r.db.Table(utils.RemoveAliasFromTable(utils.TABLE_KNOWLEDGE_CONTENT_OPTION)).
		Where("knowledge_content_id = ?", idKM).Delete(&content).
		Error; errDeleteDocument != nil {
		r.log.Errorln("Failed Execute DeleteDocument in Repo: ", errDeleteDocument)
		return errDeleteDocument
	}
	return nil
}

func (r *UpdateRepos) UpdateQuestionAndOptions(question entities.Option, idKM int) error {
	if err := UpdateQuestionAndOptionsRecursive(r.db, question, nil, idKM); err != nil {
		return err
	}
	return nil
}

func UpdateQuestionAndOptionsRecursive(db *gorm.DB, question entities.Option, parentOptionID *int, idKM int) error {
	optionData := entities.KnowledgeContentOption{
		Label:              question.Label,
		Solution:           question.Answer,
		Question:           question.Question,
		OptionParentId:     parentOptionID,
		KnowledgeContentID: idKM,
	}
	err := db.Create(&optionData).Error
	if err != nil {
		return err
	}
	idOption := optionData.ID

	for _, childOption := range question.Options {
		err = SaveQuestionAndOptionsRecursive(db, childOption, &idOption, idKM)
		if err != nil {
			return err
		}
	}

	return nil
}
