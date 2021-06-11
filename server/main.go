package main

import (
	"server/mongo"
)

func main() {
	mongo.GetMg().Collection("price")
}
