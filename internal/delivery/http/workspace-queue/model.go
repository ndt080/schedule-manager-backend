package workspace_queue

type WorkspaceQueueRequest struct {
	ID   int64  `db:"id" json:"id" binding:"omitempty"`
	Name string `db:"name" json:"name" binding:"required"`
}
