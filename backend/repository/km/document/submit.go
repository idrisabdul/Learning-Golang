package document

import (
	"strconv"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/utils"

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

func (r *SubmitRepos) SubmitToKnowledgeContentAttachment(data []entities.KnowledgeContentAttachment) ([]string, error) {
	r.log.Print("Execute SubmitToKnowledgeContentAttachment Function")

	if errSubmitToKnowledgeContentAttachment := r.db.CreateInBatches(&data, len(data)).Error; errSubmitToKnowledgeContentAttachment != nil {
		r.log.Errorln("Failed to SubmitToKnowledgeContentAttachment: ", errSubmitToKnowledgeContentAttachment)
		return nil, errSubmitToKnowledgeContentAttachment
	}

	var IDFile []string

	for _, v := range data {
		IDFile = append(IDFile, utils.GenerateNumberEncode(strconv.Itoa(v.ID)))
	}

	return IDFile, nil
}
