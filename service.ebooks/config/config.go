package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Host        string `mapstructure:"HOST"`
	Port        uint   `mapstructure:"PORT"`
	PostgresURL string `mapstructure:"POSTGRES_URL"`
}

func (c *Config) LoadConfig() (err error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("PORT", "8000")

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

func (c *Config) GetAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
