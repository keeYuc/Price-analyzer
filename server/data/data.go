package data

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
)

func GetDataResponse(read io.Reader) (DataResponse, error) {
	r := DataResponse{}
	g, err := gzip.NewReader(read)
	defer g.Close()
	datas, err := ioutil.ReadAll(g)
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(datas, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

//* 直接拿json 去网站转化一下即可
type DataResponse struct {
	Data struct {
		Count int `json:"count"`
		List  []struct {
			Symbol             string  `json:"symbol"`
			NetProfitCagr      float64 `json:"net_profit_cagr"`
			Ps                 float64 `json:"ps"`
			Type               int     `json:"type"`
			Percent            float64 `json:"percent"`
			HasFollow          bool    `json:"has_follow"`
			TickSize           float64 `json:"tick_size"`
			PbTtm              float64 `json:"pb_ttm"`
			FloatShares        int     `json:"float_shares"`
			Current            float64 `json:"current"`
			Amplitude          float64 `json:"amplitude"`
			Pcf                float64 `json:"pcf"`
			CurrentYearPercent float64 `json:"current_year_percent"`
			FloatMarketCapital int64   `json:"float_market_capital"`
			MarketCapital      int64   `json:"market_capital"`
			DividendYield      float64 `json:"dividend_yield"`
			LotSize            int     `json:"lot_size"`
			RoeTtm             float64 `json:"roe_ttm"`
			TotalPercent       float64 `json:"total_percent"`
			Percent5M          float64 `json:"percent5m"`
			IncomeCagr         float64 `json:"income_cagr"`
			Amount             float64 `json:"amount"`
			Chg                float64 `json:"chg"`
			IssueDateTs        int64   `json:"issue_date_ts"`
			MainNetInflows     float64 `json:"main_net_inflows"`
			Volume             int     `json:"volume"`
			VolumeRatio        float64 `json:"volume_ratio"`
			Pb                 float64 `json:"pb"`
			Followers          float64 `json:"followers"`
			TurnoverRate       float64 `json:"turnover_rate"`
			FirstPercent       float64 `json:"first_percent"`
			Name               string  `json:"name"`
			PeTtm              float64 `json:"pe_ttm"`
			TotalShares        int     `json:"total_shares"`
		} `json:"list"`
	} `json:"data"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}
