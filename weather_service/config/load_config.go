package config

import "github.com/spf13/viper"

type EnviromentVar struct {
	CepServicePath     string `mapstructure:"CEP_SERVICE_PATH"`
	WeatherServicePath string `mapstructure:"WEATHER_SERVICE_PATH"`
	WebPort            string `mapstructure:"WEB_PORT"`
}

func LoadEnviromentVars(path string) (*EnviromentVar, error) {
	var env *EnviromentVar

	viper.SetConfigName("channel-go")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&env)

	if err != nil {
		return nil, err
	}

	return env, nil
}
