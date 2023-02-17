package configs

import "github.com/spf13/viper"

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Host string
	Port int
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// Global config
var cfg *config

// Init default config
func init() {
	// Default api config
	viper.SetDefault("api.host", "localhost")
	viper.SetDefault("api.port", "8080")

	// Default database config
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

// Load config from TOML file
func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Host: viper.GetString("api.host"),
		Port: viper.GetInt("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

func GetDBConfig() DBConfig {
	return cfg.DB
}

func GetAPIConfig() APIConfig {
	return cfg.API
}
