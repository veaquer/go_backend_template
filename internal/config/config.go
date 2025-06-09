package config

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Env        string      `mapstructure:"ENV"`
	BackendUrl string      `mapstructure:"BACKEND_URL"`
	Port       string      `mapstructure:"PORT"`
	DSN        string      `mapstructure:"DSN"`
	Redis      Redis       `mapstructure:",squash"`
	JWT        JWTConfig   `mapstructure:",squash"`
	Email      EmailConfig `mapstructure:",squash"`
}

type Redis struct {
	Addr     string `mapstructure:"REDIS_ADDR"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"REDIS_DB"`
}

type JWTConfig struct {
	RefreshSecret     string        `mapstructure:"JWT_REFRESH_SECRET"`
	AccessSecret      string        `mapstructure:"JWT_ACCESS_SECRET"`
	RefreshExpiration time.Duration `mapstructure:"JWT_REFRESH_EXPIRATION"`
	AccessExpiration  time.Duration `mapstructure:"JWT_ACCESS_EXPIRATION"`
}

type EmailConfig struct {
	From     string `mapstructure:"EMAIL_FROM"`
	Host     string `mapstructure:"EMAIL_HOST"`
	Port     int    `mapstructure:"EMAIL_PORT"`
	Username string `mapstructure:"EMAIL_USER"`
	Password string `mapstructure:"EMAIL_PASS"`
}

func Load(logger *zap.Logger) *Config {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	if err := godotenv.Load(); err != nil {
		logger.Warn(".env file not found or failed to load, proceeding with system env")
	}
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("Failed to read config", zap.Error(err))
		return nil
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		logger.Fatal("Failed to unmarshal config", zap.Error(err))
		return nil
	}
	return &config
}
