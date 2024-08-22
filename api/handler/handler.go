package handler

import (
	"server/engines"
	"server/store"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Store        *store.Queries
	App          *gin.Engine
	SearchEngine engines.SearchEngine
}
