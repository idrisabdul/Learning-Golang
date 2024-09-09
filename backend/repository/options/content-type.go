package options

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ContentTypeRepos struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewContentTypeRepos(db *gorm.DB, log *logrus.Logger) *ContentTypeRepos {
	return &ContentTypeRepos{db, log}
}

func (r *ContentTypeRepos) GetContentType() ([]entities.ContentType, error) {
	var contents []entities.ContentType

	if errGetContentType := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_LIST).Find(&contents).Error; errGetContentType != nil {
		r.log.Errorln("Error while running GetContentType Function on Repo: ", errGetContentType)
		return nil, errGetContentType
	}

	return contents, nil
}
