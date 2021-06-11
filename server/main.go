package main

import (
	"server/jobs"
	"server/mongo"
)

func main() {
	mongo.GetMg().Collection("price")
	jobs.RunCrawler()
}
