package tools

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/config"
)

func GetTest() {
	url := "https://xueqiu.com/service/v5/stock/screener/quote/list?page=2&size=90&order=desc&orderby=percent&order_by=percent&market=CN&type=sh_sz&_=1623397949815"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	for _, v := range config.Get().CommonHead {
		request.Header.Set(v.Key, v.Value)
	}
	c := http.Client{}
	response, err := c.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	reader, err := gzip.NewReader(response.Body)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	datas, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./test/data.json", datas, 0644)
	if err != nil {
		panic(err)
	}
}
