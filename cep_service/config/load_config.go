package config

import "github.com/spf13/viper"

type EnviromentVar struct {
	ServiceURL string `mapstructure:"SERVICE_URL"`
	WebPort    string `mapstructure:"WEB_PORT"`
}

func LoadEnviromentVars(path string) (*EnviromentVar, error) {
	var env *EnviromentVar

	viper.SetConfigName("cep_service")
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
