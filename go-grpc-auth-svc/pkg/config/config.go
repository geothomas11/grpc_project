package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	SSLMode    string `mapstructure:"DB_SSLMODE"`
}

func (c *Config) GetDBUrl() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.SSLMode,
	)
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	fmt.Println("Using DB password:", config.DBPassword)

	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return

}
