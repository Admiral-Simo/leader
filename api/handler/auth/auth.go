package auth

import (
	"server/api/handler"

	"github.com/gin-gonic/gin"
)

func Routes(h handler.Handler) {
	h.App.POST("/signup", func(ctx *gin.Context) {
	})
	h.App.POST("/login", func(ctx *gin.Context) {
	})
}
