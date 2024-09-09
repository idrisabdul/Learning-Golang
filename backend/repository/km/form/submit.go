package form

import (
	"sygap_new_knowledge_management/backend/entities"
	workdetail "sygap_new_knowledge_management/backend/repository/work-detail"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SubmitRepos struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewSubmitRepos(db *gorm.DB, log *logrus.Logger) *SubmitRepos {
	return &SubmitRepos{db, log}
}

func (r *SubmitRepos) SubmitToKnowledgeContent(data *entities.KnowledgeContent) (int, error) {
	r.log.Print("Execute SubmitToKnowledgeContent Function")

	var response entities.KnowledgeContent

	if errSubmitToKnowledgeContent := r.db.Create(&data).Error; errSubmitToKnowledgeContent != nil {
		r.log.Errorln("Failed to SubmitToKnowledgeContent: ", errSubmitToKnowledgeContent.Error())
		return 0, errSubmitToKnowledgeContent
	}

	if errReturnDataKnowledgeContent := r.db.Find(&response, data.ID).Error; errReturnDataKnowledgeContent != nil {
		r.log.Errorln("Failed to ReturnDataKnowledgeContent: ", errReturnDataKnowledgeContent.Error())
		return 0, errReturnDataKnowledgeContent
	}

	return response.ID, nil
}

func (r *SubmitRepos) SubmitToKnowledgeContentDetail(data *entities.KnowledgeContentDetail) error {
	r.log.Print("Execute SubmitToKnowledgeContentDetail Function")

	if errSubmitToKnowledgeContentDetail := r.db.Create(&data).Error; errSubmitToKnowledgeContentDetail != nil {
		r.log.Errorln("Failed to SubmitToKnowledgeContentDetail: ", errSubmitToKnowledgeContentDetail.Error())
		return errSubmitToKnowledgeContentDetail
	}

	return nil
}

func (r *SubmitRepos) SubmitToKnowledgeContentLog(data entities.KnowledgeContentLog) error {
	r.log.Print("Execute SubmitToKnowledgeContentLog Function")

	if errSubmitToKnowledgeContentLog := r.db.Create(&data).Error; errSubmitToKnowledgeContentLog != nil {
		r.log.Errorln("Failed to SubmitToKnowledgeContentLog: ", errSubmitToKnowledgeContentLog)
		return errSubmitToKnowledgeContentLog
	}

	return nil
}

func (r *SubmitRepos) SubmitToKnowledgeContentOption(data []entities.KnowledgeContentOption) error {
	r.log.Print("Execute SubmitToKnowledgeContentKeyword Function")

	if errSubmitToKnowledgeContentOption := r.db.CreateInBatches(&data, len(data)).Error; errSubmitToKnowledgeContentOption != nil {
		r.log.Println("Failed to SubmitToKnowledgeContentOption: ", errSubmitToKnowledgeContentOption.Error())
		return errSubmitToKnowledgeContentOption
	}

	return nil
}

func (r *SubmitRepos) SubmitToKnowledgeContentQuestion(data []entities.KnowledgeContentQuestion) ([]int, error) {
	r.log.Print("Execute SubmitToKnowledgeContentKeyword Function")

	var temporal []int

	if errSubmitToKnowledgeContentOption := r.db.CreateInBatches(&data, len(data)).Error; errSubmitToKnowledgeContentOption != nil {
		r.log.Println("Failed to SubmitToKnowledgeContentQuestion: ", errSubmitToKnowledgeContentOption.Error())
		return nil, errSubmitToKnowledgeContentOption
	}

	for _, v := range data {
		temporal = append(temporal, v.ID)
	}

	return temporal, nil
}

func (r *SubmitRepos) SubmitToKnowledgeContentLogVersion(data entities.KnowledgeContentLogVersion) error {

	if errSubmitToKnowledgeContentLogVersion := r.db.Create(&data).Error; errSubmitToKnowledgeContentLogVersion != nil {
		return errSubmitToKnowledgeContentLogVersion
	}

	return nil
}

func (r *SubmitRepos) InstanceSubmitToWorkdetailKnowledgeManagement(data entities.WorkdetailKnowledgeManagement) (int, error) {
	return workdetail.NewWorkDetail(r.db, r.log).SubmitToWorkdetailKnowledgeManagement(data)
}

func (r *SubmitRepos) InstanceSubmitToKnowledgeRelationToTicket(data []entities.KnowledgeRelationToTicketPopup) error {
	if err := r.db.CreateInBatches(&data, len(data)).Error; err != nil {
		r.log.Println("Failed Execute function Save Knowledge Relation")
		return err
	}
	return nil
}

func (r *SubmitRepos) InstanceDetailLogVersion(data int) ([]entities.KnowledgeContentLogVersion, error) {
	return NewDetailRepos(r.db, r.log).DetailLogVersion(data)
}
