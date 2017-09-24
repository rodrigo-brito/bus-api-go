package config

import (
	"bytes"
	"io/ioutil"

	"path/filepath"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("yaml")
	path, err := filepath.Abs("./config/settings.yaml")
	if err != nil {
		panic(err)
	}
	settings, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	viper.ReadConfig(bytes.NewBuffer(settings))
}
