package auth

import (
	"server/api/handler"

	"github.com/gin-gonic/gin"
)

func Routes(h handler.Handler) {
	h.App.GET("/signup", func(ctx *gin.Context) {
	})
	h.App.GET("/login", func(ctx *gin.Context) {})
}
