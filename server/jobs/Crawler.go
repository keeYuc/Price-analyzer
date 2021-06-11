package jobs

import (
	"fmt"
	"math/rand"
	"net/http"
	"server/config"
	"time"
)

func getHttpUrl(page, size int) string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100*100*10) + 1623397949815
	return fmt.Sprintf("https://xueqiu.com/service/v5/stock/screener/quote/list?page=%d&size=%d&order=desc&orderby=percent&order_by=percent&market=CN&type=sh_sz&_=%d", page, size, num)
}

//*已经设置好随机cookie的请求了 直接用
func GetRandRequest(page, size int) (*http.Request, error) {
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

// func GetTest() {
// 	c := http.Client{}
// 	response, err := c.Do(request)
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
