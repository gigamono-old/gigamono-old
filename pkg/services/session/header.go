package session

import "github.com/gin-gonic/gin"

// SetCustomHeader sets a custom header.
func SetCustomHeader(ctx *gin.Context, key string, value string) {
	ctx.Header("X-"+key, value) // X prefix for a custom header.
}
