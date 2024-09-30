package form

import (
	"strings"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/repository/km/form"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
)

type UpdateService struct {
	repo *form.UpdateRepos
	log  *logrus.Logger
}

func NewUpdateService(repo *form.UpdateRepos, log *logrus.Logger) *UpdateService {
	return &UpdateService{repo, log}
}

func (s *UpdateService) UpdateKM(payload entities.SubmitKMNonDecisionTree, author int, step string) error {

	// get current status
	currentStatus := utils.GetCurrentStatusNumber(payload.Status)

	// Populate Data to knowledge content
	knowledge_content := entities.KnowledgeContent{
		ID:        payload.ID,
		Status:    KMStatus(payload.Status, step),
		Keyword:   strings.Join(payload.Keyword, ";"),
		UpdatedAt: utils.GetTimeNow("normal"),
		UpdatedBy: author,
	}

	// append to knowledge_content if the status < Publish Approval
	if currentStatus < 4 {
		knowledge_content.Title = payload.Title
		knowledge_content.CompanyID = payload.Company
		knowledge_content.OperationCategory1ID = payload.OperationCategory1ID
		knowledge_content.OperationCategory2ID = payload.OperationCategory2ID
		knowledge_content.ServiceNameID = payload.ServiceType
		knowledge_content.ServiceCategory1ID = payload.ServiceCategory1ID
		knowledge_content.ServiceCategory2ID = payload.ServiceCategory2ID
		knowledge_content.ExpertGroup = payload.ExpertGroup
		knowledge_content.Expertee = payload.Expertee
	}

	// append to knowledge_content based on next status
	switch KMStatus(payload.Status, step) {
	case "PUBLISHED":
		knowledge_content.PublishedDate = utils.ConvertStringToTime(utils.GetTimeNow("normal"))
		knowledge_content.RetireDate = knowledge_content.PublishedDate.AddDate(1, 0, 0)
	case "RETIRED":
		knowledge_content.IsRetired = 2
	}

	// then assign it to table
	if errUpdateToKnowledgeContent := s.repo.UpdateToKnowledgeContent(knowledge_content); errUpdateToKnowledgeContent != nil {
		return errUpdateToKnowledgeContent
	}

	// do content update if the status < Publish Approval
	if currentStatus < 4 {
		// populate data to knowledge content detail
		knowledge_content_detail := entities.KnowledgeContentDetail{
			KnowledgeContentID: payload.ID,
			Question:           payload.Question,
			Error:              payload.Error,
			RootCause:          payload.RootCause,
			Workaround:         payload.Workaround,
			FixSolution:        payload.FixSolution,
			TechnicalNote:      payload.TechnicalNote,
			Reference:          payload.Reference,
		}

		// then assign it to table
		if errUpdateToKnowledgeContentDetail := s.repo.UpdateToKnowledgeContentDetail(knowledge_content_detail); errUpdateToKnowledgeContentDetail != nil {
			return errUpdateToKnowledgeContentDetail
		}
	}

	//to knowledge content log
	knowledge_content_log := entities.KnowledgeContentLog{
		KnowledgeContentID: payload.ID,
		Action:             LogAction(step),
		Note:               payload.Note,
		Status:             KMStatus(payload.Status, step),
		CreatedBy:          author,
	}

	// then assign to table
	if errSubmitToKnowledgeContentLog := s.repo.InstanceSubmitToKnowledgeContentLog(knowledge_content_log); errSubmitToKnowledgeContentLog != nil {
		return errSubmitToKnowledgeContentLog
	}

	if payload.Status != KMStatus(payload.Status, step) {
		workdetail_knowledge_management := entities.WorkdetailKnowledgeManagement{
			IDParent:  payload.ID,
			Type:      "Status Updated",
			Note:      "Status Updated To : " + KMStatus(payload.Status, step) + " <br>Note : " + payload.Note,
			CreatedBy: author,
		}

		_, errSubmitToWorkdetailKnowledgeManagement := s.repo.InstanceSubmitToWorkdetailKnowledgeManagement(workdetail_knowledge_management)
		if errSubmitToWorkdetailKnowledgeManagement != nil {
			return errSubmitToWorkdetailKnowledgeManagement
		}

	}

	return nil
}

