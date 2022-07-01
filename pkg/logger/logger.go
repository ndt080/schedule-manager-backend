package logger

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
)

func GetLoggerHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		var reqMethod = c.Request.Method
		var reqUri = c.Request.RequestURI
		var statusCode = c.Writer.Status()
		var clientIP = c.ClientIP()

		log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
		log.SetFormatter(&log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})

		log.Infof("| %3d | %15s | %s | %s |", statusCode, clientIP, reqMethod, reqUri)
	}
}

func InitStdoutLogger() {
	l := log.New()
	l.Out = os.Stdout
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
}
