package form

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DetailRepos struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewDetailRepos(db *gorm.DB, log *logrus.Logger) *DetailRepos {
	return &DetailRepos{db, log}
}

func (r *DetailRepos) DetailKM(idKM string) (entities.DetailKM, error) {
	r.log.Print("Execute DetailKM in Repo")

	var response entities.DetailKM

	if errDetailKM := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT).
		Select(
			`kc.id,
			 kc.created_by,
			 kc.knowledge_id,
			 kc.knowledge_content_list_id,
			 kc.version,
			 kc.title,
			 kcd.question,
			 kcd.error,
			 kcd.root_cause,
			 kcd.workaround,
			 kcd.fix_solution,
			 kcd.technical_note,
			 kcd.reference,
			 kc.company_id,
			 kc.operation_category_1_id,
			 kc.operation_category_2_id,
			 kc.service_name_id,
			 kc.service_category_1_id,
			 kc.service_category_2_id,
			 kc.expert_group,
			 kc.expertee,
			 kc.status,
			 e.employee_name AS author,
			 kc.retire_date,
			 kc.published_date,
			 kc.keyword,
			 CASE 
			 WHEN kclv.key_content IS NOT NULL THEN kclv.key_content 
			 ELSE NULL END AS key_content
			 `).
		Joins("LEFT JOIN "+utils.TABLE_KNOWLEDGE_CONTENT_DETAIL+" ON kc.id = kcd.knowledge_content_id").
		Joins("LEFT JOIN "+utils.TABLE_EMPLOYEE+" ON e.id = kc.created_by").
		Joins("LEFT JOIN "+utils.TABLE_KNOWLEDGE_CONTENT_LOG_VERSION+" ON kc.id = kclv.id_knowledge_content").
		Where("kc.id = ?", idKM).Find(&response).Error; errDetailKM != nil {
		r.log.Errorln("Failed Execute DetailKM in Repo: ", errDetailKM.Error())
		return entities.DetailKM{}, errDetailKM
	}

	return response, nil
}

func (r *DetailRepos) DetailKMDecisionTree(idKM string) (entities.DetailKMDecisionTree, error) {
	r.log.Print("Execute DetailKMDecisionTree in Repo")

	var response entities.DetailKMDecisionTree

	if errDetailKMDecisionTree := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT).Select(
		`kc.id,
			 kc.knowledge_id,
			 kc.knowledge_content_list_id,
			 kc.version,
			 kc.title,
			 kc.company_id,
			 kc.operation_category_1_id,
			 kc.operation_category_2_id,
			 kc.service_name_id,
			 kc.service_category_1_id,
			 kc.service_category_2_id,
			 kc.expert_group,
			 kc.expertee,
			 kc.status,
			 e.employee_name AS author,
			 kc.retire_date,
			 kc.published_date,
			 kc.keyword,
			 kcd.question,
			 CASE 
			 WHEN kclv.key_content IS NOT NULL THEN kclv.key_content 
			 ELSE NULL END AS key_content
			 `).
		Joins("LEFT JOIN "+utils.TABLE_KNOWLEDGE_CONTENT_DETAIL+" ON kc.id = kcd.knowledge_content_id").
		Joins("LEFT JOIN "+utils.TABLE_EMPLOYEE+" ON e.id = kc.created_by").
		Joins("LEFT JOIN "+utils.TABLE_KNOWLEDGE_CONTENT_LOG_VERSION+" ON kc.id = kclv.id_knowledge_content").
		Where("kc.id = ?", idKM).Find(&response).Error; errDetailKMDecisionTree != nil {
		r.log.Errorln("Failed Execute DetailKMDecisionTree in Repo: ", errDetailKMDecisionTree.Error())
		return entities.DetailKMDecisionTree{}, nil
	}

	return response, nil
}

func (r *DetailRepos) GetQuestionDecisionTree(idkm string) ([]entities.KnowledgeContentQuestion, error) {
	r.log.Print("Execute GetQuestionDecisionTree in Repo")

	var questions []entities.KnowledgeContentQuestion

	if errQuestion := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_QUESTION).
		Where("kcq.knowledge_content_id = ?", idkm).
		Where("kcq.deleted_at IS NULL").
		Find(&questions).Error; errQuestion != nil {
		r.log.Errorln("Failed Execute GetQuestionDecisionTree in Repo: ", errQuestion.Error())
		return nil, errQuestion
	}

	return questions, nil
}

func (r *DetailRepos) GetOptionDecisionTree(idQuestion int) ([]entities.KnowledgeContentOption, error) {
	r.log.Print("Execute GetOptionDecisionTree in Repo")

	var option []entities.KnowledgeContentOption

	if errOption := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_OPTION).
		Where("kco.knowledge_content_question_id = ?", idQuestion).
		Where("kco.deleted_at IS NULL").
		Find(&option).Error; errOption != nil {
		r.log.Errorln("Failed Execute GetOptionDecisionTree in Repo: ", errOption)
		return nil, errOption
	}

	return option, nil
}

func (r *DetailRepos) DetailLogVersion(idKM int) ([]entities.KnowledgeContentLogVersion, error) {

	var list []entities.KnowledgeContentLogVersion

	if errDetailLogVersion := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_LOG_VERSION).
		Where("key_content = ?", idKM).
		Or("id_knowledge_content = ?", idKM).
		Find(&list).Error; errDetailLogVersion != nil {
		r.log.Errorln("Error while get list version log: ", errDetailLogVersion)
		return nil, errDetailLogVersion
	}

	return list, nil
}

func (r *DetailRepos) IsknowledgeManager(IDEmp int, knowledgeManagerID int) (bool, error) {
	var count int64
	err := r.db.Table("organization_employee_has_role").
		Where("id_employee = ? AND id_role = ?", IDEmp, knowledgeManagerID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}

}

func (r *DetailRepos) GetListOptionDecisionTree(idKM int) ([]entities.KnowledgeContentOption, error) {
	r.log.Print("Execute GetOptionDecisionTree in Repo")

	var option []entities.KnowledgeContentOption

	if errOption := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_OPTION).
		Where("kco.knowledge_content_id = ?", idKM).
		Find(&option).Error; errOption != nil {
		r.log.Errorln("Failed Execute GetOptionDecisionTree in Repo: ", errOption)
		return nil, errOption
	}

	return option, nil
}
