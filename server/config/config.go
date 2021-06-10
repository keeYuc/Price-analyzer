package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Header []Data `yaml:"header"`
	Gos    int64  `yaml:"gos"`
}

type Data struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

var config Config

func init() {
	fd, err := os.Open("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	filedata, err := ioutil.ReadAll(fd)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(filedata, &config)
	if err != nil {
		panic(err)
	}
}

func GetConfig() Config {
	return config
}
