package s3play

import (
	"github.com/spf13/viper"
)

type Config struct {
	AWS *AWS
}

type AWS struct {
	Endpoint        string `env:"AWS_ENDPOINT"`
	AccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
	Bucket          string `env:"S3_BUCKET"`
}

func provideConfig() (*Config, error) {
	cfg := &Config{}
	cfgFileName := "settings/setting.toml"

	viper.SetConfigType("toml")
	viper.SetConfigFile(cfgFileName)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return cfg, nil
}
