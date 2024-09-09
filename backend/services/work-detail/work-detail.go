package workdetail

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"sygap_new_knowledge_management/backend/entities"
	workdetail "sygap_new_knowledge_management/backend/repository/work-detail"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type WorkDetailService struct {
	repo *workdetail.WorkDetailRepository
	log  *logrus.Logger
}

func NewWorkDetailService(repo *workdetail.WorkDetailRepository, log *logrus.Logger) *WorkDetailService {
	return &WorkDetailService{repo, log}
}

func (s *WorkDetailService) GetWorkDetail(idKM string) ([]entities.WorkDetailChangeResponse, error) {
	return s.repo.GetWorkDetail(idKM)
}

func (s *WorkDetailService) SubmitWorkDetail(data entities.SubmitWorkDetail) error {
	workdetail_knowledge_management := entities.WorkdetailKnowledgeManagement{
		IDParent:  data.IDParent,
		Type:      data.Type,
		Note:      data.Note,
		CreatedBy: data.Submitter,
	}

	IDWorkDetail, errSubmitToWorkdetailKnowledgeManagement := s.repo.SubmitToWorkdetailKnowledgeManagement(workdetail_knowledge_management)
	if errSubmitToWorkdetailKnowledgeManagement != nil {
		return errSubmitToWorkdetailKnowledgeManagement
	}

	if len(data.Attachment) != 0 {
		var files []entities.WorkDetailHasDocument
		for k, v := range data.Attachment {
			fileExt := strings.TrimPrefix(filepath.Ext(v.Filename), ".")
			fileName := strings.TrimSuffix(filepath.Base(v.Filename), filepath.Ext(v.Filename))
			DateTime := utils.GetTimeNow("normal")

			_, errUpload := utils.UploadFile(&fiber.Ctx{}, v, FileName(strconv.Itoa(k), strconv.Itoa(data.IDParent), DateTime, fileExt))
			if errUpload != nil {
				continue
			}

			files = append(files, entities.WorkDetailHasDocument{
				IDWorkDetail: IDWorkDetail,
				Type:         "knowledge",
				FileHash:     FileName(strconv.Itoa(k), strconv.Itoa(data.IDParent), DateTime, fileExt),
				FileOri:      v.Filename,
				FileType:     fileExt,
				FileSize:     strconv.Itoa(int(v.Size)),
				FileName:     fileName,
				CreatedAt:    utils.ConvertStringToTime(DateTime),
				CreatedBy:    data.Submitter,
			})
		}

		if errSubmitToContentReportAttachment := s.repo.SubmitToWorkDetailHasDocument(files); errSubmitToContentReportAttachment != nil {
			return errSubmitToContentReportAttachment
		}
	}
	return nil
}

func FileName(index, idKM, date, ext string) string {
	index = utils.GenerateEncoded(index)
	idKM = utils.GenerateEncoded(idKM)
	return fmt.Sprintf("Workdetail-KM-%v-%v-%v.%v", index, idKM, date, ext)
}
