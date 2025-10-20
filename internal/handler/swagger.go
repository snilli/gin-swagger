package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gin-swagger-api/config"
)

type SystemHandler struct {
	config *config.Config
}

func NewSystemHandler(cfg *config.Config) *SystemHandler {
	return &SystemHandler{
		config: cfg,
	}
}

func (h *SystemHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/health", h.HealthCheck)

	if h.config.EnableSwagger {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		log.Info().
			Str("url", fmt.Sprintf("http://%s/swagger/index.html", h.config.ServerAddr())).
			Msg("Swagger documentation enabled")
	}
}

// HealthCheck godoc
// @Summary Health check
// @Description Check if the service is running
// @Tags system
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (h *SystemHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "service is running",
	})
}
