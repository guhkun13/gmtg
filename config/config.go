package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const envFilename string = ".env"

func LoadEnv() (env *EnvironmentVariable, err error) {
	log.Info().Msg("Load Env Here")

	v := viper.New()

	// read static env file
	v.SetConfigFile(envFilename)
	err = v.ReadInConfig()
	if err != nil {
		log.Error().Err(err).Str("filename", envFilename).Msg("viper error read config")
	}
	SetDefaultValue(v)

	v.AutomaticEnv()
	err = v.Unmarshal(&env)
	if err != nil {
		log.Error().Err(err).Msg("viper error unmarshal config")
	}

	return
}

func SetDefaultValue(v *viper.Viper) {
	v.SetDefault("IS_DEBUG_ENABLED", false)
}

type EnvironmentVariable struct {
	IsDebugEnabled bool `mapstructure:"IS_DEBUG_ENABLED"`
}
