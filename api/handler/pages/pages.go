package pages

import (
	"fmt"
	"net/http"
	"server/api/handler"
	"server/api/web/pages/authtempl"
	"server/api/web/pages/maintempl"
	"server/store"

	"github.com/gin-gonic/gin"
)

func Routes(h handler.Handler) {
	// home
	h.App.GET("/", func(ctx *gin.Context) {
		maintempl.Home().Render(ctx, ctx.Writer)
	})

	// authentication
	h.App.GET("/signup", func(ctx *gin.Context) {
		authtempl.SignUp(nil).Render(ctx, ctx.Writer)
	})

	h.App.GET("/login", func(ctx *gin.Context) {
		authtempl.Login(nil).Render(ctx, ctx.Writer)
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

	PostContactData(h)
}

func PostContactData(h handler.Handler) {
	h.App.POST("/contact", func(ctx *gin.Context) {
		email := ctx.PostForm("email")
		name := ctx.PostForm("name")
		message := ctx.PostForm("message")
		_, err := h.Store.CreateMessage(ctx, store.CreateMessageParams{
			Name:    name,
			Email:   email,
			Message: message,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		// redirect to "/"
		ctx.Redirect(http.StatusSeeOther, "/")
	})
}
