package config

import (
	"fmt"
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

// Load initializes configuration with elegant fallback mechanisms
func Load(logger *zap.Logger) *Config {
	// Step 1: Initialize Viper with automatic environment variable binding
	viper.AutomaticEnv()
	
	// Step 2: Set rational defaults to ensure system stability
	setDefaults()
	
	// Step 3: Attempt to load .env file (graceful failure)
	if err := godotenv.Load(); err != nil {
		logger.Info("No .env file found, utilizing system environment variables", 
			zap.String("reason", "This is the expected behavior in production environments"))
	} else {
		logger.Info("Successfully loaded .env file")
	}
	
	// Step 4: Configure Viper for optional config file reading
	viper.SetConfigFile(".env")
	
	// Step 5: Attempt to read config file (optional, non-fatal)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Info("No configuration file found, relying on environment variables and defaults")
		} else {
			logger.Warn("Configuration file found but failed to read", zap.Error(err))
		}
	} else {
		logger.Info("Using configuration file", zap.String("file", viper.ConfigFileUsed()))
	}
	
	// Step 6: Unmarshal configuration with proper error handling
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		logger.Fatal("Failed to unmarshal configuration - this indicates a structural problem", 
			zap.Error(err))
		return nil
	}
	
	// Step 7: Validate critical configuration values
	if err := validateConfig(&config); err != nil {
		logger.Fatal("Configuration validation failed", zap.Error(err))
		return nil
	}
	
	logger.Info("Configuration loaded successfully", 
		zap.String("env", config.Env),
		zap.String("port", config.Port))
	
	return &config
}

// setDefaults establishes rational baseline values
func setDefaults() {
	viper.SetDefault("ENV", "development")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("BACKEND_URL", "http://localhost:8080")
	
	// Redis defaults
	viper.SetDefault("REDIS_ADDR", "localhost:6379")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DB", 0)
	
	// JWT defaults (should be overridden in production)
	viper.SetDefault("JWT_ACCESS_EXPIRATION", "15m")
	viper.SetDefault("JWT_REFRESH_EXPIRATION", "168h") // 7 days
	
	// Email defaults
	viper.SetDefault("EMAIL_HOST", "localhost")
	viper.SetDefault("EMAIL_PORT", 587)
}

// validateConfig ensures configuration integrity
func validateConfig(config *Config) error {
	if config.JWT.AccessSecret == "" {
		return fmt.Errorf("JWT_ACCESS_SECRET is required for security")
	}
	
	if config.JWT.RefreshSecret == "" {
		return fmt.Errorf("JWT_REFRESH_SECRET is required for security")
	}
	
	if config.DSN == "" {
		return fmt.Errorf("DSN (database connection string) is required")
	}
	
	return nil
}

// LoadWithProfile allows environment-specific configuration loading
func LoadWithProfile(logger *zap.Logger, profile string) *Config {
	// Load base configuration
	config := Load(logger)
	
	// Override with profile-specific values if available
	profileFile := fmt.Sprintf(".env.%s", profile)
	if err := godotenv.Load(profileFile); err == nil {
		logger.Info("Loaded profile-specific configuration", 
			zap.String("profile", profile))
		
		// Re-unmarshal to pick up profile changes
		if err := viper.Unmarshal(config); err != nil {
			logger.Fatal("Failed to unmarshal profile configuration", zap.Error(err))
		}
	}
	
	return config
}

// GetConfigSummary provides a diagnostic view of loaded configuration
func GetConfigSummary(config *Config) map[string]any {
	return map[string]any{
		"environment":    config.Env,
		"port":          config.Port,
		"backend_url":   config.BackendUrl,
		"redis_addr":    config.Redis.Addr,
		"email_host":    config.Email.Host,
		"jwt_configured": config.JWT.AccessSecret != "" && config.JWT.RefreshSecret != "",
		"dsn_configured": config.DSN != "",
	}
}
