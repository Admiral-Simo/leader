package api

import (
	"context"
	"fmt"
	"log"
	"server/api/handler"
	"server/api/middleware"
	"server/engines"
	"server/store"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type API struct {
	port   int
	store  *store.Queries
	engine engines.SearchEngine
	app    *gin.Engine
}

type APIBuilder struct {
	port   int
	dbURL  string
	store  *store.Queries
	engine engines.SearchEngine
	app    *gin.Engine
}

func NewAPIBuilder() *APIBuilder {
	return &APIBuilder{
		port:   4000,
		engine: engines.NewGoogleSearchEngine(),
		dbURL:  "postgres://postgres:strongpassword@db:5432/leader?sslmode=disable",
	}
}

func (b *APIBuilder) WithPort(port int) *APIBuilder {
	b.port = port
	return b
}

func (b *APIBuilder) WithDBURL(dbURL string) *APIBuilder {
	b.dbURL = dbURL
	return b
}

func (b *APIBuilder) WithSearchEngine(engine engines.SearchEngine) *APIBuilder {
	b.engine = engine
	return b
}

func (b *APIBuilder) Build() *API {
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, b.dbURL)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	queries := store.New(dbpool)

	return &API{
		port:   b.port,
		store:  queries,
		engine: b.engine,
		app:    gin.Default(),
	}
}

func (a *API) StartServer() {
	a.registerMiddlewares()
	a.registerRoutes()
	port := fmt.Sprintf(":%d", a.port)

	// serve static files in public/ directory
	a.app.Static("/public", "./public")

	a.app.Run(port)
}

func (a *API) registerMiddlewares() {
	h := handler.Handler{
		Store:        a.store,
		SearchEngine: a.engine,
		App:          a.app,
	}
	a.app.Use(middleware.JWTAuthMiddleware(h))
}
