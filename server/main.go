package main

import (
	"fmt"
	"server/config"
)

// "github.com/tidwall/gjson"

func main() {
	fmt.Println(config.Get().CommonHead)
	// gjson.Get(string(datas))
}
