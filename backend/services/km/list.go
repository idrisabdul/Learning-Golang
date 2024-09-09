package km

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/repository/km"

	"github.com/sirupsen/logrus"
)

type KMListService struct {
	repo *km.KMListRepos
	log  *logrus.Logger
}

func NewKMListService(repo *km.KMListRepos, log *logrus.Logger) *KMListService {
	return &KMListService{repo, log}
}

func (s *KMListService) GetListKM(payload entities.SearchListKM, permission map[string]any) ([]entities.ListKM, error) {

	filters := map[string]any{
		"search":         payload.Search,
		"company":        payload.Company,
		"content_type":   payload.ContentType,
		"created_from":   payload.CreateFrom,
		"created_to":     payload.CreatedTo,
		"expertee_group": payload.ExperteeGroup,
		"expertee":       payload.Expertee,
		"published_date": payload.PublishedDate,
		"created_at":     payload.CreatedDate,
		"product_name":   payload.ProductName,
		"status":         payload.Status,
		"assign_to":      "",
		"submitted_by":   "",
	}

	if payload.AssignTo == "ME" {
		filters["assign_to"] = payload.AssignTo + " " + permission["user"].(map[string]any)["id"].(string)
	}

	if payload.AssignTo == "GROUP" {
		filters["assign_to"] = payload.AssignTo + " " + permission["user"].(map[string]any)["organization"].(string)
	}

	if payload.SubmittedByMe {
		filters["submitted_by"] = permission["user"].(map[string]any)["id"].(string)
	}

	return s.repo.GetListKM(filters)
}
