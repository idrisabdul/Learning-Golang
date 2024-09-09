package updaterequest

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"sygap_new_knowledge_management/backend/entities"
	updaterequest "sygap_new_knowledge_management/backend/repository/update-request"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SubmitService struct {
	repo *updaterequest.SubmitRepos
	log  *logrus.Logger
}

func NewSubmitService(repo *updaterequest.SubmitRepos, log *logrus.Logger) *SubmitService {
	return &SubmitService{repo, log}
}

func (s *SubmitService) SubmitUpdateRequest(data entities.SubmitUpdateRequest) error {

	knowledge_content_update_request := entities.KnowledgeContentUpdateRequest{
		KnowledgeID:       data.KnowledgeID,
		UpdateRequestType: data.UpdateRequestType,
		Status:            data.Status,
		ArticleVersion:    data.ArticleVersion,
		RequestSummary:    data.RequestSummary,
		RequestDetail:     data.RequestDetail,
		SubmitterID:       data.Submitter,
		SubmitDate:        data.SubmitDate,
	}

	if strings.EqualFold(data.Status, "in progress") {
		knowledge_content_update_request.ActualStartDate = utils.ConvertStringToTime(utils.GetTimeNow("normal"))
	}

	IDUpdateRequest, errSubmitToKnowledgeContentUpdateRequest := s.repo.SubmitToKnowledgeContentUpdateRequest(knowledge_content_update_request)
	if errSubmitToKnowledgeContentUpdateRequest != nil {
		return errSubmitToKnowledgeContentUpdateRequest
	}

	if len(data.Attachment) != 0 {
		var files []entities.KnowledgeContentReportAttachment
		for k, v := range data.Attachment {
			fileExt := strings.TrimPrefix(filepath.Ext(v.Filename), ".")
			DateTime := utils.GetTimeNow("normal")

			_, errUpload := utils.UploadFile(&fiber.Ctx{}, v, FileName(strconv.Itoa(k), strconv.Itoa(data.KnowledgeID), DateTime, fileExt))
			if errUpload != nil {
				continue
			}

			files = append(files, entities.KnowledgeContentReportAttachment{
				KnowledgeContentID:              data.KnowledgeID,
				Size:                            strconv.Itoa(int(v.Size)),
				Attachment:                      v.Filename,
				Filename:                        FileName(strconv.Itoa(k), strconv.Itoa(data.KnowledgeID), DateTime, fileExt),
				KnowledgeContentUpdateRequestID: IDUpdateRequest,
			})
		}

		if errSubmitToContentReportAttachment := s.repo.SubmitToKnowledgeUpdateReportAttachment(files); errSubmitToContentReportAttachment != nil {
			return errSubmitToContentReportAttachment
		}
	}

	return nil
}

func FileName(index, idKM, date, ext string) string {
	index = utils.GenerateEncoded(index)
	idKM = utils.GenerateEncoded(idKM)
	return fmt.Sprintf("Update-Request-%v-%v-%v.%v", index, idKM, date, ext)
}
