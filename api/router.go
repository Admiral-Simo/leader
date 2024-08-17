package api

import (
	"server/api/handler"
	"server/api/handler/auth"
)

// StartServer registers the server (implementation to be added)
func (a *API) registerRoutes() {
	h := handler.Handler{
		Store:        a.store,
		SearchEngine: &a.engine,
        App: a.app,
	}
	auth.Routes(h)
}
