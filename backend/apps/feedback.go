package apps

import (
	feedback_hdl "sygap_new_knowledge_management/backend/handler/feedback"
	feedback_repo "sygap_new_knowledge_management/backend/repository/feedback"
	feedback_svc "sygap_new_knowledge_management/backend/services/feedback"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func setupFeedback(mysql *gorm.DB, logger *logrus.Logger) *feedback_hdl.FeedbackHandler {
	feedbackRepo := feedback_repo.NewFeedbackRepository(mysql, logger)
	feedbackService := feedback_svc.NewFeedbackService(feedbackRepo, logger)
	return feedback_hdl.NewFeedbackHandler(feedbackService, logger)
}
