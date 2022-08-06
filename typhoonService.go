package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TyphoonService struct {
	StartYear  string
	EndYear    string
	StartMonth string
	EndMonth   string
}

const BASE_URL = "https://www.kma.go.kr/w//renew2021/rest/typhoon/search.do?json="

func (ts *TyphoonService) fetchTyphoons() ([]TyphoonModel, error) {
	// Sort Type: seq에 따라 정렬
	resp, err := http.Get(BASE_URL + url.QueryEscape(
		fmt.Sprintf(`{"startYear":"%v","endYear":"%v","startYear2":"%v","endYear2":"%v","startMonth":"%v","endMonth":"%v","startMonth2":"%v","endMonth2":"%v","sortType":"0"}`,
			ts.StartYear, ts.EndYear, ts.StartYear, ts.EndYear, ts.StartMonth, ts.EndMonth, ts.StartMonth, ts.EndMonth,
		),
	))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	typhoons := make([]TyphoonModel, 0)
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &typhoons)
	if err != nil {
		return nil, err
	}
	return typhoons, nil
}
