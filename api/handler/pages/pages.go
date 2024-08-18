package pages

import (
	"net/http"
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

	h.App.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
}
