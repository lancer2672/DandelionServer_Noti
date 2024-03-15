package utils

import "github.com/spf13/viper"

type Config struct {
	RABBITMQ_CONN string `mapstructure:"RABBITMQ_CONN"`
}

// overrided by env if exists
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return

}
