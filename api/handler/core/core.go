package core

import (
	"fmt"
	"server/api/handler"
	"server/api/web/pages/maintempl"
	"server/store"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func Routes(h handler.Handler) {
	h.App.POST("/emails", func(ctx *gin.Context) {
		keyword := ctx.PostForm("keyword")
		if len(keyword) < 3 {
			err := fmt.Errorf("the keyword is too short.")
			maintempl.Home(err).Render(ctx, ctx.Writer)
			return
		}
		// engine = massive energy
		user := ctx.Value("user").(store.User)

		if user.Credits.Int32 < 5 {
			err := fmt.Errorf("please fill your credit points.")
			maintempl.Home(err).Render(ctx, ctx.Writer)
			return
		}

		result := h.SearchEngine.GetEmailsAddressByQuery(keyword)
		createHistory(user.ID, keyword, result, h, ctx)
		fmt.Println("user credits left:")
		maintempl.EmailList(result).Render(ctx, ctx.Writer)
	})
}

func createHistory(userID int64, keyword string, result map[string]map[string]struct{}, h handler.Handler, ctx *gin.Context) {
	if len(result) == 0 {
		return
	}

	searchHistory, _ := h.Store.CreateSearchHistory(ctx, store.CreateSearchHistoryParams{
		UserID:  pgtype.Int8{Int64: userID, Valid: true},
		Keyword: keyword,
	})

	h.Store.SubtractCreditsByUserID(ctx, userID)

	for website, emails := range result {
		webRef, _ := h.Store.CreateWebsite(ctx, store.CreateWebsiteParams{
			SearchHistoryID: pgtype.Int8{Int64: searchHistory.ID, Valid: true},
			Url:             website,
		})
		for email := range emails {
			h.Store.CreateEmail(ctx, store.CreateEmailParams{
				WebsiteID: pgtype.Int8{Int64: webRef.ID, Valid: true},
				Email:     email,
				Date:      pgtype.Timestamptz{Time: time.Now()},
			})
		}
	}
}
