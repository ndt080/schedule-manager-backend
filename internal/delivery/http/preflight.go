package http

import "github.com/gin-gonic/gin"

func (handler *Handler) preflight(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept-Encoding, X-CSRF-Token, User, accept, origin, Cache-Control, X-Requested-With")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
	}
	ctx.Next()
}
