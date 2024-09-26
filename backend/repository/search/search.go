package search

import (
	"fmt"
	"strconv"
	"strings"
	"sygap_new_knowledge_management/backend/entities"
	"sygap_new_knowledge_management/backend/utils"

	"github.com/gofiber/fiber/v2"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SearchRepos struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewSearchRepos(db *gorm.DB, log *logrus.Logger) *SearchRepos {
	return &SearchRepos{db, log}
}

func (r *SearchRepos) GetSearchList(c *fiber.Ctx, user_id string) ([]entities.SearchListResponse, error) {
	// advance search source, product, company
	company_id := c.Query("company_id")
	source_id := c.Query("source_id")
	service_type_id := c.Query("service_type_id")
	created_from := strings.Split(c.Query("start_date"), "T")[0]
	created_to := strings.Split(c.Query("end_date"), "T")[0]
	created_date := c.Query("created_date")
	published_date := c.Query("published_date")

	// search title, keyword
	search_keyword := c.Query("search")

	// bookmark
	bookmark := c.Query("bookmark")

	var search_list []entities.SearchListResponse

	query := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT).
		Select(`
		kc.*, knowledge_content_list.content_name as type
	`).
		Joins(`
		LEFT JOIN knowledge_content_list ON knowledge_content_list.id = kc.knowledge_content_list_id
	`)

	if search_keyword != "" {
		query.Where(
			`kc.title LIKE ? OR kc.keyword LIKE ?
		`, "%"+search_keyword+"%", "%"+search_keyword+"%")
	}

	if company_id != "" {
		query.Where(
			`kc.company_id = ?
		`, company_id)
	}

	if source_id != "" {
		query.Where(
			`kc.knowledge_content_list_id = ?
		`, source_id)
	}

	if service_type_id != "" {
		query.Where(
			`kc.service_type_id = ?
		`, service_type_id)
	}

	if bookmark == "true" {
		query.Joins(`
			LEFT JOIN knowledge_content_bookmark ON knowledge_content_bookmark.knowledge_content_id = kc.id
		`).Where("knowledge_content_bookmark.employee_id = ?", user_id)
	}

	if created_from != "" && created_to != "" {
		query.Where("DATE(kc.created_at) BETWEEN ? AND ?", created_from, created_to)
	} else if created_from != "" {
		query.Where("DATE(kc.created_at) >= ?", created_from)
	} else if created_to != "" {
		query.Where("DATE(kc.created_at) <= ?", created_to)
	}

	if created_date != "" {
		query.Order(fmt.Sprintf("kc.created_at %v", created_date))
	} else if published_date != "" {
		query.Order(fmt.Sprintf("kc.published_date %v", published_date))
	}

	if err := query.Where("kc.status = 'PUBLISHED'").Order("kc.published_date desc").Find(&search_list).Error; err != nil {
		r.log.Errorln("Error retrieving list keyword: ", err)
		return nil, err
	}

	for i, item := range search_list {
		search_list[i].Keywords = map[bool][]string{true: strings.Split(item.Keyword, ";"), false: {}}[item.Keyword != ""]
		switch item.Type { // workaround
		case "How To", "Problem Solution":
			search_list[i].Description = r.GetDetailFieldContent(item.ID, "workaround")
		case "Reference": //reference
			search_list[i].Description = r.GetDetailFieldContent(item.ID, "reference")
		case "Known Error": //fix_solution
			search_list[i].Description = r.GetDetailFieldContent(item.ID, "fix_solution")
		case "Decision Tree": //question & option
			search_list[i].Description = r.GetDetailFieldQuestion(item.ID)
		default:
			search_list[i].Description = ""
		}
	}

	return search_list, nil
}

func (r *SearchRepos) GetDetailFieldContent(id int, column string) string {
	var descriptions entities.SearchListDescriptionResponse
	if err := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_DETAIL).Where(`
		kcd.knowledge_content_id = ?
	`, id).First(&descriptions).Error; err != nil {
		r.log.Errorln("Error retrieving list companies: ", err)
		return ""
	}
	switch column {
	case "workaround":
		return descriptions.Workaround
	case "reference":
		return descriptions.Reference
	case "fix_solution":
		return descriptions.FixSolution
	default:
		return ""
	}
}

func (r *SearchRepos) GetDetailFieldQuestion(id int) string {
	var search_list []entities.SearchListQuestionResponse
	query := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_QUESTION).Select(`
		kcq.question, knowledge_content_option.option
	`).Joins(`
		LEFT JOIN knowledge_content_option ON knowledge_content_option.knowledge_content_question_id = kcq.id
	`).Where("kcq.knowledge_content_id = ?", id)
	if err := query.Limit(2).Find(&search_list).Error; err != nil {
		return ""
	}
	var description string = ""
	for i, item := range search_list {
		if i == 0 {
			description += item.Question
			description += "<br/> - " + item.Option
		} else {
			description += "<br/> - " + item.Option
		}
	}

	return description
}

