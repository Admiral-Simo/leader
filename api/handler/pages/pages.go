package pages

import (
	"server/api/handler"
	"server/api/web/pages/authtempl"
	"server/api/web/pages/maintempl"

	"github.com/gin-gonic/gin"
)

func Routes(h handler.Handler) {
	// home
	h.App.GET("/", func(ctx *gin.Context) {
		maintempl.Home().Render(ctx, ctx.Writer)
	})

	// authentication
	h.App.GET("/signup", func(ctx *gin.Context) {
		authtempl.SignUp().Render(ctx, ctx.Writer)
	})

	h.App.GET("/login", func(ctx *gin.Context) {
		authtempl.Login().Render(ctx, ctx.Writer)
	})

	h.App.GET("/services", func(ctx *gin.Context) {
		maintempl.Services().Render(ctx, ctx.Writer)
	})

	h.App.GET("/about", func(ctx *gin.Context) {
		maintempl.About().Render(ctx, ctx.Writer)
	})

	h.App.GET("/contact", func(ctx *gin.Context) {
		maintempl.Contact().Render(ctx, ctx.Writer)
	})
}
