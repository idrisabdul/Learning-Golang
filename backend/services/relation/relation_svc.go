package relation_svc

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/model"
	repository "sygap_new_knowledge_management/backend/repository/relation"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
)

type RelationSvc struct {
	repo   *repository.RelationRepo
	logger *logrus.Logger
}

func NewRelationSvc(repo *repository.RelationRepo, logger *logrus.Logger) *RelationSvc {
	return &RelationSvc{repo, logger}
}

func (s *RelationSvc) GetRelationList(params entities.SearchRelationParams) ([]model.RelationFilteredResponseModel, entities.Priority, error) {
	// Retrieves related incidents from the repository
	allRelations, err := s.repo.GetRelationList(params)
	if err != nil {
		return nil, entities.Priority{}, err
	}

	// Initialize priority counters
	priority := entities.Priority{}

	// Function to increment priority fields
	incrementPriority := func(p string) {
		switch strings.ToLower(p) {
		case "critical":
			priority.Critical++
		case "major":
			priority.Mayor++
		case "high":
			priority.High++
		case "medium":
			priority.Medium++
		case "low":
			priority.Low++
		}
	}

	// Loop through related incidents and increment priorities
	for _, emp := range allRelations {
		incrementPriority(emp.Priority)
	}

	return allRelations, priority, nil
}

func (s *RelationSvc) GetReqTypeRelationList(params entities.SearchReqTypeRelationParams, permission map[string]interface{}, codeList []string) ([]model.ReqTypeRelationResponseModel, error) {

	idUser, ok := permission["user"].(map[string]interface{})["id"].(string)
	if !ok {
		return nil, errors.New("unable to retrieve user ID")
	}

	// Retrieves related incidents from the repository
	allRelations, err := s.repo.GetReqTypeRelationList(params, idUser, codeList)
	if err != nil {
		return nil, err
	}

	return allRelations, nil
}

func (s *RelationSvc) InsertRelation(payload model.InsertRelationParams, permission map[string]interface{}) error {
	userId, ok := permission["user"].(map[string]interface{})["id"].(string)
	if !ok {
		return errors.New("user ID not found in permission")
	}

	userIdToInt, err := strconv.Atoi(userId)
	if err != nil {
		return err
	}

	return s.repo.InsertRelation(payload, userIdToInt)
}

func (s *RelationSvc) DeleteRelations(id string, permission map[string]interface{}) error {
	userId, ok := permission["user"].(map[string]interface{})["id"].(string)
	if !ok {
		return errors.New("user ID not found in permission")
	}

	userIdToInt, err := strconv.Atoi(userId)
	if err != nil {
		return err
	}

	idSlice := make([]int, 0)
	for _, v := range strings.Split(id, ",") {
		id, err := strconv.Atoi(v)
		if err == nil {
			idSlice = append(idSlice, id)
		}
	}

	return s.repo.DeleteRelations(idSlice, userIdToInt)
}

// GetRelationList retrieves related incident from the repository.
func (s *RelationSvc) ExportRelationList(params entities.SearchRelationParams) ([][]string, error) {
	allRelations, err := s.repo.GetRelationList(params)
	if err != nil {
		return nil, err
	}

	var response [][]string
	newData := []string{"ID", "Summary", "Note", "Relation Type", "Priority", "Status", "Created Date", "Resolution", "Resolution Date"}
	response = append(response, newData)
	for _, param := range allRelations {
		var created_date string = ""
		var resolution_date string = ""

		if param.CreatedDate != nil {
			created_date = utils.ConvertTimeToString(*param.CreatedDate, "fullname")
		}
		if param.ResolutionDate != nil {
			resolution_date = utils.ConvertTimeToString(*param.ResolutionDate, "fullname")
		}
		cleanSummaryHtml := removeHTMLTags(param.Summary)
		cleanNoteHtml := removeHTMLTags(param.Note)
		cleanResolutionHtml := removeHTMLTags(param.Resolution)
		newParam := []string{param.ID, cleanSummaryHtml, cleanNoteHtml, param.RelationType, param.Priority, param.Status, created_date, cleanResolutionHtml, resolution_date}
		response = append(response, newParam)
	}

	return response, nil
}

func (p *RelationSvc) GetHeaderTittleExcelSvc(ID string) (string, error) {
	p.logger.Println("Execute function GetHeaderTittleExcelSvc")
	KMID, _ := strconv.Atoi(ID)
	detailKM, _ := p.repo.GetDetailKM(KMID)
	now := utils.ConvertStringToTime(utils.GetTimeNow("datetime"))

	nowCustom := utils.ConvertTimeToString(now, "fullname")

	title := fmt.Sprintf("Ticket code : %v - Exported at : %v ", detailKM.KnowledgeID, nowCustom)

	return title, nil
}

func removeHTMLTags(input string) string {
	// Compile the regular expression to match HTML tags
	re := regexp.MustCompile(`<.*?>`)
	// Replace all HTML tags with an empty string
	return re.ReplaceAllString(input, "")
}