func (r *SearchRepos) GetSearchDetail(id string, user_id string) (interface{}, error) {
	var search entities.KnowledgeContentSearch
	var search_detail_howto entities.SearchDetailResponse

	query := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT).Select(`
		kc.*, knowledge_content_list.content_name as type, employee.employee_name as author, kc.last_visitor
	`)

	query.Joins(`
		LEFT JOIN knowledge_content_list ON knowledge_content_list.id = kc.knowledge_content_list_id
	`).Joins(`
		LEFT JOIN employee ON employee.id = kc.created_by
	`)

	query.Where("kc.status = 'PUBLISHED' AND kc.id = ?", id)
	if err := query.First(&search).Error; err != nil {
		r.log.Errorln("Error retrieving list keyword: ", err)
		return entities.KnowledgeContentSearch{}, err
	}

	switch search.Type {
	case "How To", "Reference", "Problem Solution", "Known Error":
		detail := r.GetSearchChildDetail(id)
		search_detail_howto = entities.SearchDetailResponse{
			ID:          search.ID,
			Type:        search.Type,
			Title:       search.Title,
			LastVisitor: search.LastVisitor,
			Keywords:    map[bool][]string{true: strings.Split(search.Keyword, ";"), false: {}}[search.Keyword != ""],
			Content:     detail,
			Sidebar: entities.SearchDetailSidebarResponse{
				KnowledgeId:   search.KnowledgeId,
				Author:        search.Author,
				PublishedDay:  utils.ConvertTimeToString(search.PublishedDate, "fullname"),
				Version:       strconv.Itoa(search.Version) + ".0",
				ReportArticle: r.GetMyReport(user_id, search.ID),
				Bookmark:      r.GetMyBookmark(user_id, search.ID),
			},
			Attachment: r.GetContentAttachment(search.ID),
		}
		return search_detail_howto, nil
	case "Decision Tree":
		question := r.GetSearchQuestionDetail(id)
		search_detail_howto = entities.SearchDetailResponse{
			ID:          search.ID,
			Type:        search.Type,
			Title:       search.Title,
			LastVisitor: search.LastVisitor,
			Keywords:    map[bool][]string{true: strings.Split(search.Keyword, ";"), false: {}}[search.Keyword != ""],
			Content:     question,
			Sidebar: entities.SearchDetailSidebarResponse{
				KnowledgeId:   search.KnowledgeId,
				Author:        search.Author,
				PublishedDay:  utils.ConvertTimeToString(search.CreatedAt, "fullname"),
				Version:       strconv.Itoa(search.Version) + ".0",
				ReportArticle: r.GetMyReport(user_id, search.ID),
				Bookmark:      r.GetMyBookmark(user_id, search.ID),
			},
			Attachment: r.GetContentAttachment(search.ID),
		}
		return search_detail_howto, nil
	default:
		return nil, nil
	}
}

func (r *SearchRepos) GetSearchChildDetail(id string) interface{} {
	var detail entities.SearchDetailChildResponse
	if err := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_DETAIL).Select(`
		kcd.question, kcd.workaround, kcd.fix_solution, kcd.technical_note, kcd.reference, kcd.error, kcd.root_cause
	`).Where("kcd.knowledge_content_id = ?", id).First(&detail).Error; err != nil {
		r.log.Errorln("Error retrieving detail: ", err)
		return entities.SearchDetailChildResponse{}
	}

	return detail
}

func (r *SearchRepos) GetSearchQuestionDetail(id string) interface{} {
	var question []entities.SearchDetailQuestionResponse
	if err := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_QUESTION).Preload("Options").Where("kcq.knowledge_content_id = ?", id).Find(&question).Error; err != nil {
		r.log.Errorln("Error retrieving detail: ", err)
		return []string{}
	}

	return question
}

func (r *SearchRepos) GetMyBookmark(user_id string, content_id int) bool {
	var detail entities.KnowledgeContentBookmark
	err := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_BOOKMARK).Where("kcb.knowledge_content_id = ? AND kcb.employee_id = ?", content_id, user_id).First(&detail).Error
	if err != nil {
		return false
	} else {
		return true
	}
}

func (r *SearchRepos) GetMyReport(user_id string, content_id int) bool {
	var detail entities.KnowledgeContentReport
	err := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_REPORT).Where("kcr.knowledge_content_id = ? AND kcr.employee_id = ?", content_id, user_id).First(&detail).Error
	if err != nil {
		return false
	} else {
		return true
	}
}

