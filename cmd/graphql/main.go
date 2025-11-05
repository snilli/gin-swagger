package main

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/graceful"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/snilli/ormprovider"
	"github.com/snilli/ormprovider/ent"
	entgraphql "github.com/snilli/ormprovider/ent/graphql"
	"go.uber.org/fx"

	"gin-swagger-api/config"
)

func main() {
	fx.New(
		// Provide config
		fx.Provide(provideConfig),

		// Provide database (ORM client wrapper)
		fx.Provide(provideDatabase),

		// Provide Ent client (extract from wrapper)
		fx.Provide(provideEntClient),

		// Provide GraphQL resolver (using Ent's built-in resolver)
		fx.Provide(entgraphql.NewResolver),

		// Provide GraphQL server
		fx.Provide(provideGraphQLServer),

		// Provide Gin engine
		fx.Provide(provideGinEngine),

		// Invoke server startup
		fx.Invoke(runServer),
	).Run()
}

// provideConfig loads and sets up configuration
func provideConfig() (*config.Config, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	cfg.SetupLogger()

	log.Info().
		Str("mode", cfg.ServerMode).
		Str("address", cfg.ServerAddr()).
		Msg("Starting GraphQL server")

	gin.SetMode(cfg.ServerMode)

	return cfg, nil
}

// provideDatabase creates database connection and returns Ent client
func provideDatabase(lc fx.Lifecycle, cfg *config.Config) (*ormprovider.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := ormprovider.NewClient(ctx, ormprovider.Config{
		Host:     cfg.DatabaseHost,
		Port:     cfg.DatabasePort,
		User:     cfg.DatabaseUser,
		Password: cfg.DatabasePassword,
		DBName:   cfg.DatabaseName,
		SSLMode:  cfg.DatabaseSSLMode,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Info().
		Str("host", cfg.DatabaseHost).
		Str("database", cfg.DatabaseName).
		Msg("Connected to database successfully")

	// Register lifecycle hooks
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("Closing database connection")
			return client.Close()
		},
	})

	return client, nil
}

// provideEntClient extracts *ent.Client from *ormprovider.Client
func provideEntClient(client *ormprovider.Client) *ent.Client {
	return client.Client
}

// provideGraphQLServer creates GraphQL server with Ent resolver
func provideGraphQLServer(r *entgraphql.Resolver) *handler.Server {
	srv := handler.NewDefaultServer(
		entgraphql.NewExecutableSchema(
			entgraphql.Config{Resolvers: r},
		),
	)

	log.Info().Msg("GraphQL server initialized")
	return srv
}

// provideGinEngine creates and configures Gin engine
func provideGinEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	return r
}

// graphqlHandler wraps GraphQL handler for Gin
func graphqlHandler(srv *handler.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

// playgroundHandler wraps GraphQL playground for Gin
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL Playground", "/graphql")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// runServer sets up routes and starts the server
func runServer(
	lc fx.Lifecycle,
	cfg *config.Config,
	r *gin.Engine,
	graphqlServer *handler.Server,
) {
	// GraphQL endpoint
	a := graphqlHandler(graphqlServer)
	r.POST("/graphql", a)
	r.GET("/graphql", a)

	// GraphQL Playground (only in development)
	if cfg.ServerMode != "release" {
		r.GET("/playground", playgroundHandler())
		log.Info().
			Str("url", fmt.Sprintf("http://%s/playground", cfg.ServerAddr())).
			Msg("GraphQL Playground available")
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Info().
		Str("url", fmt.Sprintf("http://%s/graphql", cfg.ServerAddr())).
		Msg("GraphQL server is ready")

	srv, err := graceful.New(r, graceful.WithAddr(fmt.Sprintf(":%s", cfg.ServerPort)))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create graceful server")
	}

	// Register lifecycle hooks
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.RunWithContext(context.Background()); err != nil {
					log.Error().Err(err).Msg("Server error")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("Server shutdown gracefully")
			return srv.Shutdown(ctx)
		},
	})
}
