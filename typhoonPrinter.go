package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"math"
	"time"
)

type TyphoonPrinter struct {
	ts []TyphoonModel
}

func (tp TyphoonPrinter) WriteFile() {
	ioutil.WriteFile("tyhoon-history.txt", []byte(formatTyphoons(tp.ts)), fs.ModePerm.Perm())
}

func formatTyphoons(ts []TyphoonModel) string {
	result := ""
	for _, t := range ts {
		result += formatTyphoon(t)
		result += "\n"
	}
	return result
}

func formatTyphoon(t TyphoonModel) string {
	seq := fmt.Sprintf("%v", t.Seq)
	if t.Seq < 10 {
		seq = fmt.Sprintf("0%v", t.Seq)
	}
	result := fmt.Sprintf("%v%v %v", t.Year, seq, t.Name)
	result += "\n"

	for i := len(t.Tracks) - 1; i >= 0; i-- {
		result += formatTrack(t.Tracks[i])
		result += "\n"
	}

	return result
}

func formatTrack(t Track) string {
	date := formatDate(t.Date)                                // 날짜
	latitude, longitude := t.Latitude*10, t.Longitude*10      // 위도 경도
	phs := formatFloat(math.Round(t.WindSpeed * 3600 / 1000)) // 시속
	windSpeed := formatFloat(t.WindSpeed)
	rad15 := formatFloat(t.Rad15) // 강풍 반경

	ws25 := formatWs25(t.Ws25)             //폭풍반경
	strength := formatStrength(t.Strength) // 강도
	speed := formatFloat(t.Speed)
	result := fmt.Sprintf("%v  %v  %v  %v  %v  %v  %v  %v  %v  %v  %v", date, latitude, longitude, t.Pressure, windSpeed, phs, rad15, ws25, strength, t.Direction, speed)
	return result
}

func formatDate(d string) string {
	layout := "200601021504"
	t, err := time.Parse(layout, d)
	if err != nil {
		log.Fatal(err)
	}
	return t.Format("2006/01/02 15:04")
}

func formatStrength(s string) string {
	if s == "" {
		return "-"
	}
	return s
}

func formatWs25(s string) string {
	if s == "" || s == "0.0" || s == "-1.0" {
		return "-"
	}
	return s
}
func formatFloat(f float64) string {
	if f == 0.0 || f == -1 {
		return "-"
	}
	return fmt.Sprintf("%v", int(f))
}
