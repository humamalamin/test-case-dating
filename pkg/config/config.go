package config

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type ContextKey string

type Config struct {
	AppEnv                     string `mapstructure:"APP_ENV"`
	AppTz                      string `mapstructure:"APP_TZ"`
	DatabasePort               int    `mapstructure:"DATABASE_PORT"`
	DatabaseDriver             string `mapstructure:"DATABASE_DRIVER"`
	DatabasHost                string `mapstructure:"DATABASE_HOST"`
	DatabaseUser               string `mapstructure:"DATABASE_USER"`
	DatabasePassword           string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseName               string `mapstructure:"DATABASE_NAME"`
	DatabaseMaxOpenConnections int    `mapstructure:"DATABASE_MAX_OPEN_CONNECTION"`
	DatabaseMaxIdleConnections int    `mapstructure:"DATABASE_MAX_IDLE_CONNECTION"`
	JwtSigningKey              string `mapstructure:"JWT_SIGNING_KEY"`
	JwtIssuer                  string `mapstructure:"JWT_ISSUER"`
	JwtAccessTokenDuration     int    `mapstructure:"JWT_ACCESS_TOKEN_DURATION"`
	JwtRefreshTokenDuration    int    `mapstructure:"JWT_REFRESH_TOKEN_DURATION"`
	PortHttpServer             string `mapstructure:"SERVER_PORT"`
	ServerHTTPWriteTimeout     int    `mapstructure:"SERVER_WRITE_TIMEOUT"`
	ServerHTTPReadTimeout      int    `mapstructure:"SERVER_READ_TIMEOUT"`
	SaltPassword               string `mapstructure:"SALT_PASSWORD"`
}

func NewConfig() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SetConfigName(".env.example")

			if err := viper.ReadInConfig(); err != nil {
				log.Error().Err(err).Msg("failed to read config file")
				return nil, err
			}
		} else {
			log.Error().Err(err).Msg("failed to read config file")
			return nil, err
		}
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		log.Error().Err(err).Msg("failed to unmarshal config")
		return nil, err
	}

	return cfg, nil
}
