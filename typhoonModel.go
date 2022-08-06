package main

type TyphoonModel struct {
	Year   int     `json:"year"`
	Seq    int     `json:"seq"`
	Name   string  `json:"nameEn"`
	Tracks []Track `json:"tracks"`
}

type Track struct {
	Date      string  `json:"tm"`       //일시
	Latitude  float64 `json:"lat"`      // 위도
	Longitude float64 `json:"lon"`      // 경도
	Pressure  float64 `json:"ps"`       // 중심 기압
	WindSpeed float64 `json:"ws"`       // 최대 풍속 m/s
	Rad15     float64 `json:"rad15"`    //강풍 반경
	Ws25      string  `json:"ws25"`     //폭풍 반경
	Strength  string  `json:"strength"` // 강도
	Direction string  `json:"dir"`      // 진행방향
	Speed     float64 `json:"sp"`       // 이동속도
}
