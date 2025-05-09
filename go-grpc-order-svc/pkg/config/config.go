package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	DBHost        string `mapstructure:"DBHOST"`
	DBPort        string `mapstructure:"DBPORT"`
	DBUser        string `mapstructure:"DBUSER"`
	DBPassword    string `mapstructure:"DBPASSWORD"`
	DBName        string `mapstructure:"DBNAME"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return

}
