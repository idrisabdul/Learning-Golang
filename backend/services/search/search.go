package search

import (
	"mime/multipart"
	"strconv"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/model"
	"sygap_new_knowledge_management/backend/repository/search"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SearchService struct {
	repo *search.SearchRepos
	log  *logrus.Logger
}

func NewSearchService(repo *search.SearchRepos, log *logrus.Logger) *SearchService {
	return &SearchService{repo, log}
}

func (s *SearchService) GetSearchList(c *fiber.Ctx, user_id string) (interface{}, error) {
	data, err := s.repo.GetSearchList(c, user_id)
	return data, err
}

func (s *SearchService) GetContentDetail(c *fiber.Ctx, user_id string) (interface{}, error) {
	content_id, _ := utils.GenerateDecoded(c.Params("content_id"))
	data, err := s.repo.GetSearchDetail(content_id, user_id)
	s.GetUpsertVisitor(content_id, user_id)
	return data, err
}

func (s *SearchService) GetContentFeedback(c *fiber.Ctx, user_id string) (interface{}, error) {
	content_id, _ := utils.GenerateDecoded(c.Params("content_id"))
	data, err := s.repo.GetContentFeedback(content_id, user_id)
	return data, err
}

func (s *SearchService) ReportContentDetail(c *fiber.Ctx, user_id string, employee_name string) (interface{}, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}
	content_id, _ := utils.GenerateDecoded(c.Params("content_id"))
	content_id_int, _ := strconv.Atoi(content_id)
	user_id_int, _ := strconv.Atoi(user_id)
	// Notes & Save Report
	notes := form.Value["Notes"]
	if len(notes) == 0 {
		return []string{}, nil
	}
	idUpdateRequest, report_err := s.repo.CreateReportContent(entities.SearchContentReport{
		KnowledgeContentId: content_id_int,
		EmployeeId:         user_id_int,
		Notes:              notes[0],
	})
	if report_err != nil {
		return nil, report_err
	}

	// Attachment & Save Attachment
	var payload_attachment []entities.SearchContentAttachemntReport
	file_uploads := form.File["Attachment"]
	if len(file_uploads) > 0 {
		for i := range file_uploads {
			var filename string = ""
			var files *multipart.FileHeader = file_uploads[i]
			_, file_extension := utils.GetFileExtension(files.Filename)
			var counter string = utils.GetTimeNow("date_time")
			filename = "Knowledge-Management-Report" + utils.GenerateEncoded(c.Params("content_id")) + "_" + counter + "." + file_extension
			_, errs := utils.UploadFile(c, files, filename)
			s.log.Println(">>>>>>>>>>>>>>>> print <<<<<<<<<<<<<<<<", errs)
			payload_attachment = append(payload_attachment,
				entities.SearchContentAttachemntReport{
					KnowledgeContentId:              content_id_int,
					Attachment:                      files.Filename,
					Filename:                        filename,
					Size:                            int(files.Size),
					KnowledgeContentUpdateRequestID: idUpdateRequest,
				},
			)
		}
		if len(payload_attachment) > 0 {
			s.repo.CreateReportAttachmentContent(payload_attachment)
		}
	}

	// Send Email Notification
	detail, err := s.repo.GetMinimumDetail(content_id_int)
	if err == nil {
		utils.SendEmailNotification("reported", []string{
			detail.KnowledgeId, detail.Author, strconv.Itoa(detail.Version), employee_name,
			detail.Title,
		}, c)
	}
	return "Success to report this content", nil
}

func (s *SearchService) BookmarkContentDetail(c *fiber.Ctx, user_id string) (interface{}, error) {
	content_id, _ := utils.GenerateDecoded(c.Params("content_id"))
	content_id_int, _ := strconv.Atoi(content_id)
	return s.repo.BookmarkContentDetail(entities.KnowledgeContentBookmark{
		KnowledgeContentID: content_id_int,
		EmployeeId:         user_id,
	})
}

func (s *SearchService) FeedbackContentDetail(c *fiber.Ctx, user_id string, payload model.FeedbackContent, user_name string) (interface{}, error) {
	content_id, _ := utils.GenerateDecoded(c.Params("content_id"))
	content_id_int, _ := strconv.Atoi(content_id)
	user_id_int, _ := strconv.Atoi(user_id)
	_, err := s.repo.FeedbackContentDetail(entities.KnowledgeContentFeedback{
		KnowledgeId: content_id_int,
		SubmitterId: user_id_int,
		Usefull:     "",
		Rating:      payload.Rating,
		Comment:     payload.Comment,
	})
	// Send email notification
	if payload.Rating <= 3 {
		detail, err := s.repo.GetMinimumDetail(content_id_int)
		if err == nil {
			utils.SendEmailNotification("feedback", []string{
				detail.KnowledgeId, detail.Author, strconv.Itoa(detail.Version), strconv.Itoa(payload.Rating), user_name,
				detail.Title, payload.Comment,
			}, c)
		}
	}

	return "Success to feedback this content", err
}

func (s *SearchService) GetUpsertVisitor(contentId string, userId string) bool {
	user_id, _ := strconv.Atoi(userId)
	content_id, _ := strconv.Atoi(contentId)
	code := userId + "-" + contentId
	return s.repo.GetUpsertVisitor(content_id, user_id, code)
}
