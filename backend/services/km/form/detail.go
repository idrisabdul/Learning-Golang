package form

import (
	"strconv"
	"strings"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/repository/km/form"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
)

type DetailService struct {
	repo *form.DetailRepos
	log  *logrus.Logger
}

func NewDetailService(repo *form.DetailRepos, log *logrus.Logger) *DetailService {
	return &DetailService{repo, log}
}

func (s *DetailService) DetailKM(idKM string, permission map[string]interface{}) (entities.DetailKM, error) {

	detailKM, errDetailKM := s.repo.DetailKM(idKM)
	if errDetailKM != nil {
		return entities.DetailKM{}, errDetailKM
	}

	keyword := strings.Split(detailKM.Keyword, ";")
	if keyword[0] == "" {
		detailKM.Keywords = []string{}
	} else {
		detailKM.Keywords = keyword
	}

	getPermission, _ := s.GetPermissionKm(idKM, permission)
	detailKM.Permission = getPermission

	return detailKM, nil
}

func (s *DetailService) DetailKMDecisionTree(idKM string, permission map[string]interface{}) (entities.DetailKMDecisionTree, error) {
	detailKMDecisionTree, errDetailKMDecisionTree := s.repo.DetailKMDecisionTree(idKM)
	if errDetailKMDecisionTree != nil {
		return entities.DetailKMDecisionTree{}, errDetailKMDecisionTree
	}

	id, _ := strconv.Atoi(idKM)
	option, errGetListOption := s.repo.GetListOptionDecisionTree(id)
	if errGetListOption != nil {
		return entities.DetailKMDecisionTree{}, errGetListOption
	}

	keywords := strings.Split(detailKMDecisionTree.Keyword, ";")
	if keywords[0] == "" {
		detailKMDecisionTree.Keywords = []string{}
	} else {
		detailKMDecisionTree.Keywords = keywords
	}

	getPermission, _ := s.GetPermissionKm(idKM, permission)
	return entities.DetailKMDecisionTree{
		ID:                     detailKMDecisionTree.ID,
		KnowledgeID:            detailKMDecisionTree.KnowledgeID,
		KnowledgeContentListID: detailKMDecisionTree.KnowledgeContentListID,
		CompanyID:              detailKMDecisionTree.CompanyID,
		Version:                detailKMDecisionTree.Version,
		Title:                  detailKMDecisionTree.Title,
		Question:               detailKMDecisionTree.Question,
		Content:                buildOptionTree(option),
		Keywords:               detailKMDecisionTree.Keywords,
		OperationCategory1ID:   detailKMDecisionTree.OperationCategory1ID,
		OperationCategory2ID:   detailKMDecisionTree.OperationCategory2ID,
		ServiceNameID:          detailKMDecisionTree.ServiceNameID,
		ServiceCategory1ID:     detailKMDecisionTree.ServiceCategory1ID,
		ServiceCategory2ID:     detailKMDecisionTree.ServiceCategory2ID,
		ExpertGroup:            detailKMDecisionTree.ExpertGroup,
		Expertee:               detailKMDecisionTree.Expertee,
		Status:                 detailKMDecisionTree.Status,
		Author:                 detailKMDecisionTree.Author,
		RetireDate:             detailKMDecisionTree.RetireDate,
		PublishedDate:          detailKMDecisionTree.PublishedDate,
		KeyContent:             detailKMDecisionTree.KeyContent,
		Permission:             getPermission,
	}, nil
}

func buildOptionTree(options []entities.KnowledgeContentOption) []*entities.KnowledgeContentOption {
	node := make(map[int]*entities.KnowledgeContentOption)
	result := make([]*entities.KnowledgeContentOption, 0)

	parentIds := make([]int, 0)

	for _, i := range options {
		nodeTemp := &i
		node[i.ID] = nodeTemp
	}

	for _, option := range options {
		if option.OptionParentId != nil {
			node[*option.OptionParentId].Options = append(node[*option.OptionParentId].Options, node[option.ID])
		} else {
			parentIds = append(parentIds, option.ID)
		}
	}

	for _, id := range parentIds {
		result = append(result, node[id])
	}

	return result
}

func (s *DetailService) GetPermissionKm(idKm string, permission map[string]interface{}) (entities.ButtonPermission, error) {
	s.log.Println("Execute function GetPermissionKm")

	var permissionResult entities.ButtonPermission

	detailKC, _ := s.repo.DetailKM(idKm)

	userID := permission["user"].(map[string]interface{})["id"].(string)
	UserIDToInt, _ := strconv.Atoi(userID)

	// get role
	getRoleManager, _ := s.repo.IsknowledgeManager(UserIDToInt, 26) // 26 is id knowledge manager in table role

	isAdmin := userID == "1"
	IsknowledgeManager := getRoleManager // knowledge user as knowledge manager
	isSme := UserIDToInt == *detailKC.Expertee
	isRequestor := userID == detailKC.CreatedBy

	currentStatus := utils.GetCurrentStatusNumber(detailKC.Status)

	// save/update and submit
	if currentStatus == 1 && (isAdmin || IsknowledgeManager || isRequestor) {
		permissionResult.Save = true
		permissionResult.Submit = true
	}

	// next step
	if currentStatus == 2 && (isAdmin || IsknowledgeManager || isRequestor) {
		permissionResult.NextStep = true
		permissionResult.Cancel = true
	}

	// approve, cancel & reject for sme
	if (isSme || isAdmin || IsknowledgeManager) && currentStatus == 3 {
		permissionResult.Approve = true
		permissionResult.Reject = true
		permissionResult.Cancel = true
	}

	// approve & reject
	if (isAdmin || IsknowledgeManager) && currentStatus == 4 {
		permissionResult.Approve = true
		permissionResult.Reject = true
	}

	// return article from publish approval for sme (exclusive for sme reviewer)
	if isSme && currentStatus == 4 {
		permissionResult.Return = true
	}

	// retire article
	if (isAdmin || IsknowledgeManager) && (currentStatus == 5) {
		permissionResult.Retire = true
	}

	// approve & reject
	if (isAdmin || IsknowledgeManager || isSme) && currentStatus == 6 {
		permissionResult.Approve = true
		permissionResult.Reject = true
	}

	// new version
	if (isAdmin || IsknowledgeManager) && (currentStatus == 7) {
		permissionResult.NewVersion = true
	}

	return permissionResult, nil
}
