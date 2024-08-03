package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/temaxuck/WUR/service.ebooks/internal/constants"
)

type Config struct {
	Host        string `mapstructure:"HOST"`
	Port        uint   `mapstructure:"PORT"`
	PostgresURL string `mapstructure:"POSTGRES_URL"`
	MaxFileSize uint   `mapstructure:"MAX_FILE_SIZE"` // in bytes
}

var configInstance *Config = nil

func (c *Config) loadConfig() (err error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("PORT", "8000")
	viper.SetDefault("MAX_FILE_SIZE", constants.GIGABYTE)

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
		return err
	}

	if err = viper.Unmarshal(&c); err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func GetConfig() (*Config, error) {
	var err error
	if configInstance == nil {
		configInstance = new(Config)
		err = configInstance.loadConfig()
	}
	return configInstance, err
}

func (c *Config) GetAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
