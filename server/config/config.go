package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

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

var confMutex sync.Mutex

func Get() Config {
	return config
}

func init() {
	confMutex.Lock()
	defer confMutex.Unlock()
	fd, err := os.Open("./config/config.yaml")
	if err != nil {
		fmt.Println("1")
		panic(err)
	}
	filebytes, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println("2")
		panic(err)
	}
	err = yaml.Unmarshal(filebytes, &config)
	if err != nil {
		fmt.Println("3")
		panic(err)
	}
}
