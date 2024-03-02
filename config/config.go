package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	EncryptionKey string `json:"encryptionKey"`
}

func LoadConfig(configPath string) (*viper.Viper, error) {
	v := viper.New()
	// set this before running
	v.AddConfigPath(configPath)
	v.SetConfigName("config")
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
