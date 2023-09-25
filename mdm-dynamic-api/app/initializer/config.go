package initializer

import (
	"app/domain"
	"github.com/spf13/viper"
)

func LoadConfig(path string) (config *domain.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("json")
	viper.SetConfigName("conf")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
