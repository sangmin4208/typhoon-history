package main

import (
	"log"
)

func main() {
	ts := &TyphoonService{
		StartYear:  "1977",
		EndYear:    "2021",
		StartMonth: "01",
		EndMonth:   "12",
	}
	typhoons, err := ts.fetchTyphoons()
	if err != nil {
		log.Fatal(err)
	}
	tp := TyphoonPrinter{ts: typhoons}
	tp.WriteFile()
}
