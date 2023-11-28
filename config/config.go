package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	Port        int    `mapstructure:"PORT"`
	DbURL       string `mapstructure:"DB_URL"`
	DbUser      string `mapstructure:"DB_USER"`
	DbPassword  string `mapstructure:"DB_PASSWORD"`
	DbName      string `mapstructure:"DB_NAME"`
	RedisURL    string `mapstructure:"REDIS_URL"`
	KafkaURL    string `mapstructure:"KAFKA_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv() // 환경 변수를 자동으로 읽어오도록 설정

	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to unmarshal config, %s", err)
	}

	return config, err
}
