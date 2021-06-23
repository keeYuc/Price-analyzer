package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fd, err := os.Open("./xxx.html")
	if err != nil {
		fmt.Println(1, err)
		return
	}
	g, err := goquery.NewDocumentFromReader(fd)
	if err != nil {
		fmt.Println(1, err)
		return
	}
	g.Find("#app").Filter("a[href]")
	if err != nil {
		fmt.Println(3, err)
		return
	}
	// fmt.Println(s)
}
