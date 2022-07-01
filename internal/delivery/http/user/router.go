package user

import (
	"github.com/gin-gonic/gin"
)

func (handler *Handler) InitRoutes(user *gin.RouterGroup) {
	user.GET("/user/me", handler.GetCurrentUser)
	user.GET("/user/:id", handler.GetUserById)
	user.GET("/users", handler.GetUsers)
}
