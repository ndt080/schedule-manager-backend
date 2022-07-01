package main

import (
	"github.com/gin-gonic/gin"
	s "github.com/ndt080/schedule-manager-backend/internal/server"
	log "github.com/sirupsen/logrus"
)

// @title           Schedule Manager Swagger API
// @version         1.0
// @description     Swagger API for Golang Project Schedule Manager.
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @contact.name   API Support
// @contact.email  andreipetrov080@gmail.com

// @securityDefinitions.apikey  AuthorizationKey
// @in                          header
// @name                        Authorization
func main() {
	gin.SetMode(gin.DebugMode)

	server := &s.Server{}
	err := server.Run()

	if err != nil {
		log.Fatalf("Error while running server %s", err.Error())
	}
}
