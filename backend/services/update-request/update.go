package updaterequest

import (
	"path/filepath"
	"strconv"
	"strings"
	"sygap_new_knowledge_management/backend/entities"
	updaterequest "sygap_new_knowledge_management/backend/repository/update-request"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UpdateService struct {
	repo *updaterequest.UpdateRepos
	log  *logrus.Logger
}

func NewUpdateService(repo *updaterequest.UpdateRepos, log *logrus.Logger) *UpdateService {
	return &UpdateService{repo, log}
}

func (s *UpdateService) UpdateUpdateRequest(data entities.SubmitUpdateRequest) error {

	knowledge_content_update_request := entities.KnowledgeContentUpdateRequest{
		ID:                data.ID,
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

	if strings.EqualFold(data.Status, "completed") {
		knowledge_content_update_request.ActualEndDate = utils.ConvertStringToTime(utils.GetTimeNow("normal"))
	}

	if errUpdateToKnowledgeContentUpdateRequest := s.repo.UpdateToKnowledgeContentUpdateRequest(knowledge_content_update_request); errUpdateToKnowledgeContentUpdateRequest != nil {
		return errUpdateToKnowledgeContentUpdateRequest
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
				KnowledgeContentUpdateRequestID: data.ID,
			})

			if errSubmitToContentReportAttachment := s.repo.InstanceSubmitToKnowledgeContentReportAttachment(files); errSubmitToContentReportAttachment != nil {
				return errSubmitToContentReportAttachment
			}
		}
	}
	return nil
}

func (s *UpdateService) DeleteUpdateRequest(id, author int) error {

	knowledge_content_update_request := entities.KnowledgeContentUpdateRequest{
		ID:        id,
		DeletedAt: utils.ConvertStringToTime(utils.GetTimeNow("normal")),
		DeletedBy: author,
	}

	if errUpdateToKnowledgeContentUpdateRequest := s.repo.UpdateToKnowledgeContentUpdateRequest(knowledge_content_update_request); errUpdateToKnowledgeContentUpdateRequest != nil {
		return errUpdateToKnowledgeContentUpdateRequest
	}

	return nil
}

func (s *UpdateService) DeleteUpdateRequestAttachment(idFile int, author string) error {

	knowledge_content_report_attachement := entities.KnowledgeContentReportAttachment{
		ID:        idFile,
		DeletedAt: utils.ConvertStringToTime(utils.GetTimeNow("normal")),
		DeletedBy: author,
	}

	if errUpdateToKnowledgeContentReportAttachment := s.repo.UpdateToKnowledgeContentReportAttachment(knowledge_content_report_attachement); errUpdateToKnowledgeContentReportAttachment != nil {
		return errUpdateToKnowledgeContentReportAttachment
	}

	return nil
}
