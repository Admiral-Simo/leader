package core

import (
	"server/api/handler"
	"server/api/web/pages/maintempl"

	"github.com/gin-gonic/gin"
)

func Routes(h handler.Handler) {
	h.App.POST("/emails", func(ctx *gin.Context) {
		keyword := ctx.PostForm("keyword")
		result := h.SearchEngine.GetEmailsAddressByQuery(keyword)
		maintempl.EmailList(result).Render(ctx, ctx.Writer)
	})
}
