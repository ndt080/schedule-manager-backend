package auth

import (
	"github.com/gin-gonic/gin"
)

func (handler *Handler) InitRoutes(auth *gin.RouterGroup) {
	auth.POST("/sign-in", handler.SignIn)
	auth.POST("/sign-up", handler.SignUp)
	auth.POST("/refresh-token", handler.RefreshToken)
	auth.GET("/verify-email", handler.VerifyEmail)
	auth.POST("/confirm-email-again", handler.ConfirmEmailAgain)
}

func (handler *Handler) InitRoutesWithIdentifyUser(auth *gin.RouterGroup) {
	auth.GET("/status", handler.CheckAuthStatus)
}
