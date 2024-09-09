package document

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

func (r *DetailRepos) DetailDocumentKM(idKM string) ([]entities.KnowledgeContentAttachment, error) {
	r.log.Print("Execute DetailDocumentKM in Repo")

	var response []entities.KnowledgeContentAttachment

	if errDetailDocumentKM := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_ATTACHMENT).
		Where("kca.knowledge_content_id = ?", idKM).
		Where("kca.deleted_at IS NULL").
		Find(&response).Error; errDetailDocumentKM != nil {
		return nil, errDetailDocumentKM
	}

	return response, nil
}
