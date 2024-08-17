package api

import (
	"context"
	"fmt"
	"log"
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

// APIBuilder provides a builder for the API struct
type APIBuilder struct {
	port   int
	dbURL  string
	store  *store.Queries
	engine engines.SearchEngine
	app    *gin.Engine
}

// NewAPIBuilder creates a new instance of APIBuilder
func NewAPIBuilder() *APIBuilder {
	return &APIBuilder{
		port:   4000,                                                                      // default port
		dbURL:  "postgres://postgres:my_password@localhost:5432/sqlctest?sslmode=disable", // default DB URL
		engine: engines.NewGoogleSearchEngine(),                                           // default engine
		app:    gin.Default(),                                                             // default engine
	}
}

// WithPort sets the port for the API
func (b *APIBuilder) WithPort(port int) *APIBuilder {
	b.port = port
	return b
}

// WithDBURL sets the database URL for the API
func (b *APIBuilder) WithDBURL(dbURL string) *APIBuilder {
	b.dbURL = dbURL
	return b
}

// WithSearchEngine sets the search engine for the API
func (b *APIBuilder) WithSearchEngine(engine engines.SearchEngine) *APIBuilder {
	b.engine = engine
	return b
}

// Build constructs the API instance with the specified configuration
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
		app:    b.app,
	}
}

// StartServer starts the server (implementation to be added)
func (a *API) StartServer() {
	a.registerMiddlewares()
	a.registerRoutes()
	port := fmt.Sprintf(":%d", a.port)
	fmt.Println("starting server at port ", port)
	// Server logic goes here
}

// StartServer registers the server (implementation to be added)
func (a *API) registerMiddlewares() {
}
