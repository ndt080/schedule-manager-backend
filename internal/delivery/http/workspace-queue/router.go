package workspace_queue

import (
	"github.com/gin-gonic/gin"
)

func (handler *WorkspaceQueueHandler) InitRoutes(group *gin.RouterGroup) {
	group.POST("/workspace/:id/queue", handler.CreateWorkspaceQueue)
	group.POST("/workspace/queue/:id/join", handler.JoinToWorkspaceQueue)
	group.DELETE("/workspace/queue/:id", handler.RemoveWorkspaceQueue)
	group.DELETE("/workspace/queue/:id/leave/:uid", handler.LeaveWorkspaceQueue)
}
