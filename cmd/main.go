package main

import (
	"context"
	"fmt"
	"time"

	ormprovider "github.com/example/orm-provider-api"
	"github.com/gin-contrib/graceful"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"gin-swagger-api/config"
	_ "gin-swagger-api/docs"
	"gin-swagger-api/internal/handler"
	"gin-swagger-api/internal/handler/orderhdl"
	"gin-swagger-api/internal/handler/producthdl"
	"gin-swagger-api/internal/handler/userhdl"
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
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	cfg.SetupLogger()

	log.Info().
		Str("mode", cfg.ServerMode).
		Str("address", cfg.ServerAddr()).
		Msg("Starting server")

	gin.SetMode(cfg.ServerMode)

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
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer db.Close()

	log.Info().
		Str("host", cfg.DatabaseHost).
		Str("database", cfg.DatabaseName).
		Msg("Connected to database successfully")

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	userRepo := userrepo.New(db)
	productRepo := productrepo.New(db)
	orderRepo := orderrepo.New(db)

	userService := usersvc.New(userRepo)
	productService := productsvc.New(productRepo)
	orderService := ordersvc.New(orderRepo)

	systemHandler := handler.NewSystemHandler(cfg)
	userHandler := userhdl.NewHandler(userService)
	productHandler := producthdl.NewHandler(productService)
	orderHandler := orderhdl.NewHandler(orderService)

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

	serverCtx := context.Background()
	if err := srv.RunWithContext(serverCtx); err != nil {
		log.Fatal().Err(err).Msg("Server error")
	}

	log.Info().Msg("Server shutdown gracefully")
}
