package document

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"strings"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/repository/km/document"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SubmitService struct {
	repo *document.SubmitRepos
	log  *logrus.Logger
}

func NewSubmitService(repo *document.SubmitRepos, log *logrus.Logger) *SubmitService {
	return &SubmitService{repo, log}
}

func (s *SubmitService) SubmitDocument(file []*multipart.FileHeader, idKM int) ([]string, error) {

	var files []entities.KnowledgeContentAttachment

	for k, v := range file {
		fileExt := strings.TrimPrefix(filepath.Ext(v.Filename), ".")
		DateTime := utils.GetTimeNow("normal")

		_, errUpload := utils.UploadFile(&fiber.Ctx{}, v, FileName(strconv.Itoa(k), strconv.Itoa(idKM), DateTime, fileExt))
		if errUpload != nil {
			return nil, errUpload
		}

		files = append(files, entities.KnowledgeContentAttachment{
			KnowledgeContentID: idKM,
			Size:               int(v.Size),
			Attachment:         v.Filename,
			Filename:           FileName(strconv.Itoa(k), strconv.Itoa(idKM), DateTime, fileExt),
		})
	}

	IDFiles, errSubmitToKnowledgeContentAttachment := s.repo.SubmitToKnowledgeContentAttachment(files)
	if errSubmitToKnowledgeContentAttachment != nil {
		return nil, errSubmitToKnowledgeContentAttachment
	}

	return IDFiles, nil
}

func FileName(index, idKM, date, ext string) string {
	index = utils.GenerateEncoded(index)
	idKM = utils.GenerateEncoded(idKM)
	return fmt.Sprintf("KM-Reference-%v-%v-%v.%v", index, idKM, date, ext)
}
