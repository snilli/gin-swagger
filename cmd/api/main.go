package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-contrib/graceful"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/snilli/ormprovider"
	"go.uber.org/fx"

	"gin-swagger-api/config"
	_ "gin-swagger-api/docs"
	"gin-swagger-api/internal/handler"
	"gin-swagger-api/internal/handler/orderhdl"
	"gin-swagger-api/internal/handler/producthdl"
	"gin-swagger-api/internal/handler/userhdl"
	portorderrepo "gin-swagger-api/internal/port/repository/orderrepo"
	portproductrepo "gin-swagger-api/internal/port/repository/productrepo"
	portuserrepo "gin-swagger-api/internal/port/repository/userrepo"
	portordersvc "gin-swagger-api/internal/port/service/ordersvc"
	portproductsvc "gin-swagger-api/internal/port/service/productsvc"
	portusersvc "gin-swagger-api/internal/port/service/usersvc"
	"gin-swagger-api/internal/repository/orderrepo"
	"gin-swagger-api/internal/repository/productrepo"
	"gin-swagger-api/internal/repository/userrepo"
	"gin-swagger-api/internal/service/ordersvc"
	"gin-swagger-api/internal/service/productsvc"
	"gin-swagger-api/internal/service/usersvc"
)

// @title Gin Swagger API
// @version 1.0
// @description API documentation for Gin Swagger API service with Ent ORM
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /api/v1
// @schemes http https
func main() {
	fx.New(
		// Provide config
		fx.Provide(provideConfig),

		// Provide database
		fx.Provide(provideDatabase),

		// Provide repositories
		fx.Provide(
			fx.Annotate(
				userrepo.New,
				fx.As(new(portuserrepo.Repository)),
			),
			fx.Annotate(
				productrepo.New,
				fx.As(new(portproductrepo.Repository)),
			),
			fx.Annotate(
				orderrepo.New,
				fx.As(new(portorderrepo.Repository)),
			),
		),

		// Provide services
		fx.Provide(
			fx.Annotate(
				usersvc.New,
				fx.As(new(portusersvc.Service)),
			),
			fx.Annotate(
				productsvc.New,
				fx.As(new(portproductsvc.Service)),
			),
			fx.Annotate(
				ordersvc.New,
				fx.As(new(portordersvc.Service)),
			),
		),

		// Provide handlers
		fx.Provide(
			handler.NewSystemHandler,
			userhdl.NewHandler,
			producthdl.NewHandler,
			orderhdl.NewHandler,
		),

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
		Msg("Starting server")

	gin.SetMode(cfg.ServerMode)

	return cfg, nil
}

// provideDatabase creates database connection
func provideDatabase(lc fx.Lifecycle, cfg *config.Config) (*ormprovider.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db, err := ormprovider.NewClient(ctx, ormprovider.Config{
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
			return db.Close()
		},
	})

	return db, nil
}

// provideGinEngine creates and configures Gin engine
func provideGinEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	return r
}

// runServer sets up routes and starts the server
func runServer(
	lc fx.Lifecycle,
	cfg *config.Config,
	r *gin.Engine,
	systemHandler *handler.SystemHandler,
	userHandler *userhdl.Handler,
	productHandler *producthdl.Handler,
	orderHandler *orderhdl.Handler,
) {
	// Register routes
	systemHandler.RegisterRoutes(r)

	v1 := r.Group("/api/v1")
	{
		userHandler.RegisterRoutes(v1)
		productHandler.RegisterRoutes(v1)
		orderHandler.RegisterRoutes(v1)
	}

	log.Info().
		Str("url", fmt.Sprintf("http://%s", cfg.ServerAddr())).
		Msg("Server is ready")

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
