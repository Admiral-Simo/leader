package pages

import (
	"fmt"
	"net/http"
	"server/api/handler"
	"server/api/web/pages/authtempl"
	"server/api/web/pages/historytempl"
	"server/api/web/pages/maintempl"
	"server/store"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func Routes(h handler.Handler) {
	// home
	h.App.GET("/", func(ctx *gin.Context) {
		maintempl.Home(nil).Render(ctx, ctx.Writer)
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

	h.App.GET("/history", func(ctx *gin.Context) {
		history := getHistoryData(h, ctx)
		historytempl.HistoryPage(history).Render(ctx, ctx.Writer)
	})

	PostContactData(h)
}

func getHistoryData(h handler.Handler, ctx *gin.Context) handler.HistoryResult {
	var result handler.HistoryResult

	// Get the user ID from the context
	userID := ctx.Value("user").(store.User).ID

	// Fetch search history for the user
	history, err := h.Store.GetSearchHistoryByUserID(ctx, pgtype.Int8{Int64: userID, Valid: true})
	if err != nil {
		// Handle the error appropriately (e.g., log it, return an error response, etc.)
		return result
	}

	// Iterate over the search history items
	for _, item := range history {
		var websites []handler.Website

		// Fetch websites for each search history item
		websitesData, err := h.Store.GetWebsitesBySearchHistoryID(ctx, pgtype.Int8{Int64: item.ID, Valid: true})
		if err != nil {
			// Handle the error appropriately
			continue
		}

		// Iterate over the websites and fetch emails for each website
		for _, websiteData := range websitesData {
			emails, err := h.Store.GetEmailsByWebsiteID(ctx, pgtype.Int8{Int64: websiteData.ID, Valid: true})
			if err != nil {
				// Handle the error appropriately
				continue
			}

			// Append the website with its associated emails to the websites slice
			websites = append(websites, handler.Website{
				Url:    websiteData.Url,
				Emails: emails, // Assuming emails is of type []store.Email
			})
		}

		// Append the history item to the result
		result.HistoryItems = append(result.HistoryItems, handler.HistoryItem{
			Keyword:    item.Keyword,
			SearchTime: item.SearchTime.Time,
			Websites:   websites,
		})
	}

	return result
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
