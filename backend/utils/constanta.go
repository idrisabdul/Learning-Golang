package utils

import "strings"

// Table
const (

	//Options
	TABLE_OPERATIONAL_CATEGORIES         = "operational_categories oc"
	TABLE_PRODUCT_CATEGORIES             = "product_categories pc"
	TABLE_COMPANY                        = "company comp"
	TABLE_ORGANIZATION                   = "organization o"
	TABLE_EMPLOYEE                       = "employee e"
	TABLE_ORGANIZATION_EMPLOYEE_HAS_ROLE = "organization_employee_has_role oehr"
	TABLE_KNOWLEDGE_CONTENT_LIST         = "knowledge_content_list kcl"
	TABLE_STATUS                         = "status s"
	TABLE_WORK_DETAIL_TYPE               = "work_detail_type wdt"

	//Knowledge Content
	TABLE_KNOWLEDGE_CONTENT             = "knowledge_content kc"
	TABLE_KNOWLEDGE_CONTENT_DETAIL      = "knowledge_content_detail kcd"
	TABLE_KNOWLEDGE_CONTENT_ATTACHMENT  = "knowledge_content_attachment kca"
	TABLE_KNOWLEDGE_CONTENT_QUESTION    = "knowledge_content_question kcq"
	TABLE_KNOWLEDGE_CONTENT_OPTION      = "knowledge_content_option kco"
	TABLE_KNOWLEDGE_CONTENT_LOG         = "knowledge_content_log kclo"
	TABLE_KNOWLEDGE_CONTENT_LOG_VERSION = "knowledge_content_log_version kclv"
	TABLE_HISTORY_KNOWLEDGE             = "history_knowledge hk"

	//Details
	TABLE_KNOWLEDGE_WORK_DETAIL               = "knowledge_work_detail kwd"
	TABLE_KNOWLEDGE_CONTENT_BOOKMARK          = "knowledge_content_bookmark kcb"
	TABLE_KNOWLEDGE_CONTENT_REPORT            = "knowledge_content_report kcr"
	TABLE_KNOWLEDGE_CONTENT_FEEDBACK          = "knowledge_content_feedback kcf"
	TABLE_KNOWLEDGE_FEEDBACK                  = "knowledge_feedback kf"
	TABLE_KNOWLEDGE_RELATION                  = "knowledge_relation kr"
	TABLE_DISCOVERY_TICKET_RELATION           = "discovery_ticket_relation dtr"
	TABLE_KNOWLEDGE_RELATION_TO_TICKET        = "knowledge_relation_to_ticket krt"
	TABLE_KNOWLEDGE_CONTENT_REPORT_ATTACHMENT = "knowledge_content_report_attachment kcra"
	TABLE_WORKDETAIL_KNOWLEDGE_MANAGEMENT     = "workdetail_knowledge_management wkm"
	TABLE_WORK_DETAIL_HAS_DOCUMENT            = "work_detail_has_document wdhd"

	//CI Relation
	TABLE_CI_TYPE                          = "ci_type ct"
	TABLE_CI_RELATION_TYPE                 = "ci_relation_type crt"
	TABLE_DISCOVERY_MACHINE_LOG            = "discovery_machine_log dml"
	TABLE_DISCOVERY_STORAGE_LOG            = "discovery_storage_log dsl"
	TABLE_DISCOVERY_PACKAGE_LOG            = "discovery_package_log dpl"
	TABLE_KNOWLEDGE_CONTENT_UPDATE_REQUEST = "knowledge_content_update_request kcur"
)

var Status = map[string]int{
	"IN PROGRESS":      1,
	"DRAFT":            2,
	"SME REVIEW":       3,
	"PUBLISH APPROVAL": 4,
	"PUBLISHED":        5,
	"RETIRE APPROVAL":  6,
	"RETIRED":          7,
	"CLOSED VERSION":   8,
	"CANCELLED":        9,
}

func RemoveAliasFromTable(table string) string {
	return strings.Split(table, " ")[0]
}

func GetCurrentStatusNumber(status string) int {
	return Status[status]
}

func GetNextStatus(i int) string {
	var nextStatus string
	for k, _ := range Status {
		if Status[k] == i {
			nextStatus = k
		}
	}
	return nextStatus
}
