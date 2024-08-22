package api

import (
	"server/api/handler"
	"server/api/handler/auth"
	"server/api/handler/pages"
)

// StartServer registers the server
func (a *API) registerRoutes() {
	h := handler.Handler{
		Store:        a.store,
		SearchEngine: &a.engine,
        App: a.app,
	}
	auth.Routes(h)
	pages.Routes(h)
}
