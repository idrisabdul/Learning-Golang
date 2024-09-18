package apps

import (
	"os"
	"sygap_new_knowledge_management/backend/middleware"
	"sygap_new_knowledge_management/backend/routes"
	"sygap_new_knowledge_management/config/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func StartApps() {
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Set("Access-Control-Allow-Credentials", "true")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.Next()
	})

	app.Use(middleware.CustomRecoverMiddleware)

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  "2006/01/02 15:04:05",
		DisableTimestamp: false,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
			logrus.FieldKeyFunc:  "@caller",
		},
	})

	log := logrus.New()
	log.SetOutput(os.Stdout)

	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		logrus.Fatalf("Error loading .env file: %v", errLoadEnv)
	}

	mysql := setupMySQLConnection()

	KMRoutes := routes.KMRoutes{
		App:     app,
		List:    setupList(mysql, log),
		CRUD:    setupCRUD(mysql, log),
		History: setupHistory(mysql, log),
	}
	KMRoutes.SetupKMRoutes()

	DocumentRoutes := routes.DocumentRoutes{
		App:      app,
		Document: setupDocumentCRUD(mysql, log),
	}
	DocumentRoutes.SetupDocumentRoutes()

	OptionsRoute := routes.Options{
		App:                 app,
		OpCatHandler:        SetupOpCat(mysql, log),
		Product:             SetupProduct(mysql, log),
		ProductType:         SetupProductType(mysql, log),
		Organization:        SetupOrganization(mysql, log),
		Company:             SetupCompany(mysql, log),
		Expertee:            setupExpertees(mysql, log),
		ContentType:         setupContentType(mysql, log),
		Status:              setupStatus(mysql, log),
		Symptoms:            setupSymptoms(mysql, log),
		UpdateRequestStatus: setupRequestUpdateStatus(mysql, log),
		WorkDetailType:      setupWorkDetailType(mysql, log),
		RelationType:        setupRelationType(mysql, log),
	}
	OptionsRoute.SetupOptions()

	// Update Request
	updateRequestRoute := routes.UpdateRequestRoutes{
		App:  app,
		CRUD: setupUpdateRequest(mysql, log),
	}
	updateRequestRoute.SetupUpdateRequestRoutes()

	// Search
	SearchRoute := routes.Search{
		App:            app,
		SearchHandler:  SetupSearch(mysql, log),
		HistoryHandler: SetupHistory(mysql, log),
	}
	SearchRoute.SearchRoute()

	// Work Detail
	WorkDetailRoute := routes.WorkDetailRoute{
		App:            app,
		WorkDetailHdlr: setupWorkDetail(mysql, log),
	}

	WorkDetailRoute.SetupWorkDetail()

	//HistoryKM
	HistoryRoute := routes.HistoryRoute{
		App:         app,
		HistoryHdlr: setupHistoryListKM(mysql, log),
	}

	HistoryRoute.HistoryRoute()

	//Feedback
	FeedbackRoute := routes.FeedbackRoute{
		App:          app,
		FeedbackHdlr: setupFeedback(mysql, log),
	}

	FeedbackRoute.SetupFeedback()

	// Relation
	RelationRoute := routes.RelationRoute{
		App:          app,
		RelationHdlr: setupRelation(mysql, log),
	}
	RelationRoute.SetupRelation()

	// CI Relation
	CiRelationRoute := routes.CiRelationRoute{
		App:            app,
		CiRelationHdlr: setupCiRelation(mysql, log),
	}
	CiRelationRoute.SetupCiRelation()

	// knowledge relation
	KnowledgeRelationRoute := routes.KnowledgeRelationRoute{
		App:                   app,
		KnowledgeRelationHdlr: setupKnowledgeRelation(mysql, log),
	}
	KnowledgeRelationRoute.SetupKnowledgeRelation()

	errApp := app.Listen(":8080")
	if errApp != nil {
		logrus.Fatalf("Error starting Fiber app: %v", errApp)
	}
}

func setupMySQLConnection() *gorm.DB {
	hostMysql := os.Getenv("HOST_MYSQL")
	usernameMysql := os.Getenv("USSERNAME_MYSQL")
	passwordMysql := os.Getenv("PASSWORD_MYSQL")
	dbMysql := os.Getenv("DB_MYSQL")

	return db.MysqlConnect(hostMysql, usernameMysql, passwordMysql, dbMysql)
}
