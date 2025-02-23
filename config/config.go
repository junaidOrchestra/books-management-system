package config

import (
	"github.com/spf13/viper"
	"log"
)

// Config struct to hold all configuration
type Config struct {
	Redis RedisConfig
	Kafka KafkaConfig
}
type KafkaConfig struct {
	Broker string
}

// RedisConfig holds Redis settings
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

var AppConfig Config

// InitConfig loads configuration from file
func InitConfig() {
	viper.SetConfigName("config")    // Filename without extension
	viper.SetConfigType("yaml")      // File type
	viper.AddConfigPath("../config") // Directory where config.yaml is stored

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	log.Println("✅ Configuration loaded successfully!")
}
