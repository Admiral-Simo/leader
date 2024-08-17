package auth

import (
	"server/api/handler"
	"server/api/web/pages/authtempl"

	"github.com/gin-gonic/gin"
)

func Routes(h handler.Handler) {
	h.App.GET("/signup", func(ctx *gin.Context) {
		authtempl.Login().Render(ctx, ctx.Writer)
	})
	h.App.GET("/login", func(ctx *gin.Context) {
		authtempl.SignUp().Render(ctx, ctx.Writer)
	})
}
