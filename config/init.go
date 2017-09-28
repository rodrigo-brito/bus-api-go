package config

import (
	"bytes"
	"io/ioutil"

	"github.com/rodrigo-brito/bus-api-go/lib/environment"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("yaml")
	path := environment.AbsPath("./config/settings.yaml")
	settings, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	viper.ReadConfig(bytes.NewBuffer(settings))
}
