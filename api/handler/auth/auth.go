package auth

import (
	"server/api/handler"
	"server/api/web/layout"

	"github.com/gin-gonic/gin"
)

func Routes(h handler.Handler) {
	h.App.GET("/signup", func(ctx *gin.Context) {
		layout.Base("signup").Render(ctx, ctx.Writer)
	})
	h.App.GET("/login", func(ctx *gin.Context) {
		layout.Base("login").Render(ctx, ctx.Writer)
	})
}
