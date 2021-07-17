package jobs

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"server/config"
	"server/data"
	"server/mongo"
	"time"
)

func getHttpUrl(page, size int) string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100*100*10) + 1623397949815
	return fmt.Sprintf("https://xueqiu.com/service/v5/stock/screener/quote/list?page=%d&size=%d&order=desc&orderby=percent&order_by=percent&market=CN&type=sh_sz&_=%d", page, size, num)
}

//*已经设置好随机cookie的请求了 直接用
func getRandRequest(page, size int) (*http.Request, error) {
	rand.Seed(time.Now().UnixNano())
	request, err := http.NewRequest("GET", getHttpUrl(page, size), nil)
	if err != nil {
		return nil, err
	}
	for _, v := range config.Get().CommonHead {
		request.Header.Set(v.Key, v.Value)
	}
	ran_num := rand.Intn(len(config.Get().Cookies))
	request.Header.Set(config.Get().Cookies[ran_num].Key, config.Get().Cookies[ran_num].Value)
	return request, nil
}

func RunCrawler() {
	// size := config.Get().Size
	c := http.Client{}
	i := 136
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Printf("读到【%d】页了", i)
	}()
	for ; ; i++ {
		request, err := getRandRequest(i, 90)
		if err != nil {
			fmt.Println(err)
		}
		response, err := c.Do(request)
		if err != nil {
			fmt.Println(err)
		}
		temp, err := data.GetDataResponse(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		table := mongo.GetMg().Collection("price")
		for _, v := range temp.Data.List {
			table.InsertOne(context.TODO(), v)
		}
	}
}

// func GetTest() {
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer response.Body.Close()
// 	reader, err := gzip.NewReader(response.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer reader.Close()
// 	datas, err := ioutil.ReadAll(reader)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = ioutil.WriteFile("./test/data.json", datas, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// }
