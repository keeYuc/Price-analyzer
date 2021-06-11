package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/config"
)

func main() {
	url := "https://xueqiu.com/service/v5/stock/screener/quote/list?page=1&size=90&order=desc&orderby=percent&order_by=percent&market=CN&type=sh_sz&_=1623336144629s"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	for _, v := range config.Get().Header {
		request.Header.Set(v.Key, v.Value)
	}
	c := http.Client{}
	response, err := c.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	datas, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	f := make(map[string]interface{})
	json.Unmarshal(datas, f)
	fmt.Println(f)
	defer response.Body.Close()
}
