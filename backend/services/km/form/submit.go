package form

import (
	"math"
	"strconv"
	"strings"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/repository/km/form"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
)

type SubmitService struct {
	repo *form.SubmitRepos
	log  *logrus.Logger
}

func NewSubmitService(repos *form.SubmitRepos, log *logrus.Logger) *SubmitService {
	return &SubmitService{repos, log}
}

func (s *SubmitService) SubmitKM(payload entities.SubmitKMNonDecisionTree, author int, toDraft bool) (string, error) {

	// to knowledge_content
	// populate data first
	knowledge_content := &entities.KnowledgeContent{
		Title:                  payload.Title,
		Version:                payload.Version,
		KnowledgeContentListID: payload.KnowledgeContentListID,
		CompanyID:              payload.Company,
		Status:                 SetStatusSubmit(toDraft),
		OperationCategory1ID:   payload.OperationCategory1ID,
		OperationCategory2ID:   payload.OperationCategory2ID,
		ServiceNameID:          payload.ServiceType,
		ServiceCategory1ID:     payload.ServiceCategory1ID,
		ServiceCategory2ID:     payload.ServiceCategory2ID,
		ExpertGroup:            payload.ExpertGroup,
		Expertee:               payload.Expertee,
		CreatedBy:              author,
		Keyword:                strings.Join(payload.Keyword, ";"),
	}

	// then assign to db, resolve ID from submit
	IDresponse, errSubmitToKnowledgeContent := s.repo.SubmitToKnowledgeContent(knowledge_content)
	if errSubmitToKnowledgeContent != nil {
		return "", errSubmitToKnowledgeContent
	}

	// to knowledge_content_detail
	knowledge_content_detail := &entities.KnowledgeContentDetail{
		KnowledgeContentID: IDresponse,
	}

	// define field based on id knowledge content list
	switch payload.KnowledgeContentListID {
	case 1, 3:
		knowledge_content_detail.Question = payload.Question
		knowledge_content_detail.Workaround = payload.Workaround
		knowledge_content_detail.FixSolution = payload.FixSolution
		knowledge_content_detail.TechnicalNote = payload.TechnicalNote
	case 2:
		knowledge_content_detail.Reference = payload.Reference
	case 4:
		knowledge_content_detail.Error = payload.Error
		knowledge_content_detail.RootCause = payload.RootCause
		knowledge_content_detail.Workaround = payload.Workaround
		knowledge_content_detail.FixSolution = payload.FixSolution
		knowledge_content_detail.TechnicalNote = payload.TechnicalNote
	}

	if errSubmitToKnowledgeContentDetail := s.repo.SubmitToKnowledgeContentDetail(knowledge_content_detail); errSubmitToKnowledgeContentDetail != nil {
		return "", errSubmitToKnowledgeContentDetail
	}

	//to knowledge content log
	knowledge_content_log := entities.KnowledgeContentLog{
		KnowledgeContentID: IDresponse,
		Action:             "SUBMIT",
		Note:               payload.Note,
		Status:             SetStatusSubmit(toDraft),
		CreatedBy:          author,
	}

	// then assign to table
	if errSubmitToKnowledgeContentLog := s.repo.SubmitToKnowledgeContentLog(knowledge_content_log); errSubmitToKnowledgeContentLog != nil {
		return "", errSubmitToKnowledgeContentLog
	}

	workdetail_knowledge_management := entities.WorkdetailKnowledgeManagement{
		IDParent:  IDresponse,
		Type:      "Status Created",
		Note:      "Status Created To : " + SetStatusSubmit(toDraft) + " <br>Note : " + payload.Note,
		CreatedBy: author,
	}

	_, errSubmitToWorkdetailKnowledgeManagement := s.repo.InstanceSubmitToWorkdetailKnowledgeManagement(workdetail_knowledge_management)
	if errSubmitToWorkdetailKnowledgeManagement != nil {
		return "", errSubmitToWorkdetailKnowledgeManagement
	}

	knowledge_content_log_version := entities.KnowledgeContentLogVersion{
		IDKnowledgeContent: IDresponse,
		CreatedBy:          strconv.Itoa(author),
	}
	if payload.KeyContent != 0 {
		knowledge_content_log_version.KeyContent = strconv.Itoa(payload.KeyContent)
	} else {
		knowledge_content_log_version.IsFirst = "true"
	}

	if errSubmitToKnowledgeContentLogVersion := s.repo.SubmitToKnowledgeContentLogVersion(knowledge_content_log_version); errSubmitToKnowledgeContentLogVersion != nil {
		return "", errSubmitToKnowledgeContentLogVersion
	}

	if payload.KeyContent != 0 {

		list, _ := s.repo.InstanceDetailLogVersion(payload.KeyContent)
		createdAt := utils.ConvertStringToTime(utils.GetTimeNow("normal"))
		currentID := payload.KeyContent

		var fromPrevious []entities.KnowledgeRelationToTicketPopup
		var toNew []entities.KnowledgeRelationToTicketPopup

		for _, v := range list {

			if len(list) > 1 {
				currentID = v.IDKnowledgeContent
			}

			if IDresponse == v.IDKnowledgeContent && len(list) > 1 {
				continue
			}

			fromPrevious = append(fromPrevious, entities.KnowledgeRelationToTicketPopup{
				EntityType:    "knowledge_management",
				IDEntityType:  currentID,
				RequestType:   "knowledge_management",
				IDRequestType: IDresponse,
				RelationType:  "Related To",
				CreatedAt:     &createdAt,
				CreatedBy:     author,
			})

			toNew = append(toNew, entities.KnowledgeRelationToTicketPopup{
				EntityType:    "knowledge_management",
				IDEntityType:  IDresponse,
				RequestType:   "knowledge_management",
				IDRequestType: currentID,
				RelationType:  "Related To",
				CreatedAt:     &createdAt,
				CreatedBy:     author,
			})

		}

		combinedData := append(fromPrevious, toNew...)

		if errSubmitToKnowledgeRelationToTicket := s.repo.InstanceSubmitToKnowledgeRelationToTicket(combinedData); errSubmitToKnowledgeRelationToTicket != nil {
			return "", errSubmitToKnowledgeRelationToTicket
		}

	}

	return utils.GenerateNumberEncode(strconv.Itoa(IDresponse)), nil
}

