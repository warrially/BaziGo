package main

import (
	"github.com/warrially/BaziGo/Days"
	"github.com/warrially/BaziGo/Lunar"
	"log"
)

func main() {
	// 年月日时分秒 性别(1男0女)
	// bazi := BaziGo.GetBazi(1995, 12, 22, 21, 30, 0, 1)
	// log.Println(bazi)
	// BaziGo.PrintBazi(bazi)

	// 计算某一天的时间戳
	nTimeStamp := Days.Get64TimeStamp(2018, 5, 17, 0, 0, 0)

	//
	log.Println("农历是", Lunar.GetDateFrom64TimeStamp(nTimeStamp))
}
