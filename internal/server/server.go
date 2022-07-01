package server

import (
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/ndt080/schedule-manager-backend/internal/configs"
	httpHandler "github.com/ndt080/schedule-manager-backend/internal/delivery/http"
	"github.com/ndt080/schedule-manager-backend/internal/delivery/smtp"
	"github.com/ndt080/schedule-manager-backend/internal/repository"
	"github.com/ndt080/schedule-manager-backend/internal/service"
	"github.com/ndt080/schedule-manager-backend/pkg/auth"
	"github.com/ndt080/schedule-manager-backend/pkg/logger"
	"github.com/ndt080/schedule-manager-backend/pkg/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	instance *manners.GracefulServer
}

func (server *Server) Run() error {
	logger.InitStdoutLogger()
	config := configs.NewServerConfig("configs")

	router := inject(config)
	server.instance = manners.NewWithServer(&http.Server{
		Addr:           ":" + configs.GetServerPort(config, "PORT"),
		Handler:        router,
		MaxHeaderBytes: config.Http.MaxHeaderBytes,
		ReadTimeout:    config.Http.ReadTimeout,
		WriteTimeout:   config.Http.WriteTimeout,
	})

	log.Infoln("Server is running")
	return server.instance.ListenAndServe()
}

func (server *Server) Shutdown() {
	log.Infoln("Shutting down")
	server.instance.Close()
}

func inject(serverConfig *configs.ServerConfig) *gin.Engine {
	var smtpConfig configs.SmtpConfig
	if err := envconfig.Process("", &smtpConfig); err != nil {
		log.Fatalf(err.Error())
	}

	dbConnection, err := createDatabaseConnection("postgres")
	if err != nil {
		log.Fatalf(err.Error())
	}

	tokenManager, err := auth.NewManager(
		serverConfig.Auth.SigningKey,
		serverConfig.Auth.AccessTokenTTL,
		serverConfig.Auth.RefreshTokenTTL,
	)
	if err != nil {
		log.Fatalf(err.Error())
	}

	smtpServiceInstance := smtp.NewSmtpService(smtpConfig)
	repositoryInstance := repository.NewRepository(dbConnection)
	serviceInstance := service.NewService(repositoryInstance, smtpServiceInstance, tokenManager)
	handlerInstance := httpHandler.NewHandler(serviceInstance, tokenManager)

	router := handlerInstance.InitRoutes()
	router.Use(logger.GetLoggerHandlerFunc())
	router.Use(middleware.CORS())
	return router
}

func createDatabaseConnection(driverName string) (*sqlx.DB, error) {
	var databaseConfig configs.DatabaseConfig
	if err := envconfig.Process("", &databaseConfig); err != nil {
		return nil, err
	}

	dbConnection, err := sqlx.Open(driverName, databaseConfig.ToString())
	if err != nil {
		return nil, err
	}

	if err = dbConnection.Ping(); err != nil {
		return nil, err
	}

	return dbConnection, nil
}