func (s *SubmitService) SubmitKMDecisionTree(payload entities.SubmitKMDecisionTree, author int, toDraft bool) (string, error) {

	// to knowledge content
	knowledge_content := &entities.KnowledgeContent{
		Version:                payload.Version,
		Title:                  payload.Title,
		Status:                 SetStatusSubmit(toDraft),
		CompanyID:              payload.Company,
		KnowledgeContentListID: payload.KnowledgeContentListID,
		OperationCategory1ID:   payload.OperationCategory1ID,
		OperationCategory2ID:   payload.OperationCategory2ID,
		ServiceNameID:          payload.ServiceType,
		ServiceCategory1ID:     payload.ServiceCategory1ID,
		ServiceCategory2ID:     payload.ServiceCategory2ID,
		ExpertGroup:            payload.ExpertGroup,
		Expertee:               payload.Expertee,
		CreatedBy:              author,
		Keyword:                strings.Join(payload.Keyword, ";"),
	}

	IDResponse, errSubmitToKnowledgeContent := s.repo.SubmitToKnowledgeContent(knowledge_content)
	if errSubmitToKnowledgeContent != nil {
		return "", errSubmitToKnowledgeContent
	}

	knowledge_content_detail := entities.KnowledgeContentDetail{
		KnowledgeContentID: IDResponse,
		Question:           payload.Question,
	}

	if errSubmitToKnowledgeContentDetail := s.repo.SubmitToKnowledgeContentDetail(&knowledge_content_detail); errSubmitToKnowledgeContentDetail != nil {
		return "", errSubmitToKnowledgeContentDetail
	}

	// populate data first
	var questions []entities.KnowledgeContentQuestion
	var options []entities.KnowledgeContentOption
	for _, question := range payload.Content {
		questions = append(questions, entities.KnowledgeContentQuestion{
			KnowledgeContentID: IDResponse,
			Question:           question.Question,
		})

		for _, v := range question.Options {
			options = append(options, entities.KnowledgeContentOption{
				Option:   v.Option,
				Solution: v.Answer,
			})
		}
	}

	// assign questions to knowledge content question, the resolve the id question
	IDQuestion, errSubmitToKnowledgeContentQuestion := s.repo.SubmitToKnowledgeContentQuestion(questions)
	if errSubmitToKnowledgeContent != nil {
		return "", errSubmitToKnowledgeContentQuestion
	}

	// insert the id question to options
	// credit to anton
	var i float64 = 2 * float64(len(IDQuestion))
	var x float64 = 0

	for x < i {
		var j float64 = x / 2
		k := int(math.Floor(j))
		l := IDQuestion[k]
		options[int(x)].KnowledgeContentQuestionID = l
		x += 1
	}

	// then assign options to knowledge content option
	if errSubmitToKnowledgeContentOption := s.repo.SubmitToKnowledgeContentOption(options); errSubmitToKnowledgeContentOption != nil {
		return "", errSubmitToKnowledgeContentOption
	}

	//to knowledge content log
	knowledge_content_log := entities.KnowledgeContentLog{
		KnowledgeContentID: IDResponse,
		Action:             "SUBMIT",
		Note:               payload.Note,
		Status:             SetStatusSubmit(toDraft),
		CreatedBy:          author,
	}

	// then assign to table
	if errSubmitToKnowledgeContentLog := s.repo.SubmitToKnowledgeContentLog(knowledge_content_log); errSubmitToKnowledgeContentLog != nil {
		return "", errSubmitToKnowledgeContentLog
	}

	workdetail_knowledge_management := entities.WorkdetailKnowledgeManagement{
		IDParent:  IDResponse,
		Type:      "Status Created",
		Note:      "Status Created To : " + SetStatusSubmit(toDraft) + " <br>Note : " + payload.Note,
		CreatedBy: author,
	}

	_, errSubmitToWorkdetailKnowledgeManagement := s.repo.InstanceSubmitToWorkdetailKnowledgeManagement(workdetail_knowledge_management)
	if errSubmitToWorkdetailKnowledgeManagement != nil {
		return "", errSubmitToWorkdetailKnowledgeManagement
	}

	knowledge_content_log_version := entities.KnowledgeContentLogVersion{
		IDKnowledgeContent: IDResponse,
		CreatedBy:          strconv.Itoa(author),
	}
	if payload.KeyContent != 0 {
		knowledge_content_log_version.KeyContent = strconv.Itoa(payload.KeyContent)
	} else {
		knowledge_content_log_version.IsFirst = "true"
	}

	if errSubmitToKnowledgeContentLogVersion := s.repo.SubmitToKnowledgeContentLogVersion(knowledge_content_log_version); errSubmitToKnowledgeContentLogVersion != nil {
		return "", errSubmitToKnowledgeContentLogVersion
	}

	if payload.KeyContent != 0 {

		list, _ := s.repo.InstanceDetailLogVersion(payload.KeyContent)
		createdAt := utils.ConvertStringToTime(utils.GetTimeNow("normal"))
		currentID := payload.KeyContent

		var fromPrevious []entities.KnowledgeRelationToTicketPopup
		var toNew []entities.KnowledgeRelationToTicketPopup

		for _, v := range list {

			if len(list) > 1 {
				currentID = v.IDKnowledgeContent
			}

			if IDResponse == v.IDKnowledgeContent && len(list) > 1 {
				continue
			}

			fromPrevious = append(fromPrevious, entities.KnowledgeRelationToTicketPopup{
				EntityType:    "knowledge_management",
				IDEntityType:  currentID,
				RequestType:   "knowledge_management",
				IDRequestType: IDResponse,
				RelationType:  "Related To",
				CreatedAt:     &createdAt,
				CreatedBy:     author,
			})

			toNew = append(toNew, entities.KnowledgeRelationToTicketPopup{
				EntityType:    "knowledge_management",
				IDEntityType:  IDResponse,
				RequestType:   "knowledge_management",
				IDRequestType: currentID,
				RelationType:  "Related To",
				CreatedAt:     &createdAt,
				CreatedBy:     author,
			})

		}

		combinedData := append(fromPrevious, toNew...)

		if errSubmitToKnowledgeRelationToTicket := s.repo.InstanceSubmitToKnowledgeRelationToTicket(combinedData); errSubmitToKnowledgeRelationToTicket != nil {
			return "", errSubmitToKnowledgeRelationToTicket
		}

	}

	return utils.GenerateNumberEncode(strconv.Itoa(IDResponse)), nil
}

func SetStatusSubmit(isDraft bool) string {
	if isDraft {
		return "DRAFT"
	} else {
		return "IN PROGRESS"
	}
}
