package km

import (
	"fmt"
	"strings"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type KMListRepos struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewKMListRepos(db *gorm.DB, log *logrus.Logger) *KMListRepos {
	return &KMListRepos{db, log}
}

func (r *KMListRepos) GetListKM(payload map[string]any) ([]entities.ListKM, error) {
	var KMList []entities.ListKM

	query := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT).
		Select(`kc.id, 
		kc.knowledge_id, 
		kc.title, 
		kc.status, 
		kcl.content_name AS source_name, 
		kc.status, 
		e1.employee_name AS author,
		e2.employee_name AS assignee,
		CASE 
		WHEN kc.updated_at IS NOT NULL THEN kc.updated_at 
		ELSE kc.created_at END AS modified_date,
		CASE
		WHEN kcur.knowledge_id IS NULL THEN "No"
		ELSE "Yes" END AS request_update
		`).
		Joins("LEFT JOIN " + utils.TABLE_EMPLOYEE + "1 ON kc.created_by = e1.id").
		Joins("LEFT JOIN " + utils.TABLE_EMPLOYEE + "2 ON kc.expertee = e2.id").
		Joins("LEFT JOIN " + utils.TABLE_KNOWLEDGE_CONTENT_LIST + " ON kc.knowledge_content_list_id = kcl.id").
		Joins("LEFT JOIN " + utils.TABLE_KNOWLEDGE_CONTENT_UPDATE_REQUEST + " ON kc.id = kcur.knowledge_id")

	for k, v := range payload {
		switch k {
		case "search":
			if v != "" {
				query = query.Where(r.db.Where("kc.title LIKE ?", "%"+v.(string)+"%").Or("kc.knowledge_id LIKE ?", "%"+v.(string)+"%"))
			}
		case "status":
			if v != "" {
				query = query.Where("kc.status = ?", v)
			}
		case "assign_to":
			if v != "" {
				assignTo := strings.Split(v.(string), " ")
				if assignTo[0] == "ME" {
					query = query.Where("kc.expertee = ? ", assignTo[1])
				} else {
					query = query.Where("kc.expert_group = ?", assignTo[1])
				}
			}
		case "company":
			if v != "" {
				query = query.Where("kc.company_id = ?", v)
			}
		case "content_type":
			if v != "" {
				query = query.Where("kc.knowledge_content_list_id = ?", v)
			}
		case "created_from", "created_to":
			if v != "" {
				payload["created_from"] = strings.Split(payload["created_from"].(string), "T")[0]
				payload["created_to"] = strings.Split(payload["created_to"].(string), "T")[0]
				if payload["created_from"] != "" && payload["created_to"] != "" {
					query = query.Where("DATE(kc.created_at) BETWEEN ? AND ?", payload["created_from"], payload["created_to"])
				} else if payload["created_from"] != "" {
					query = query.Where("DATE(kc.created_at) >= ?", payload["created_from"])
				} else if payload["created_to"] != "" {
					query = query.Where("DATE(kc.created_at) <= ?", payload["created_to"])
				}
			}
		case "expertee_group":
			if v != "" {
				query = query.Where("kc.expert_group = ?", v)
			}
		case "expertee":
			if v != "" {
				query = query.Where("kc.expertee = ?", v)
			}
		case "product_name":
			if v != "" {
				query = query.Where("kc.service_name_id = ?", v)
			}
		case "submitted_by":
			if v != "" {
				query = query.Where("kc.created_by = ?", v)
			}
		}
	}
	if errGetKMList := query.Group("kc.id").
		Scopes(func(d *gorm.DB) *gorm.DB {
			if payload["published_date"] != "" {
				return d.Order(fmt.Sprintf("kc.published_date %v", payload["published_date"].(string)))
			}
			if payload["created_at"] != "" {
				return d.Order(fmt.Sprintf("kc.created_at %v", payload["created_at"].(string)))
			}
			return query.Order("kc.created_at DESC")
		}).
		Find(&KMList).Error; errGetKMList != nil {
		r.log.Errorln("Error on GetListKM in Repo: ", errGetKMList)
		return nil, errGetKMList
	}

	return KMList, nil
}
