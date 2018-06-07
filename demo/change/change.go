package main

import (
	"github.com/warrially/BaziGo/Days"
	"github.com/warrially/BaziGo/Lunar"
	"log"
)

func main() {
	// 农历公里互转
	log.Println(
		Lunar.PrintLunar(
			Lunar.GetDateFrom64TimeStamp(
				Days.Get64TimeStamp(1990, 7, 21, 0, 0, 0))))
}
