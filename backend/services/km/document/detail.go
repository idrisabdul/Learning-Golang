package document

import (
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/repository/km/document"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type DetailService struct {
	repo *document.DetailRepos
	log  *logrus.Logger
}

func NewDetailService(repo *document.DetailRepos, log *logrus.Logger) *DetailService {
	return &DetailService{repo, log}
}

func (s *DetailService) DetailDocumentKM(idKM string) ([]entities.DetailDocument, error) {

	var response []entities.DetailDocument

	detailDocumentKM, errDetailDocumentKM := s.repo.DetailDocumentKM(idKM)
	if errDetailDocumentKM != nil {
		return nil, errDetailDocumentKM
	}

	for _, v := range detailDocumentKM {
		response = append(response, entities.DetailDocument{
			ID:       v.ID,
			Filename: v.Attachment,
			FileHash: v.Filename,
			Size:     int(v.Size),
		})
	}

	return response, nil
}

func (s *DetailService) GetFileLink(name string) (map[string]any, error) {
	fileURL, errFileURL := utils.GetFileURL(&fiber.Ctx{}, name)

	if errFileURL != nil {
		s.log.Errorln("Error while trying to get file url: ", errFileURL.Error())
		return nil, errFileURL
	}

	return map[string]any{
		"file_name": name,
		"file_url":  fileURL,
	}, nil
}
