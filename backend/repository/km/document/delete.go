package document

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DeleteDocument struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewDeleteDocument(db *gorm.DB, log *logrus.Logger) *DeleteDocument {
	return &DeleteDocument{db, log}
}

func (r *DeleteDocument) DeleteDocumentKM(idFile, author string) error {
	r.log.Print("Execute DeleteDocument in Repo")

	if errDeleteDocument := r.db.Table(utils.RemoveAliasFromTable(utils.TABLE_KNOWLEDGE_CONTENT_ATTACHMENT)).
		Where("id = ?", idFile).
		Updates(&entities.KnowledgeContentAttachment{
			DeletedAt: utils.ConvertStringToTime(utils.GetTimeNow("normal")),
			DeletedBy: author,
		}).Error; errDeleteDocument != nil {
		r.log.Errorln("Failed Execute DeleteDocument in Repo: ", errDeleteDocument)
		return errDeleteDocument
	}

	return nil
}
