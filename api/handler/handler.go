package handler

import (
	"server/engines"
	"server/store"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Store        *store.Queries
	App          *gin.Engine
	SearchEngine engines.SearchEngine
}

type HistoryResult struct {
	HistoryItems []HistoryItem
}

type HistoryItem struct {
	Keyword    string
	SearchTime time.Time
	Websites   []Website
}

type Website struct {
    Url string
	Emails []store.Email
}
