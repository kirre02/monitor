package util

import (
	"log"

	"github.com/spf13/viper"
)

// Config stores all configuration for api.
type Config struct {
	DatabaseUrl      string `mapstructure:"DATABASE_URL"`
	MigrationPath    string `mapstructure:"MIGRATION_PATH"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`
	PostgresSslmode  string `mapstructure:"POSTGRES_SSLMODE"`
	DatabaseDriver   string `mapstructure:"DB_DRIVER"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}
	err = viper.Unmarshal(&config)
	return
}
