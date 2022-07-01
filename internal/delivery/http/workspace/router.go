package workspace

import (
	"github.com/gin-gonic/gin"
)

func (handler *Handler) InitRoutes(group *gin.RouterGroup) {
	group.POST("/workspace", handler.CreateWorkspace)
	group.PUT("/workspace", handler.UpdateWorkspace)

	group.GET("/workspace/:id", handler.GetWorkspaceById)
	group.DELETE("/workspace/:id", handler.RemoveWorkspace)

	group.GET("/workspace/:id/invite", handler.GetInviteToken)
	group.POST("/workspace/:id/task", handler.CreateWorkspaceTask)
	group.POST("/workspace/:id/member", handler.CreateWorkspaceMember)
	group.DELETE("/workspace/task/:id", handler.RemoveWorkspaceTask)

	group.GET("/workspaces/search", handler.SearchWorkspaces)
	group.GET("/workspaces/me", handler.GetUserWorkspaces)
}