func (s *UpdateService) UpdateKMDecisionTree(payload entities.SubmitKMDecisionTree, author int, step string) error {

	// get current status
	currentStatus := utils.GetCurrentStatusNumber(payload.Status)

	// populate data to knowledge content
	knowledge_content := entities.KnowledgeContent{
		ID:        payload.ID,
		Status:    KMStatus(payload.Status, step),
		Keyword:   strings.Join(payload.Keyword, ";"),
		UpdatedAt: utils.GetTimeNow("normal"),
		UpdatedBy: author,
	}

	// append to knowledge_content if the status < Publish Approval
	if currentStatus < 4 {
		knowledge_content.Title = payload.Title
		knowledge_content.CompanyID = payload.Company
		knowledge_content.OperationCategory1ID = payload.OperationCategory1ID
		knowledge_content.OperationCategory2ID = payload.OperationCategory2ID
		knowledge_content.ServiceNameID = payload.ServiceType
		knowledge_content.ServiceCategory1ID = payload.ServiceCategory1ID
		knowledge_content.ServiceCategory2ID = payload.ServiceCategory2ID
		knowledge_content.ExpertGroup = payload.ExpertGroup
		knowledge_content.Expertee = payload.Expertee
	}

	// append to knowledge_content based on next status
	switch KMStatus(payload.Status, step) {
	case "PUBLISHED":
		knowledge_content.PublishedDate = utils.ConvertStringToTime(utils.GetTimeNow("normal"))
		knowledge_content.RetireDate = knowledge_content.PublishedDate.AddDate(1, 0, 0)
	case "RETIRED":
		knowledge_content.IsRetired = 2
	}

	// then assign it to table
	if errUpdateToKnowledgeContent := s.repo.UpdateToKnowledgeContent(knowledge_content); errUpdateToKnowledgeContent != nil {
		return errUpdateToKnowledgeContent
	}

	// do content update if the status < Publish Approval
	if currentStatus < 4 {

		// to knowledge_content_detail
		knowledge_content_detail := entities.KnowledgeContentDetail{
			KnowledgeContentID: payload.ID,
			Question:           payload.Question,
		}

		if errSubmitToKnowledgeContentDetail := s.repo.UpdateToKnowledgeContentDetail(knowledge_content_detail); errSubmitToKnowledgeContentDetail != nil {
			return errSubmitToKnowledgeContentDetail
		}

		// Populate questions and answers
		var questions []entities.KnowledgeContentQuestion
		var options []entities.KnowledgeContentOption

		var newQuestions []entities.KnowledgeContentQuestion
		var newOptions []entities.KnowledgeContentOption
		for _, q := range payload.Content {
			switch q.Action {
			case "delete":
				questions = append(questions, entities.KnowledgeContentQuestion{
					ID:        q.ID,
					DeletedAt: utils.ConvertStringToTime(utils.GetTimeNow("normal")),
					DeletedBy: author,
				})

				options = append(options,
					entities.KnowledgeContentOption{
						ID:        q.Options[0].ID,
						DeletedAt: utils.ConvertStringToTime(utils.GetTimeNow("normal")),
						DeletedBy: author,
					},
					entities.KnowledgeContentOption{
						ID:        q.Options[1].ID,
						DeletedAt: utils.ConvertStringToTime(utils.GetTimeNow("normal")),
						DeletedBy: author,
					},
				)
			case "update", "none":
				questions = append(questions, entities.KnowledgeContentQuestion{
					ID:       q.ID,
					Question: q.Question,
				})

				options = append(options,
					entities.KnowledgeContentOption{
						ID:       q.Options[0].ID,
						Label:    q.Options[0].Option,
						Solution: q.Options[0].Answer,
					},
					entities.KnowledgeContentOption{
						ID:       q.Options[1].ID,
						Label:    q.Options[1].Option,
						Solution: q.Options[1].Answer,
					},
				)
			case "new":
				newQuestions = append(newQuestions, entities.KnowledgeContentQuestion{
					KnowledgeContentID: payload.ID,
					Question:           q.Question,
				})
				newOptions = append(newOptions,
					entities.KnowledgeContentOption{
						Label:    q.Options[0].Option,
						Solution: q.Options[0].Answer,
					},
					entities.KnowledgeContentOption{
						Label:    q.Options[1].Option,
						Solution: q.Options[1].Answer,
					})
			}

		}

		// then assign questions to table
		for _, v := range questions {
			if errUpdateQuestions := s.repo.UpdateToKnowledgeContentQuestion(v); errUpdateQuestions != nil {
				return errUpdateQuestions
			}
		}

		// same goes to options
		for _, v := range options {
			if errUpdateOptions := s.repo.UpdateToKnowledgeContentOption(v); errUpdateOptions != nil {
				return errUpdateOptions
			}
		}

		// if there are any new questions and options
		if len(newQuestions) != 0 {

			// assign new questions to table first, then resolve the ids
			IDQuestions, errSubmitToKnowledgeContentQuestion := s.repo.InstanceSubmitToKnowledgeContentQuestion(newQuestions)
			if errSubmitToKnowledgeContentQuestion != nil {
				return errSubmitToKnowledgeContentQuestion
			}

			// then populate the ids to each option
			var i float64 = 2 * float64(len(IDQuestions))
			var x float64 = 0

			for x < i {
				x += 1
			}

			// finally assign the options into table
			if errSubmitToKnowledgeContentOption := s.repo.InstanceSubmitToKnowledgeContentOption(newOptions); errSubmitToKnowledgeContentOption != nil {
				return errSubmitToKnowledgeContentOption
			}
		}

	}
	//to knowledge content log
	knowledge_content_log := entities.KnowledgeContentLog{
		KnowledgeContentID: payload.ID,
		Action:             LogAction(step),
		Note:               payload.Note,
		Status:             KMStatus(payload.Status, step),
		CreatedBy:          author,
	}

	// then assign to table
	if errSubmitToKnowledgeContentLog := s.repo.InstanceSubmitToKnowledgeContentLog(knowledge_content_log); errSubmitToKnowledgeContentLog != nil {
		return errSubmitToKnowledgeContentLog
	}

	if payload.Status != KMStatus(payload.Status, step) {
		workdetail_knowledge_management := entities.WorkdetailKnowledgeManagement{
			IDParent:  payload.ID,
			Type:      "Status Updated",
			Note:      "Status Updated To : " + KMStatus(payload.Status, step) + " <br>Note : " + payload.Note,
			CreatedBy: author,
		}

		_, errSubmitToWorkdetailKnowledgeManagement := s.repo.InstanceSubmitToWorkdetailKnowledgeManagement(workdetail_knowledge_management)
		if errSubmitToWorkdetailKnowledgeManagement != nil {
			return errSubmitToWorkdetailKnowledgeManagement
		}

	}
	return nil
}

