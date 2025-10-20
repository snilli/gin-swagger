package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-simpler.org/env"
)

type Config struct {
	ServerPort string `env:"SERVER_PORT" default:"8080"`
	ServerHost string `env:"SERVER_HOST" default:"0.0.0.0"`
	ServerMode string `env:"SERVER_MODE" default:"debug"`

	DatabaseHost     string `env:"DATABASE_HOST" default:"localhost"`
	DatabasePort     string `env:"DATABASE_PORT" default:"5432"`
	DatabaseUser     string `env:"DATABASE_USER" default:"postgres"`
	DatabasePassword string `env:"DATABASE_PASSWORD"`
	DatabaseName     string `env:"DATABASE_NAME" default:"myapp"`
	DatabaseSSLMode  string `env:"DATABASE_SSL_MODE" default:"disable"`

	JWTSecret     string `env:"JWT_SECRET" default:"your-secret-key-change-this"`
	JWTExpiration int    `env:"JWT_EXPIRATION" default:"24"`

	APIVersion    string `env:"API_VERSION" default:"v1"`
	APITimeout    int    `env:"API_TIMEOUT" default:"30"`
	RateLimitRPS  int    `env:"RATE_LIMIT_RPS" default:"100"`
	MaxUploadSize int64  `env:"MAX_UPLOAD_SIZE" default:"10485760"`

	RedisHost     string `env:"REDIS_HOST" default:"localhost"`
	RedisPort     string `env:"REDIS_PORT" default:"6379"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDB       int    `env:"REDIS_DB" default:"0"`

	EnableSwagger bool `env:"ENABLE_SWAGGER" default:"true"`
	EnableCORS    bool `env:"ENABLE_CORS" default:"true"`
	EnableMetrics bool `env:"ENABLE_METRICS" default:"false"`

	LogLevel  string `env:"LOG_LEVEL" default:"info"`
	LogFormat string `env:"LOG_FORMAT" default:"json"`
}

var (
	instance *Config
	once     sync.Once
)

func Load() (*Config, error) {
	var loadErr error

	once.Do(func() {
		_ = godotenv.Load()

		cfg := &Config{}

		if err := env.Load(cfg, nil); err != nil {
			loadErr = fmt.Errorf("failed to load environment variables: %w", err)
			return
		}

		if err := cfg.Validate(); err != nil {
			loadErr = err
			return
		}

		instance = cfg
	})

	if loadErr != nil {
		return nil, loadErr
	}

	return instance, nil
}

func Get() *Config {
	if instance == nil {
		panic("config not loaded, call Load() first")
	}
	return instance
}

func (c *Config) Validate() error {
	if c.ServerPort == "" {
		return fmt.Errorf("SERVER_PORT is required")
	}

	if c.JWTSecret == "" {
		return fmt.Errorf("JWT_SECRET must be set to a secure value")
	}

	if c.ServerMode != "debug" && c.ServerMode != "release" && c.ServerMode != "test" {
		return fmt.Errorf("SERVER_MODE must be one of: debug, release, test")
	}

	if c.LogLevel != "debug" && c.LogLevel != "info" && c.LogLevel != "warn" && c.LogLevel != "error" {
		return fmt.Errorf("LOG_LEVEL must be one of: debug, info, warn, error")
	}

	if c.LogFormat != "json" && c.LogFormat != "text" {
		return fmt.Errorf("LOG_FORMAT must be one of: json, text")
	}

	return nil
}

func (c *Config) DatabaseDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DatabaseHost,
		c.DatabasePort,
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseName,
		c.DatabaseSSLMode,
	)
}

func (c *Config) RedisAddr() string {
	return fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort)
}

func (c *Config) ServerAddr() string {
	return fmt.Sprintf("%s:%s", c.ServerHost, c.ServerPort)
}

func (c *Config) SetupLogger() {
	switch c.LogLevel {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if c.LogFormat == "text" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	} else {
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}
}
