package config

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	CityLatitude  float64 `mapstructure:"CITY_LATITUDE" json:"latitude,omitempty"`
	CityLongitude float64 `mapstructure:"CITY_LONGITUDE" json:"latitude,omitempty"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env, please copy .env.sample to .env, and fill the value.", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded : ", err)
	}

	return &env
}