func (r *SearchRepos) GetContentFeedback(id string, user_id string) (interface{}, error) {
	var search_detail_feedback []entities.SearchDetailFeedbakcResponse
	query := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_FEEDBACK).Select("kcf.*,employee.employee_name as submitter").Joins(`
		LEFT JOIN employee ON employee.id = kcf.submitter_id
	`).Where("kcf.knowledge_id = ?", id)
	err := query.Order("kcf.date_submit DESC ").Find(&search_detail_feedback).Error
	if err != nil {
		return []entities.SearchDetailFeedbakcResponse{}, err
	}
	var ratingTotal int = 0
	for _, item := range search_detail_feedback {
		ratingTotal += item.Rating
	}
	total := float64(ratingTotal) / float64(len(search_detail_feedback))
	return entities.SearchDetailFeedbakcParentResponse{
		Comments:    search_detail_feedback,
		RatingTotal: fmt.Sprintf("%.1f", total),
	}, nil
}

func (r *SearchRepos) GetContentAttachment(content_id int) interface{} {
	var detail []entities.KnowledgeContentAttachment
	err := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_ATTACHMENT).
		Where("kca.knowledge_content_id = ?", content_id).
		Where("kca.deleted_at IS NULL").
		Find(&detail).Error
	if err != nil {
		return nil
	}
	return detail
}

func (r *SearchRepos) CreateReportContent(report entities.SearchContentReport) (int, error) {
	create := r.db.Create(&report)
	if err := create.Error; err != nil {
		return 0, err
	}
	// input on update request
	payload := entities.SearchContentUpdateRequest{
		KnowledgeId:       report.KnowledgeContentId,
		UpdateRequestType: "Technical Update",
		ArticleVersion:    r.GetContentVersion(report.KnowledgeContentId),
		SubmitterId:       report.EmployeeId,
		RequestSummary:    report.Notes,
	}
	create_update := r.db.Table("knowledge_content_update_request").Create(&payload)
	if errs := create_update.Error; errs != nil {
		return 0, errs
	}

	return payload.ID, nil
}

func (r *SearchRepos) CreateReportAttachmentContent(attachment []entities.SearchContentAttachemntReport) (interface{}, error) {
	create := r.db.Create(&attachment)
	if err := create.Error; err != nil {
		return nil, err
	}
	return attachment, nil
}

func (r *SearchRepos) BookmarkContentDetail(bookmark entities.KnowledgeContentBookmark) (interface{}, error) {
	var exist entities.KnowledgeContentBookmark
	err := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_BOOKMARK).Where("kcb.knowledge_content_id = ? AND kcb.employee_id = ?", bookmark.KnowledgeContentID, bookmark.EmployeeId).First(&exist).Error
	if err != nil {
		create := r.db.Create(&bookmark)
		if err := create.Error; err != nil {
			return false, err
		}
	} else {
		r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT_BOOKMARK).Where("kcb.knowledge_content_id = ? AND kcb.employee_id = ?", bookmark.KnowledgeContentID, bookmark.EmployeeId).Delete(&exist)
	}
	if bookmark.ID == 0 {
		return false, nil
	}
	return true, nil
}

func (r *SearchRepos) FeedbackContentDetail(attachment entities.KnowledgeContentFeedback) (interface{}, error) {
	create := r.db.Create(&attachment)
	if err := create.Error; err != nil {
		return nil, err
	}
	return attachment, nil
}

func (r *SearchRepos) GetContentVersion(id int) string {
	var content_km entities.KnowledgeContentSearch
	if err := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT).Where("kc.id = ?", id).First(&content_km).Error; err != nil {
		r.log.Errorln("Error retrieving detail: ", err)
		return ""
	}

	return strconv.Itoa(content_km.Version)
}

func (r *SearchRepos) GetMinimumDetail(id int) (entities.KnowledgeContentSearch, error) {
	var search entities.KnowledgeContentSearch
	query := r.db.Table(utils.TABLE_KNOWLEDGE_CONTENT).Select(`
		kc.*, knowledge_content_list.content_name as type, employee.employee_name as author
	`)

	query.Joins(`
		LEFT JOIN knowledge_content_list ON knowledge_content_list.id = kc.knowledge_content_list_id
	`).Joins(`
		LEFT JOIN employee ON employee.id = kc.created_by
	`)

	query.Where("kc.id = ?", id)
	if err := query.First(&search).Error; err != nil {
		r.log.Errorln("Error retrieving list keyword: ", err)
		return entities.KnowledgeContentSearch{}, err
	}

	return search, nil
}

func (r *SearchRepos) GetUpsertVisitor(content_id int, user_id int, code string) bool {
	var date = utils.GetTimeNow("date")
	payload := entities.KmVisitors{
		KnowledgeId: content_id,
		UserId:      user_id,
		VisitDate:   date,
		Code:        code + "-" + date,
	}

	result := r.db.Table("km_visitors").Clauses(clause.OnConflict{
		UpdateAll: false,
	}).Create(&payload)
	if result.Error != nil {
		return false
	}

	if result.RowsAffected > 0 {
		if err := r.db.Table("knowledge_content").
			Where("id = ?", content_id).
			Update("last_visitor", gorm.Expr("last_visitor + ?", 1)).Error; err != nil {
			return false
		}
	}

	return true
}
