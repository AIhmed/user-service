package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBURI     string
	JWTSecret string
	JWTTTL    time.Duration
	HTTPPort  string
	GRPCPort  string
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No config file found, using defaults/environment")
	}

	return &Config{
		DBURI:     viper.GetString("DB_URI"),
		JWTSecret: viper.GetString("JWT_SECRET"),
		JWTTTL:    viper.GetDuration("JWT_TTL"),
		HTTPPort:  viper.GetString("HTTP_PORT"),
		GRPCPort:  viper.GetString("GRPC_PORT"),
	}
}