func (s *UpdateService) SetClosedVersion(idKM, author int, note string) error {

	// set the date first
	now := utils.GetTimeNow("normal")

	// populate data to knowledge_content
	knowledge_content := entities.KnowledgeContent{
		ID:        idKM,
		Status:    "CLOSED VERSION",
		UpdatedAt: now,
		UpdatedBy: author,
		DeletedAt: now,
		DeletedBy: author,
		IsRetired: 1,
	}

	// then assign to table
	if errUpdateToKnowledgeContent := s.repo.UpdateToKnowledgeContent(knowledge_content); errUpdateToKnowledgeContent != nil {
		return errUpdateToKnowledgeContent
	}

	// populate data to knowledge_content_log
	knowledge_content_log := entities.KnowledgeContentLog{
		KnowledgeContentID: idKM,
		Action:             "UPDATE",
		Status:             "CLOSED VERSION",
		Note:               note,
		CreatedBy:          author,
	}

	// then assign it to table
	if errSubmitToKnowledgeContentLog := s.repo.InstanceSubmitToKnowledgeContentLog(knowledge_content_log); errSubmitToKnowledgeContentLog != nil {
		return errSubmitToKnowledgeContentLog
	}

	workdetail_knowledge_management := entities.WorkdetailKnowledgeManagement{
		IDParent:  idKM,
		Type:      "Status Updated",
		Note:      "Status Updated To : CLOSED VERSION <br>Note : " + note,
		CreatedBy: author,
	}

	_, errSubmitToWorkdetailKnowledgeManagement := s.repo.InstanceSubmitToWorkdetailKnowledgeManagement(workdetail_knowledge_management)
	if errSubmitToWorkdetailKnowledgeManagement != nil {
		return errSubmitToWorkdetailKnowledgeManagement

	}

	return nil
}

/*
This is local function, make sure only used in this file. if you want to use it globally, move it to utils
*/

// This function is used for decide which action will inserted to knowledge_content_log
func LogAction(step string) string {
	if step != "none" && step != "update" {
		return strings.ToUpper(step)
	}
	return "UPDATE"
}

// This function is used for decide the next status of the KM
func KMStatus(status, step string) string {
	currentStatus := utils.GetCurrentStatusNumber(status)
	var nextStatus string

	if step == "next" {
		nextStatus = utils.GetNextStatus(currentStatus + 1)
	} else if step == "reject" {
		if currentStatus == 4 {
			nextStatus = "DRAFT"
		} else {
			nextStatus = utils.GetNextStatus(currentStatus - 1)
		}
	} else if step == "cancel" {
		nextStatus = "CANCELLED"
	} else if step == "return" {
		nextStatus = "DRAFT"
	} else {
		nextStatus = status
	}

	return nextStatus
}
