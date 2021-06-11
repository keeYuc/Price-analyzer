package main

import (
	"fmt"
	"service/conf"
)

func main() {
	fmt.Println(conf.Get().Mongo.Uri)
}
