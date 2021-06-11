package conf

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Mongo Mongo `yaml:"mongo"`
}

type Mongo struct {
	Uri string `yaml:"uri"`
}

var conf *Config
var confMutex sync.Mutex

func Get() *Config {
	return conf
}

func init() {
	confMutex.Lock()
	defer confMutex.Unlock()
	fd, err := os.Open("./conf/conf.yaml")
	if err != nil {
		fmt.Println("1")
		panic(err)
	}
	filebytes, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println("2")
		panic(err)
	}
	err = yaml.Unmarshal(filebytes, &conf)
	if err != nil {
		fmt.Println("3")
		panic(err)
	}
}
