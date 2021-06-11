package main

import (
	"fmt"
	"server/jobs"
	"server/mongo"
)

func main() {
	mongo.GetMg().Collection("price")
	r, err := jobs.GetRandRequest(1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Header.Values("cookie"))
}
