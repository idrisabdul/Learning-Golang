package feedback_svc

import (
	"fmt"
	"strconv"
	repository "sygap_new_knowledge_management/backend/repository/feedback"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type FeedbackService struct {
	repo *repository.FeedbackRepository
	log  *logrus.Logger
}

func NewFeedbackService(repo *repository.FeedbackRepository, logger *logrus.Logger) *FeedbackService {
	return &FeedbackService{
		repo: repo,
		log:  logger,
	}
}

func (p *FeedbackService) GetFeedbackList(c *fiber.Ctx) (interface{}, error) {
	km_id, _ := utils.GenerateDecoded(c.Params("km_id"))
	return p.repo.GetFeedbackList(km_id)
}

// GetRelationList retrieves related incident from the repository.
func (s *FeedbackService) ExportFeedbackList(c *fiber.Ctx) ([][]string, error) {
	km_id, _ := utils.GenerateDecoded(c.Params("km_id"))
	allFeedback, err := s.repo.GetFeedbackList(km_id)
	if err != nil {
		return nil, err
	}

	var response [][]string
	newData := []string{"ID", "Code", "Submitter", "Rating", "Comment", "Date Submit"}
	response = append(response, newData)
	for _, param := range allFeedback {
		var id = strconv.Itoa(param.ID)
		var rating = strconv.Itoa(param.Rating) + "/5"
		var submittedDate string = ""
		// var rating = fmt.Sprintf("%.1f", float64(param.Rating)/float64(5))

		if param.DateSubmit != nil {
			submittedDate = utils.ConvertTimeToString(*param.DateSubmit, "fullname")
		}

		newParam := []string{id, param.Code, param.Submitter, rating, param.Comment, submittedDate}
		response = append(response, newParam)
	}

	return response, nil
}

func (p *FeedbackService) GetHeaderTitleExcelSvc(c *fiber.Ctx) (string, error) {
	p.log.Println("Execute function GetHeaderTitleExcelSvc")
	km_id, _ := utils.GenerateDecoded(c.Params("km_id"))
	// detailFeedback, _ := p.repo.GetDetailFeedback(c.Params("km_id"))
	detailFeedback, _ := p.repo.GetKnowledgeContent(km_id)
	now := utils.ConvertStringToTime(utils.GetTimeNow("datetime"))

	nowCustom := utils.ConvertTimeToString(now, "fullname")

	title := fmt.Sprintf("Knowledge Code : %v - Exported at : %v ", detailFeedback.KnowledgeID, nowCustom)

	return title, nil
}

// func removeHTMLTags(input string) string {
// 	// Compile the regular expression to match HTML tags
// 	re := regexp.MustCompile(`<.*?>`)
// 	// Replace all HTML tags with an empty string
// 	return re.ReplaceAllString(input, "")
// }
