package http

import (
	"github.com/gin-gonic/gin"
	_ "github.com/ndt080/schedule-manager-backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(handler.preflight)

	handler.createRootRouterGroup(router)
	handler.createAuthRouterGroup(router)
	handler.createUserRouterGroup(router)
	handler.createWorkspaceRouterGroup(router)
	return router
}

func (handler *Handler) createRootRouterGroup(router *gin.Engine) {
	root := router.Group("/")

	root.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	root.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "alive"})
	})
}

func (handler *Handler) createAuthRouterGroup(router *gin.Engine) {
	group := router.Group("/auth")
	handler.Auth.InitRoutes(group)

	groupWithIdentity := router.Group("/auth", handler.service.Authorization.IdentifyUser)
	handler.Auth.InitRoutesWithIdentifyUser(groupWithIdentity)
}

func (handler *Handler) createUserRouterGroup(router *gin.Engine) {
	group := router.Group("/", handler.service.Authorization.IdentifyUser)
	handler.User.InitRoutes(group)
}

func (handler *Handler) createWorkspaceRouterGroup(router *gin.Engine) {
	group := router.Group("/", handler.service.Authorization.IdentifyUser)
	handler.Workspace.InitRoutes(group)
	handler.WorkspaceQueue.InitRoutes(group)

}
